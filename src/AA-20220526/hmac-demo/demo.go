package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

func main() {
	var applayDep = ApplyDepository{
		Appid:            "appid",
		NotarialOfficeId: "notarialOfficeId",
		Category:         "category",
		NotaryData:       "NotaryData",
		TxIds:            []string{"751899ede6ebde781b51b45d8eaf0d13e92d7160c10b679ae33af8d086b5331d"},
		File:             []byte("file"),
		CertName:         "cert001",
		CertObjectStr:    "CertObjectStr",
	}

	s := sign(applayDep, time.Now().UnixMilli(), []byte("yZYMDea3pk0sl2wzq"))
	fmt.Println("s:", s)

}

type ApplyDepository struct {
	Appid            string   `json:"appid"`
	NotarialOfficeId string   `json:"notarialOfficeId"`
	Category         string   `json:"category"`
	NotaryData       string   `json:"notaryData"`
	TxIds            []string `json:"txIds"`
	File             []byte   `json:"file"`
	CertName         string   `json:"certName"`
	CertObjectStr    string   `json:"certMetaData"`
}

func sign(v interface{}, timestamp int64, secret []byte) string {
	fmt.Println("======================>", v)
	jsondata, err := json.Marshal(v)
	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	dict := make(map[string]interface{})
	json.Unmarshal(jsondata, &dict)
	fmt.Printf("dict:%+v\n", dict)
	delete(dict, "appid")
	delete(dict, "token")
	values := make([]string, 0, len(dict))
	for k := range dict {
		values = append(values, k)
	}
	sort.Strings(values)

	fmt.Println("====>", values)
	var buf bytes.Buffer
	for _, key := range values {
		v := dict[key]
		vList, ok := v.([]interface{})

		if ok {
			by, _ := json.Marshal(vList)
			fmt.Println("--->", string(by))
			// sprintf := fmt.Sprintf("\"%s\"", vList[0])
			fmt.Fprintf(&buf, "%s", by)
		} else {
			fmt.Fprintf(&buf, "%v", v)
		}
	}
	fmt.Fprintf(&buf, "%d", timestamp)

	fmt.Println("sign str:===>", buf.String())

	h := hmac.New(sha1.New, secret)
	h.Write(buf.Bytes())
	result := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}
