package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"time"
)

type BlockChain struct {
	// Blocks []*Block //存储有序的区块
	Tip []byte   // 最近的取快递Hash值
	DB  *bolt.DB // 数据库对象
}

/*
	1.仅仅用来创建区块链
	如果数据库存在，证明区块链存在，直接结束该方法
	否则进行创建创世区块，并存入数据库中
*/
func CreateBlockChainWithGenesisBlock(data string) {
	if dbExists() {
		fmt.Println("数据库已经存在。。。")
		return
	}
	//
	fmt.Println("创建创世区块：", data)
	// 2.数据库不存在，说明第一次创建，然后存入到数据库中
	fmt.Println("数据库不存在。。")
	// A：创建创世区块
	// 创建创世区块
	genesisBlock := CreateGenesisBlock(data)
	// B：打开数据库
	db, err := bolt.Open(DBNAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// C：存入数据表
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(BLOCKTABLENAME))
		if err != nil {
			log.Panic(err)
		}
		if b != nil {
			err = b.Put(genesisBlock.Hash, genesisBlock.Serilalize())
			if err != nil {
				log.Panic("创世区块存储有误。。。")
			}
			// 存储最新区块的hash
			b.Put([]byte("l"), genesisBlock.Hash)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	// 返回区块链对象
	// return &BlockChain{genesisBlock.Hash, db}
}

func (bc *BlockChain) AddBlockToBlockChain(data string) {
	// 创建新区块
	// newBlock := NewBlock(data,prevHash,height)
	// 添加到切片中
	// bc.Blocks = append(bc.Blocks,newBlock)
	// 1.更新数据库
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		// 2.打开表
		b := tx.Bucket([]byte(BLOCKTABLENAME))
		if b != nil {
			// 2.根据最新块的hash读取数据，并反序列化最后一个区块
			blockBytes := b.Get(bc.Tip)
			lastBlock := DeserializeBlock(blockBytes)
			// 3.创建新的区块
			newBlock := NewBlock(data, lastBlock.Hash, lastBlock.Height+1)
			// 4.将新的区块序列化并存储
			err := b.Put(newBlock.Hash, newBlock.Serilalize())
			if err != nil {
				log.Panic(err)
			}
			// 5.更新最后一个哈希值，以及blockchain的tip
			b.Put([]byte("l"), newBlock.Hash)
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// 判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(DBNAME); os.IsNotExist(err) {
		return false
	}
	return true
}

// Iterator
// 2.获取一个迭代器的方法
func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.Tip, bc.DB}
}
func (bc *BlockChain) PrintChains() {
	// 1.获取迭代器对象
	bcIterator := bc.Iterator()
	// 2.循环迭代
	for {
		block := bcIterator.Next()
		fmt.Printf("第%d个区块的信息：\n", block.Height+1)
		// 获取当前hash对应的数据，并进行反序列化
		fmt.Printf("\t高度：%d\n", block.Height)
		fmt.Printf("\t上一个区块的hash：%x\n", block.PrevBlockHash)
		fmt.Printf("\t当前的hash：%x\n", block.Hash)
		fmt.Printf("\t数据：%s\n", block.Data)
		// fmt.Printf("\t时间：%v\n", block.TimeStamp)
		fmt.Printf("\t时间：%s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("\t次数：%d\n", block.Nonce)
		// 3.直到父hash值为0
		hashInt := new(big.Int)
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(hashInt) == 0 {
			break
		}
	}
}

// GetBlockchainObject 新增方法，用于获取区块链
func GetBlockchainObject() *BlockChain {
	/*
	   1.如果数据库不存在，直接返回nil
	   2.读取数据库
	*/
	if !dbExists() {
		fmt.Println("数据库不存在，无法获取区块链。。")
		return nil
	}
	db, err := bolt.Open(DBNAME, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	var blockchain *BlockChain
	// B：读取数据库
	err = db.View(func(tx *bolt.Tx) error {
		// C：打开表
		b := tx.Bucket([]byte(BLOCKTABLENAME))
		if b != nil {
			// D：读取最后一个hash
			hash := b.Get([]byte("l"))
			// E：创建blockchain
			blockchain = &BlockChain{hash, db}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return blockchain
}
