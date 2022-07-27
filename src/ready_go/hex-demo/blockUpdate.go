package main

import (
	"github.com/go-pg/pg/v10"
	"time"
)

var (
	db *pg.DB
)

func main() {

}
func init() {
	// db = pg.Connect(&pg.Options{
	// 	Database: info.Database,
	// 	Addr:     fmt.Sprintf("%s:%d", info.Host, info.Port),
	// 	User:     info.User,
	// 	Password: info.Pass,
	// })
}

type BlockMode struct {
	tableName   struct{}  `pg:"block_info"`
	BlockNumber int64     `pg:"block_number"`
	Hash        string    `pg:"block_hash"`
	ParentHash  string    `pg:"parent_hash"`
	Miner       string    `pg:"miner"`
	Difficulty  string    `pg:"difficulty"`
	GasLimit    uint64    `pg:"gas_limit"`
	GasUsed     uint64    `pg:"gas_used"`
	CreatedTime string    `pg:"created_time"`
	Timestamp   int64     `pg:"timestamp"`
	Nonce       string    `pg:"nonce"`
	TxCount     int       `pg:"tx_count"`
	CreatedAt   time.Time `pg:"created_at"`
}
