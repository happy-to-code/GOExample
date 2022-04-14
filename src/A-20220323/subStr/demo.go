package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	var s = "0xa0f3f6ab"
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	s = s[:8]
	fmt.Println(s)

	fmt.Println("----------------")
	hexstr := "38ed17390000000000000000000000000000000000000000000000056bc75e2d63100000000000000000000000000000000000000000000000000009689df51ba16c41b000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000002f5820e55d904cfc6c43371a4352958a0f573be900000000000000000000000000000000000000000000000000000000623bd37c0000000000000000000000000000000000000000000000000000000000000003000000000000000000000000ae1107fc7cef1294f09185ac475c9886527dcd8a000000000000000000000000bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d56"
	data, _ := hex.DecodeString(hexstr)
	fmt.Println(string(data))

	var key1 = "/v1/bsc11/block/getbynum/24dfsdf53135"
	var key2 = "/v1/*/block/getByNum/@"
	// key2 = key2[:strings.Index(key2, "@")]
	// fmt.Println(key2)
	// fmt.Println(strings.Split(key1, "/"))
	// split := strings.Split(key2, "/")
	// fmt.Println(strings.Split(key2, "/"))
	// for _, s2 := range split {
	// 	fmt.Println("--->", s2, "<---")
	// }

	fmt.Println(KeyMatch(key1, key2))
}
func KeyMatch(key1, key2 string) bool {
	key1 = strings.ToLower(key1)
	key2 = strings.ToLower(key2)
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}
	index := strings.LastIndex(key2, "@")
	if index != -1 {
		key2 = key2[:index]
	}
	split1 := strings.Split(key1, "/")
	split2 := strings.Split(key2, "/")
	if len(split1) == 0 || len(split2) != len(split1) {
		return false
	}
	for i := 1; i < len(split1); i++ {
		s2 := strings.TrimSpace(split2[i])
		if s2 == "*" || s2 == "" {
			continue
		}
		s1 := strings.TrimSpace(split1[i])
		if s1 != s2 {
			return false
		}
	}
	return true
}

// 0xa0f3f6ab
