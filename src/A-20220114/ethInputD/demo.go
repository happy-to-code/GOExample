package main

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"os"
)

func main() {
	if reader, err := os.Open("E:\\20.06.16Project\\GOExample\\src\\A-20220114\\ethInputD\\contract.abi"); err != nil {
		log.Fatal(err)
	} else {
		if tokenAbi, err := abi.JSON(reader); err != nil {
			log.Fatal(err)
		} else {
			encodedData := "a9059cbb00000000000000000000000033978e63a9ff9d7cd24ef28d45ad2ba89bd36cd2000000000000000000000000000000000000000000000000000000000005047e"
			decodeData, err := hex.DecodeString(encodedData)
			if err != nil {
				log.Fatal(err)
			}
			// a9059cbb == transfer
			if method, ok := tokenAbi.Methods["transfer"]; ok {
				params, err := method.Inputs.Unpack(decodeData[4:])
				if err != nil {
					log.Fatal(err)
				}
				log.Println(params)
			}
		}
	}

}
