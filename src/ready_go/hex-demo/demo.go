package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main1() {
	// 0x00000002000001345696503e73a1a38e0227daeb0ecaa33df21be9a1632c24d0
	// 7d8200625207b9126ca10a01ee225c6a70ad06e6730f2842b5d0748aa24c2cf6

}

// HexDecode 16进制解码
func HexDecode(s string) []byte {
	dst := make([]byte, hex.DecodedLen(len(s))) // 申请一个切片, 指明大小. 必须使用hex.DecodedLen
	n, err := hex.Decode(dst, []byte(s))        // 进制转换, src->dst
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return dst[:n] // 返回0:n的数据.
}

// HexEncode 字符串转为16进制
func HexEncode(s string) []byte {
	dst := make([]byte, hex.EncodedLen(len(s))) // 申请一个切片, 指明大小. 必须使用hex.EncodedLen
	n := hex.Encode(dst, []byte(s))             // 字节流转化成16进制
	return dst[:n]
}

func main() {
	var aa uint64 = 1
	fmt.Println(aa)
	fmt.Println(aa - 1)
	s16 := "6769746875622e636f6d2f79657a696861636b"
	fmt.Println("HexDecode:", string(HexDecode(s16)))

	println(Chain_FANTOM.String())
}

type ChainName int

const (
	Chain_ETH ChainName = iota
	Chain_BSC
	Chain_HECO
	Chain_MATIC
	Chain_FANTOM
)

var chainNames = [...]string{
	Chain_ETH:    "ETH",
	Chain_BSC:    "BSC",
	Chain_HECO:   "HECO",
	Chain_MATIC:  "Polygon",
	Chain_FANTOM: "Fantom",
}

func (chain ChainName) String() string {
	return chainNames[chain]
}
