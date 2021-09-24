package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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
	    "registration_information": {
	        "register_basic_information": {
	            "register_event_type": 1,
	            "register_object_type": 2
	        },
	        "registration_rights": {
	            "basic_information_rights": {
	                "authentic_right_record": {},
	                "available_registration": {},
	                "basic_information_description": {
	                    "register_product_ref": "B8EFCABC7922DF91631931EE252228FE"
	                },
	                "description_status_information": {},
	                "freezing_registration": {},
	                "pledge_registration": {}
	            }
	        },
	        "roll_records": {
	            "basic_information_roster": {
	                "register_product_ref": "B8EFCABC7922DF91631931EE252228FE"
	            },
	            "fund_investors": [],
	            "register_creditors": [],
	            "register_shareholders": [
	                {
	                    "register_asset_holding_status": 3,
	                    "register_equity_capital": 52000000,
	                    "register_equity_capital_paidin": 52000000,
	                    "register_equity_number": 52000000,
	                    "register_equity_shareholding": 58.69,
	                    "register_equity_subject_ref": "C57B91FEFB1C94EDD52F55F7585B1848",
	                    "register_equity_subject_type": 0
	                }
	            ]
	        }
	    }
    }
}
`

func main() {
	var err error

	//	获取 register_object_type
	registerObjectTypeResult := gjson.Get(jsonStr, "body.registration_information.register_basic_information.register_object_type")
	if registerObjectTypeResult.Int() == 1 {
		jsonStr, err = sjson.Delete(jsonStr, "body.registration_information.roll_records")
		if err != nil {
			panic(fmt.Errorf("删除名册登记列表错误：%v", err))
			return
		}
	} else {
		jsonStr, err = sjson.Delete(jsonStr, "body.registration_information.registration_rights")
		if err != nil {
			panic(fmt.Errorf("删除权利登记列表错误：%v", err))
			return
		}
	}

	valid, err := checkSchemaValid(jsonStr)
	fmt.Println("valid:", valid)
	fmt.Println("err:", err)

}

func checkSchemaValid(jsonStr string) (bool, error) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://E:/20.06.16Project/GOExample/src/A-2021-09-09/jsonSchemaD/Schema-v2.0.0.json")

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
