package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

var jsonStr = `{
    "header": {
        "model": {
            "protocol": "区域性股权市场跨链业务数据模型",
            "version": "2.0.0"
        },
        "content": {
            "type": "registration",
            "object_id": "5No00017aX5PxC51u",
            "version": 0,
            "operation": "create",
            "timestamp": 1513374000
        },
        "metadata": null
    },
    "body": {
	    "subject_information": {
	        "basic_information_subject": {
	            "general_information_subject": {
	                "subject_create_time": "2015-03-26T00:00:00+08:00",
	                "subject_main_administrative_region": 1,
	                "subject_type": 2
	            },
	            "subject_qualification_information": []
	        },
	        "personal_subject_information": {
	            "personal_subject_basic_information": {
	                "subject_birthday": "1970-04-05",
	                "subject_cellphone_number": "18637128818",
	                "subject_city": "",
	                "subject_contact_address": "河南省郑州市市辖区郑东新区商都路8号东2单元19层1902号",
	                "subject_contact_number": "0371-66987768",
	                "subject_education": 0,
	                "subject_gender": 2,
	                "subject_id_address": "河南省新乡市红旗区西马小营北街14号",
	                "subject_id_doc_mailbox": "1647611158@qq.com",
	                "subject_id_number": "372930197004052196",
	                "subject_id_type": 1,
	                "subject_investor_name": "刘春生",
	                "subject_personal_fax": "0371-55985955",
	                "subject_postal_code": "450001",
	                "subject_province": ""
	            }
	        }
	    }
    }
}
`

func main() {
	var err error

	valid, err := checkSchemaValid(jsonStr)
	fmt.Println("valid:", valid)
	fmt.Println("err:", err)

}

func checkSchemaValid(jsonStr string) (bool, error) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://E:/20.06.16Project/GOExample/src/A-2021-09-09/jsonSchemaD/Schema-v2.0.0.1.json")

	// 读取验证数据
	documentLoader := gojsonschema.NewStringLoader(jsonStr)

	// 验证
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
		return false, err
	}

	var detail []string
	// 验证结果
	if result.Valid() {
		return true, nil
	} else {
		for _, desc := range result.Errors() {
			detail = append(detail, fmt.Sprintf("%s\n", desc))
		}
		return false, fmt.Errorf("%v", detail)
	}
}
