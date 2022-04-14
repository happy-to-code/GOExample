package main

import (
	"fmt"
	"log"
	"time"
)

type MysqlDBStore struct {
	ConnStr  string
	Name     string
	Password string
	DBName   string
}

type OracleDBStore struct {
	ConnStr  string
	Name     string
	Password string
	Port     int
}

var mysqlStore *MysqlDBStore
var oracleStore *OracleDBStore

func init() {

	oracleStore = &OracleDBStore{
		ConnStr:  "oracle:test@10.1.3.5",
		Name:     "admin",
		Password: "123456",
		Port:     3308,
	}

	Register("mysql", func(opt Opt) (IStore, error) {
		return newMysqlStroe(opt)
	})
}

func newMysqlStroe(opt Opt) (*MysqlDBStore, error) {
	store := &MysqlDBStore{
		ConnStr:  "mysql:test@10.1.3.1",
		Name:     "root",
		Password: "root",
		DBName:   "testDB",
	}
	return store, nil
}

type Opt struct {
	ConnectionStr string
	DupTxDuration time.Duration // 重复交易时间间隔
	IdleTimeout   time.Duration
}

type CreateStorer func(Opt) (IStore, error)

var providerMap = make(map[string]CreateStorer)

// Register 注册提供程序
func Register(provider string, fn CreateStorer) {
	providerMap[provider] = fn
}

// NewStorer 按名字创建存储提供程序
func NewStorer(provider string, opt Opt) IStore {
	if provider == "mysql" {
		// CreateWalletdb(opt.Connstr)
	}
	fn, has := providerMap[provider]
	if !has {
		log.Fatal(fmt.Errorf("Not support provider %s", provider))
	}
	storer, err := fn(opt)
	if err != nil {
		log.Fatal(err)
	}

	return storer
}

func main() {
	fmt.Println("=1=", providerMap)
	storer := providerMap["mysql"]
	fmt.Printf("=2=%s\n", storer)

	fmt.Printf("==>%+v\n", mysqlStore)
	fmt.Println(mysqlStore.MaxValue())
	fmt.Println(mysqlStore.FindByTxHash("hash1"))
	fmt.Println("==========================")
	fmt.Printf("======>%+v\n", oracleStore)
	fmt.Println(oracleStore.MaxValue())
	fmt.Println(oracleStore.FindByTxHash("hash2"))
}

// IStore 数据库接口
type IStore interface {
	MaxValue() uint64
	FindByTxHash(hash string) (error, string)
}

func (mysql *MysqlDBStore) MaxValue() uint64 {
	fmt.Println("mysql MaxValue")
	return 16
}

func (mysql *MysqlDBStore) FindByTxHash(hash string) (error, string) {
	fmt.Println("mysql FindByTxHash:", hash)
	return nil, "mysql"
}

func (oracle *OracleDBStore) MaxValue() uint64 {
	fmt.Println("oracle oracle MaxValue")
	return 18
}

func (oracle *OracleDBStore) FindByTxHash(hash string) (error, string) {
	fmt.Println("oracle oracle FindByTxHash::", hash)
	return nil, "oracle"
}
