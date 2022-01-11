package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

// Block step1:创建Block结构体
type Block struct {
	// 字段：
	// 高度Height：其实就是区块的编号，第一个区块叫创世区块，高度为0
	Height int64
	// 上一个区块的哈希值ProvHash：
	PrevBlockHash []byte
	// 交易数据Data：目前先设计为[]byte,后期是Transaction
	Data []byte
	// 时间戳TimeStamp：
	TimeStamp int64
	// 哈希值Hash：32个的字节，64个16进制数
	Hash []byte

	Nonce int64
}

// NewBlock step2：创建新的区块
func NewBlock(data string, provBlockHash []byte, height int64) *Block {
	// 创建区块
	block := &Block{
		height,
		provBlockHash,
		[]byte(data),
		time.Now().Unix(),
		nil,
		0,
	}
	// 设置哈希值
	// block.SetHash()
	// return block

	// 调用工作量证明的方法，并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// SetHash step3:设置区块的hash
func (block *Block) SetHash() {
	// 1.将高度转为字节数组
	heightBytes := IntToHex(block.Height)
	// fmt.Println(heightBytes)
	// 2.时间戳转为字节数组
	// timeBytes:=IntToHex(block.TimeStamp)
	// 转为二进制的字符串
	// fmt.Println(block.TimeStamp)
	// fmt.Printf("%x,%b\n",block.TimeStamp,block.TimeStamp)
	timeString := strconv.FormatInt(block.TimeStamp, 2)
	// fmt.Println("timeString:",timeString)
	timeBytes := []byte(timeString)
	// fmt.Println("timeStamp:",timeBytes)
	// 3.拼接所有的属性
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes}, []byte{})
	// 4.生成哈希值
	hash := sha256.Sum256(blockBytes) // 数组长度32位
	block.Hash = hash[:]
}

// CreateGenesisBlock step4:创建创世区块：
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, make([]byte, 32, 32), 0)
}

// Serilalize 将区块序列化，得到一个字节数组---区块的行为，设计为方法
func (block *Block) Serilalize() []byte {
	// 1.创建一个buffer
	var result bytes.Buffer
	// 2.创建一个编码器
	encoder := gob.NewEncoder(&result)
	// 3.编码--->打包
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// DeserializeBlock 反序列化，得到一个区块---设计为函数
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	var reader = bytes.NewReader(blockBytes)
	// 1.创建一个解码器
	decoder := gob.NewDecoder(reader)
	// 解包
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
