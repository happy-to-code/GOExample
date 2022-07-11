package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	_ "github.com/lib/pq"
	"github.com/siddontang/go/log"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	httpClient *http.Client
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout),
		},
	}
	return client
}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) (error, string) {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("httpDo err:%v\n", err), ""
	}

	req.Header.Set("Content-Type", contentType)

	resp, e := client.Do(req)
	if e != nil {
		return fmt.Errorf("接口调用出错：%v\n", e), ""
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		return fmt.Errorf("httpDo er:%v\n", er), ""
	}
	return nil, string(body)
}

const (
	UrlPrefix = "http://10.1.3.116:3334"
	// UrlPrefix = "http://10.1.4.29:8545"
)

func hex2dec(hex string) int64 {
	if strings.HasPrefix(hex, "0x") {
		hex = hex[2:]
	}

	n, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		log.Errorf("【%s】转换失败，err:%v", hex, err)
		return -1
	}
	return n
}

func weiToEther(hex string) *big.Float {
	val := hex[2:]
	wei := new(big.Int)
	wei.SetString(val, 16)

	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()

	// 定义接受参数的变量
	// var start int
	// var end int
	//
	// flag.IntVar(&start, "s", start, "开始区块")
	// flag.IntVar(&end, "e", end, "结束区块")
	//
	// // 解析命令行参数写入注册的flag里,在所有flag都注册之后，这一步一定不能少
	// flag.Parse()
	// if start >= end {
	// 	log.Errorf("结束区块：%d必须要大于开始区块:%d", end, start)
	// 	return
	// }
	os.Stdout, _ = os.Create("/var/log/log_block.log")
	log := log.NewDefault(os.Stdout)
	// 查询已经处理的最大区块高度
	err, handledBlockNum := selectMaxHasHandledBlockNum()
	if err != nil {
		time.Sleep(5 * time.Second)
		for j := 0; j < 5; j++ {
			err, handledBlockNum = selectMaxHasHandledBlockNum()
			if err != nil {
				time.Sleep(5 * time.Second)
				continue
			}
		}
	}
	if err != nil {
		panic(err)
	}
	// handledBlockNum = 11052984

	for i := handledBlockNum + 1; ; i++ {
		// 根据区块高度查询区块和交易信息
		blockNumStr := strconv.FormatInt(int64(i), 16)
		var data = `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x%s", true],"id":1}`
		data = fmt.Sprintf(data, blockNumStr)
		err, res := httpDo(httpClient, "POST", UrlPrefix, "application/1.json", data)
		if err != nil {
			log.Errorf("根据区块高度查询区块和交易信息出错：%v", err)
			continue
		}

		var blockObject BlockObject
		err = json.Unmarshal([]byte(res), &blockObject)
		if err != nil {
			log.Warn(res, ",反序列化出错：", err)
			continue
		}
		// fmt.Printf("==>blockObject:%+v\n", blockObject)

		if blockObject.Result.Timestamp == "" {
			log.Infof("第[%d]块Timestamp is null ,continue", i)
			continue
		}
		// 区块的时间戳
		blockTimeStampHexStr := blockObject.Result.Timestamp
		timeStamp := hex2dec(blockTimeStampHexStr)
		// 旷工地址
		minerAddr := strings.TrimSpace(blockObject.Result.Miner)

		if blockObject.Jsonrpc != "" && len(blockObject.Result.Transactions) > 0 {
			var txHistorys []TxHistory

			for _, tx := range blockObject.Result.Transactions {
				// 根据交易hash查询 receipt
				err, receipt := getTransactionReceipt(tx.Hash)
				if err != nil {
					log.Errorf("根据交易hash查询 receipt出错：", err)
					continue
				}
				if strings.TrimSpace(tx.To) != "" {
					gasUsed := receipt.Result.GasUsed
					if strings.HasPrefix(gasUsed, "0x") {
						gasUsed = gasUsed[2:]
					}
					// 此次交易使用掉的gas 单位是位
					wei := new(big.Int)
					wei.SetString(gasUsed, 16)
					gasUsedInt64 := wei.Int64()

					// 转账金额
					transferAmount := weiToEther(tx.Value)
					// 0x1
					txHistory := TxHistory{
						Txhash:         tx.Hash,
						From:           tx.From,
						To:             tx.To,
						Amount:         transferAmount,
						OriginalAmount: tx.Value,
						// Token:
						GasUsed: gasUsedInt64,
						// GasPrice:
						ChainName:   "ETH",
						CreatedTime: timeStamp,
						BlockNumber: int64(i),
						// Status:
					}
					if receipt.Result.Status == "0x1" {
						txHistory.Status = 1
					}

					gasPriceOriginal := tx.GasPrice
					if strings.HasPrefix(gasPriceOriginal, "0x") {
						gasPriceOriginal = gasPriceOriginal[2:]
					}

					gasPrice := new(big.Int)
					gasPrice.SetString(gasPriceOriginal, 16)
					txHistory.GasPrice = gasPrice.Int64()
					txHistory.Token = "ETH"

					txHistorys = append(txHistorys, txHistory)
				}

				// 2、更新余额表中数据 todo
				// f, _ := transferAmount.Float64()
				// // 2.1 处理from地址数据
				// err, fromAssetBalance := selectAssetBalanceByAddress(tx.From)
				// if err == nil {
				// 	fromAssetBalance.Balance = fromAssetBalance.Balance - f
				//
				// } else {
				// 	// log.Warnf("根据地址[%s]查询formAssetBalance失败,%v", tx.From, err)
				// 	sql := "INSERT INTO \"asset_balance\"(\"address\", \"token_symbol\", \"balance\", \"last_tradetime\", \"last_tradetxhash\", \"input_count\", \"output_count\", \"input_maxvalue\", \"output_maxvalue\") VALUES "
				// 	fmt.Sprintf("")
				//
				// }
				// 2.1 处理to地址数据

				// 3、更新合约数据
				if tx.To == "" || tx.To == " " || tx.To == "null" || tx.To == "nil" || tx.To == "0x" {
					log.Infof("[%d][%s]===>receipt:%v", i, tx.Hash, receipt)
					// 	合约交易
					contractInfo := ContractInfo{
						Address:              receipt.Result.ContractAddress,
						Creator:              tx.From,
						ChainName:            "ETH",
						CreatedTime:          timeStamp,
						CreatedTxHash:        []int64{int64(i), hex2dec(tx.TransactionIndex)},
						ContractCreationCode: tx.Input,
						Nonce:                tx.Nonce,
					}
					if strings.TrimSpace(contractInfo.Address) == "" {
						// 自己计算合约地址
						nonce := uint64(hex2dec(tx.Nonce))
						contractAddr := crypto.CreateAddress(common.HexToAddress(tx.From), nonce)

						contractInfo.Address = contractAddr.Hex()
					}
					if contractInfo.Address != "" {
						updateContractInfo(contractInfo)
					}
				}

				// 4、更新账号地址表数据
				err = updateAccountAddress(tx.From, tx.To, timeStamp, i, tx.TransactionIndex)
				if err != nil {
					log.Warnf("区块[%d]更新账号地址表数据失败，err:%v", i, err)
					return
				}

			}
			// 1、将数据批量存到tx_history中
			if len(txHistorys) > 0 {
				err := batchInsertDataToTxHistory(txHistorys)
				if err != nil {
					log.Warnf("区块[%d]将数据批量存到tx_history中失败，err:%v", i, err)
					return
				}
			}

			// 	将旷工地址插入数据库
			if minerAddr != "" {
				err = insertMinerAddress2DB(minerAddr, timeStamp, i)
				if err != nil {
					log.Warnf("区块[%d]将旷工地址插入数据库失败，err:", i, err)
					return
				}
			}
		}

		log.Infof("第[%d]块处理结束", i)
	}

}

func updateContractInfo(info ContractInfo) {
	sql := "INSERT INTO \"contract_info\"(\"address\", \"creator\", \"chain_name\", \"created_time\", \"created_txhash\", \"contract_creation_code\", \"nonce\") VALUES "
	sprintf := fmt.Sprintf("('%s','%s','%s',TO_TIMESTAMP(%d),ARRAY[%d,%d],'%s','%s')", info.Address, info.Creator, info.ChainName, info.CreatedTime, info.CreatedTxHash[0], info.CreatedTxHash[1], info.ContractCreationCode, info.Nonce)
	BatchInsert(sql + sprintf)
}

func updateAccountAddress(from string, to string, stamp int64, i int64, index string) error {
	sql := "INSERT INTO \"account_address\"(\"address\", \"chain_name\", \"first_occurtime\", \"first_occurtx\") VALUES "

	sprintf1 := fmt.Sprintf("('%s','%s',TO_TIMESTAMP(%d),ARRAY[%d,%d])", from, "ETH", stamp, i, hex2dec(index))
	sprintf2 := fmt.Sprintf("('%s','%s',TO_TIMESTAMP(%d),ARRAY[%d,%d])", to, "ETH", stamp, i, hex2dec(index))

	err2 := BatchInsert(sql + sprintf1)
	if err2 != nil && !strings.Contains(err2.Error(), "duplicate key value") {
		return err2
	}

	err2 = BatchInsert(sql + sprintf2)
	if err2 != nil && !strings.Contains(err2.Error(), "duplicate key value") {
		return err2
	}

	return nil
}

func insertMinerAddress2DB(addr string, stamp int64, i int64) error {
	// 新增旷工地址
	minerSql := "INSERT INTO \"account_address\"(\"address\", \"chain_name\", \"first_occurtime\", \"first_occurtx\", \"tags\") VALUES "
	sprintf3 := fmt.Sprintf("('%s','%s',TO_TIMESTAMP(%d),ARRAY[%d,%d],ARRAY['%s'])", addr, "ETH", stamp, i, -1, "miner")
	err2 := BatchInsert(minerSql + sprintf3)
	if err2 != nil && !strings.Contains(err2.Error(), "duplicate key value") {
		return err2
	}
	return nil
}

func selectAssetBalanceByAddress(addr string) (error, AssetBalance) {
	var assetBalance AssetBalance
	sprintf := fmt.Sprintf("SELECT * FROM  \"asset_balance\" where address = '%s' %s", addr, "limit 1")
	rows, err := db.Query(sprintf)
	if err != nil {
		log.Errorf("PG Statements Wrong:%v ", err)
		return err, assetBalance
	}

	for rows.Next() {
		if err := rows.Scan(&assetBalance.Address, &assetBalance.TokenSymbol, &assetBalance.Balance,
			&assetBalance.LastTradeTime, &assetBalance.LastTradeTxHash, &assetBalance.InputCount, &assetBalance.OutputCount,
			&assetBalance.InputMaxValue, &assetBalance.OutputMaxValue); err != nil {
			log.Fatal("PG Rows Scan Failed: ", err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal("PG Query Failed: ", err)
		return err, assetBalance
	}

	rows.Close()
	db.Close()
	return nil, assetBalance
}

func selectMaxHasHandledBlockNum() (error, int64) {
	sprintf := fmt.Sprintf("SELECT \"max\"(block_number) FROM \"tx_history\"")
	rows, err := db.Query(sprintf)
	if err != nil {
		log.Errorf("PG Statements Wrong: %v", err)
		return err, 0
	}

	var blockNumber int64
	for rows.Next() {
		if err := rows.Scan(&blockNumber); err != nil {
			log.Errorf("PG Rows Scan Failed:%v\n ", err)
			return err, 0
		}
	}

	if err := rows.Err(); err != nil {
		log.Errorf("PG Query Failed: ", err)
		return err, 0
	}

	rows.Close()

	return nil, blockNumber

}

var db *sql.DB
var err error

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pg_host, pg_port, pg_user, pg_password, pg_dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Errorf("Connect PG Failed:%v ", err)
	}

	db.SetMaxOpenConns(50)
	// db.SetMaxIdleConns(1000)
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Errorf("Ping GP Failed: %v", err)
	}
}

// PGs数据库信息
const (
	pg_host     = "b2904236d6.zicp.vip"
	pg_port     = 22021
	pg_user     = "root"
	pg_password = "wutongchain@ps"
	pg_dbname   = "blockdb"

	// pg_host     = "b2904236d6.zicp.vip"
	// // pg_host     = "121.239.164.179"
	// pg_port     = 22021
	// pg_user     = "root"
	// pg_password = "wutongchain@ps"
	// pg_dbname   = "testdb"
)

func BatchInsert(sqlStr string) error {
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Errorf("PG Statements Wrong:%v ", err)
		return err
	}

	res, err := stmt.Exec()
	if err != nil {
		// log.Fatal("PG Statements Exec Wrong: ", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Fatal("PG Affecte Wrong: ", err)
		return err
	}

	// fmt.Println(id)
	return nil
}
func batchInsertDataToTxHistory(txHistorys []TxHistory) error {
	sql := "INSERT INTO \"tx_history\"(\"txhash\", \"from\", \"to\", \"amount\", \"token\", \"gas_used\", \"gas_price\", \"chain_name\", \"created_time\", \"block_number\", \"status\") VALUES "

	var s = ""
	for i, history := range txHistorys {
		sprintf := fmt.Sprintf("('%s','%s','%s',%f,'%s',%d,%d,'%s',TO_TIMESTAMP(%d),%d,%d)", history.Txhash, history.From, history.To, history.Amount, history.Token, history.GasUsed, history.GasPrice, history.ChainName, history.CreatedTime, history.BlockNumber, history.Status)

		if i == 0 {
			s = fmt.Sprintf("%s%s", s, sprintf)
		} else {
			s = fmt.Sprintf("%s,%s", s, sprintf)
		}
	}
	sql = sql + s

	// fmt.Println("sql:", sql)
	return BatchInsert(sql)
}

func getTransactionReceipt(hash string) (error, Receipt) {
	var receipt Receipt
	var data = `{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["%s"],"id":1}`
	data = fmt.Sprintf(data, hash)
	err, res := httpDo(httpClient, "POST", UrlPrefix, "application/1.json", data)
	if err != nil {
		log.Errorf("根据交易hash查询Receipt信息出错：%v", err)
		return err, receipt
	}

	err = json.Unmarshal([]byte(res), &receipt)
	if err != nil {
		log.Errorf("Receipt反序列化出错：%v", err)
		return err, receipt
	}
	return nil, receipt
}

type Receipt struct {
	Jsonrpc string `1.json:"jsonrpc"`
	Id      int    `1.json:"id"`
	Result  struct {
		BlockHash         string `1.json:"blockHash"`
		BlockNumber       string `1.json:"blockNumber"`
		ContractAddress   string `1.json:"contractAddress"`
		CumulativeGasUsed string `1.json:"cumulativeGasUsed"`
		EffectiveGasPrice string `1.json:"effectiveGasPrice"`
		From              string `1.json:"from"`
		GasUsed           string `1.json:"gasUsed"`
		Logs              []struct {
			Address          string   `1.json:"address"`
			Topics           []string `1.json:"topics"`
			Data             string   `1.json:"data"`
			BlockNumber      string   `1.json:"blockNumber"`
			TransactionHash  string   `1.json:"transactionHash"`
			TransactionIndex string   `1.json:"transactionIndex"`
			BlockHash        string   `1.json:"blockHash"`
			LogIndex         string   `1.json:"logIndex"`
			Removed          bool     `1.json:"removed"`
		} `1.json:"logs"`
		LogsBloom        string `1.json:"logsBloom"`
		Status           string `1.json:"status"`
		To               string `1.json:"to"`
		TransactionHash  string `1.json:"transactionHash"`
		TransactionIndex string `1.json:"transactionIndex"`
		Type             string `1.json:"type"`
	} `1.json:"result"`
}

type BlockObject struct {
	Jsonrpc string `1.json:"jsonrpc"`
	Id      int    `1.json:"id"`
	Result  Result `1.json:"result"`
}

type Result struct {
	Difficulty       string        `1.json:"difficulty"`
	ExtraData        string        `1.json:"extraData"`
	GasLimit         string        `1.json:"gasLimit"`
	GasUsed          string        `1.json:"gasUsed"`
	Hash             string        `1.json:"hash"`
	LogsBloom        string        `1.json:"logsBloom"`
	Miner            string        `1.json:"miner"`
	MixHash          string        `1.json:"mixHash"`
	Nonce            string        `1.json:"nonce"`
	Number           string        `1.json:"number"`
	ParentHash       string        `1.json:"parentHash"`
	ReceiptsRoot     string        `1.json:"receiptsRoot"`
	Sha3Uncles       string        `1.json:"sha3Uncles"`
	Size             string        `1.json:"size"`
	StateRoot        string        `1.json:"stateRoot"`
	Timestamp        string        `1.json:"timestamp"`
	TotalDifficulty  string        `1.json:"totalDifficulty"`
	Transactions     []Transaction `1.json:"transactions"`
	TransactionsRoot string        `1.json:"transactionsRoot"`
	Uncles           []string      `1.json:"uncles"`
}

type Transaction struct {
	BlockHash        string `1.json:"blockHash"`
	BlockNumber      string `1.json:"blockNumber"`
	From             string `1.json:"from"`
	Gas              string `1.json:"gas"`
	GasPrice         string `1.json:"gasPrice"`
	Hash             string `1.json:"hash"`
	Input            string `1.json:"input"`
	Nonce            string `1.json:"nonce"`
	To               string `1.json:"to"`
	TransactionIndex string `1.json:"transactionIndex"`
	Value            string `1.json:"value"`
	Type             string `1.json:"type"`
	V                string `1.json:"v"`
	R                string `1.json:"r"`
	S                string `1.json:"s"`
}

// TxHistory ==================================================================
type TxHistory struct {
	Txhash         string     `1.json:"txhash"`
	From           string     `1.json:"from"`
	To             string     `1.json:"to"`
	Amount         *big.Float `1.json:"amount"`
	OriginalAmount string     `1.json:"original_amount"`
	Token          string     `1.json:"token"`
	GasUsed        int64      `1.json:"gas_used"`
	GasPrice       int64      `1.json:"gas_price"`
	ChainName      string     `1.json:"chain_name"`
	CreatedTime    int64      `1.json:"created_time"`
	BlockNumber    int64      `1.json:"block_number"`
	Status         int        `1.json:"status"`
}

type ContractInfo struct {
	Address              string   `1.json:"address"`
	Creator              string   `1.json:"creator"`
	Standard             []string `1.json:"standard"`
	ChainName            string   `1.json:"chain_name"`
	CreatedTime          int64    `1.json:"created_time"`
	CreatedTxHash        []int64  `1.json:"created_tx_hash"`
	Tags                 []string `1.json:"tags"`
	TokenSymbol          string   `1.json:"token_symbol"`
	TokenTotal           float64  `1.json:"token_total"`
	Code                 string   `1.json:"code"`
	Abi                  string   `1.json:"abi"`
	ContractCreationCode string   `1.json:"contract_creation_code"`
	From                 string   `1.json:"from"`
	Nonce                string   `1.json:"nonce"`
}

type AssetBalance struct {
	Address         string          `1.json:"address"`
	TokenSymbol     string          `1.json:"token_symbol"`
	Balance         float64         `1.json:"balance"`
	LastTradeTime   time.Time       `1.json:"last_trade_time"`
	LastTradeTxHash BlockAndTxIndex `1.json:"last_trade_tx_hash"`
	InputCount      int64           `1.json:"input_count"`
	OutputCount     int64           `1.json:"output_count"`
	InputMaxValue   float64         `1.json:"input_max_value"`
	OutputMaxValue  float64         `1.json:"output_max_value"`
}

type BlockAndTxIndex string

func (bat BlockAndTxIndex) BlockNumber() uint64 {
	return 0
}
