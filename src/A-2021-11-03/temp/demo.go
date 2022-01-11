package main

import (
	"encoding/json"
	"fmt"
)

// Sample 样本集
type Sample struct {
	Name           string          // 样本集名称
	Relation       string          // 关联数据源
	Owner          string          // 所有者
	Description    string          // 描述
	Authorizations []Authorization // 授权列表
	Fields         []Field         // 授权字段列表
	PreField       []string        // 预处理字段
	Seglength      int32           // 分片长度
	Status         int
}

// Authorization 授权列表
type Authorization struct {
	Member      string // 成员代号
	ExpireStart int64  // 授权起始时间
	ExpireEnd   int64  // 授权截止时间
	AuthId      string // 授权ID
	Account     string // 授权成员地址
	Grade       int    // 授权等级
}

// Field 字段信息
type Field struct {
	Field       string // 字段名称
	Attribute   string // 字段属性
	Description string // 字段描述
	Grade       int    // 授权等级
}

func main() {
	var fields = []Field{
		{
			Field:       "年龄",
			Attribute:   "age",
			Description: "age",
			Grade:       1,
		},
		{
			Field:       "姓名",
			Attribute:   "name",
			Description: "百家姓",
			Grade:       2,
		},
	}
	var authorizations = []Authorization{
		{
			Member:      "xm",
			ExpireStart: 1,
			ExpireEnd:   2,
			AuthId:      "id1",
			Account:     "account1",
			Grade:       3,
		},
		{
			Member:      "zs",
			ExpireStart: 3,
			ExpireEnd:   4,
			AuthId:      "id2",
			Account:     "account2",
			Grade:       2,
		},
	}

	var sample = Sample{
		Name:           "sample1",
		Relation:       "www.baidu.com",
		Owner:          "lisi",
		Description:    "description",
		Authorizations: authorizations,
		Fields:         fields,
		PreField:       []string{"age"},
		Seglength:      2,
		Status:         1,
	}

	fmt.Println(sample)
	marshal, _ := json.Marshal(sample)
	fmt.Println(string(marshal))

}
