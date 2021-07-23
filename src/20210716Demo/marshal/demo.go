package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	s := `{
"disas_23ce27fb56624cf8b2e9f3fa348fd62d": "",
"disas_6702009ff1fb47b896870491ec4af866": "",
"disas_ea93736ee86d47cf94efd873d8ff16d0": "",
"disas_ee7590d450924827942ca41fc966acc4": "",
"disas_4iHwYoqhLBSGumU2NH7etBBrStH6ezh6JPq2e7nnjbnrWb7NCD": "",
"disas_96a6816363364daba9b8ed806997579e": "{\"id\":\"96a6816363364daba9b8ed806997579e\",\"name\":\"大数据局\",\"method\":\"1\",\"itemNum\":1,\"status\":1,\"createdTime\":\"2021-07-12 17:20:05\",\"updatedTime\":\"2021-07-12 17:20:05\",\"publicKey\":\"-----BEGIN PUBLIC KEY-----\\nMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE3PqyPjf6cRguXXOt0fEZsbVWRM4e\\n2HpH84Lkb3X00OC+i+6gnA5hw2QBbVHf+GdYrfqkjs6L3CFZ4M5eWDr+9A==\\n-----END PUBLIC KEY-----\",\"url\":\"http://172.21.43.42:60002/jml/v1/opendata/query/do\",\"dataItemList\":[{\"id\":\"b01f43ca-ea1f-40ee-b5b7-01e9edaf075d\",\"dataSourceId\":\"96a6816363364daba9b8ed806997579e\",\"name\":\"test\",\"nameCode\":\"\",\"url\":\"\",\"requestType\":0,\"createTime\":\"\",\"useType\":0,\"mold\":0}]}",
"disas_d04e2e75d11c46288b587acfab56b978": "",
"disas_4a243249e0824ba48fe80522679648df": "",
"disas_66e7b9e2327e47a0be70ea059f8b35d0": "",
"disas_b3d3dfbf542642439598b7140dabd652": "",
"disas_ae058a06316c48e69b6af1753bc1d4a2": "",
"disas_4151150192af42a0a00249a79dda1c28": "",
"disas_9bc482df1c234712949d3ad8faa6df36": "",
"disas_00f1cfa7c0f24c58b6f9248cb599762e": "",
"disas_90b9feddcf0141a69b24adddb9dcdd1b": "",
"disas_d9100131d80a4266a8a3e118c89846cd": "",
"disas_620bb71ddf4d407883d691f05f4e9445": "",
"disas_aa0255f434c5402d8f2af73e35772156": "",
"disas_3fd2651793c04ac08d6c9837047b7a53": "",
"disas_3c755e1192614a45b0ba7f75c483c0f7": "",
"disas_588e3f5dd00c424297bf2dbd321b070f": "",
"disas_621d8cdb6d0343e2af0658bac64e582d": "",
"disas_9e6076e7d1b54870991e1454b9f8b6c6": "",
"disas_fa8cc4f9f12b4d5583c482df3be9911d": "",
"disas_1d82d9fdeeb8450d8d23aa98c05e62aa": "",
"disas_3dc4255724324d779d52aabbfaaa9d70": ""
}`

	var m map[string]string
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}

	var newM = make(map[string]string)
	for k, v := range m {
		if v == "" {
			continue
		}
		newM[k] = v
	}
	fmt.Println(newM)

	var dsiList []DataSourceAndItems
	for _, v := range newM {
		var dsi DataSourceAndItems
		err := json.Unmarshal([]byte(v), &dsi)
		if err != nil {
			log.Println("数据项反序列化错误：", err)
			continue
		}
		dsiList = append(dsiList, dsi)
	}
	fmt.Println("=================")
	fmt.Println(dsiList)

}

type DataSourceAndItems struct {
	Id          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Method      string `json:"method" binding:"required"`
	ItemNum     int    `json:"itemNum"` // 项目数
	Status      int    `json:"status"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
	PublicKey   string `json:"publicKey" binding:"required"` // 公钥
	Url         string `json:"url" binding:"required"`       // 数据源的url

	DataItemList []DataItem `json:"dataItemList" binding:"required,dive"` // 数据项
}

type DataItem struct {
	Id           string `json:"id" binding:"required"`   //  唯一ID
	DataSourceId string `json:"dataSourceId"`            //  接口拥有者
	Name         string `json:"name" binding:"required"` //  接口名称
	NameCode     string `json:"nameCode"`                //  接口名称编码（比如：marriage,grid）
	Url          string `json:"url"`                     //  接口路径
	RequestType  int    `json:"requestType"`             //  请求类型(0:GET;1:POST),默认是0
	CreateTime   string `json:"createTime"`              //  创建时间（格式：yyyy-MM-dd HH:mm:ss），默认是当前时间
	UseType      int    `json:"useType"`                 //  使用类型（0：自然人；1：法人；2：通用）默认是0
	Mold         int    `json:"mold"`                    //  分类（0：身份；1：资产；2：风险;3:其它） 默认是0
}
