package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var gdContractStr = `contract ShareContract{

    //平台公钥
    string PubKey = #########

    string ContractKey = "SGJ"
}`
	fmt.Println("0==>", gdContractStr)
	pubkey := "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE2pSmxNebDvgMXhdvopkdEJLm8EGX\nWm+YTdcha75A3vLlRgRBCqy7/fuhHCgriamrFTiwlc4DOnKEiNryp/PkZA==\n-----END PUBLIC KEY-----\n"
	fmt.Println("1==>", pubkey)
	pubkey = strconv.Quote(pubkey)
	fmt.Println("2==>", pubkey)

	replaceStr := strings.Replace(gdContractStr, "#########", pubkey, 1)
	fmt.Println("3==>", replaceStr)
}
