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
                "register_object_type": 1
            },
            "registration_rights": {
                "basic_information_rights": {
                    "authentic_right_record": {},
                    "available_registration": {
                        "register_asset_balance_after": 1687500
                    },
                    "basic_information_description": {
                        "register_asset_currency": "156",
                        "register_asset_type": 0,
                        "register_asset_unit": 4,
                        "register_authentic_right_recognition_status": 1,
                        "register_create_time": "2021-07-08T00:00:00+08:00",
                        "register_product_ref": "B5F3D6C472C5F27F2CFDDAB40A88EC25",
                        "register_serial_number": "30042724000100073276",
                        "register_subject_account_ref": "9FE737CE40662530D66AF997299C2EB6",
                        "register_subject_ref": "A56BEE2E257BF5C451E073DABEDC611B",
                        "register_subject_type": 1,
                        "register_time": "2021-07-08T00:00:00+08:00"
                    },
                    "description_status_information": {
                        "register_asset_holding_status": 3,
                        "register_source_type": 1
                    },
                    "freezing_registration": {
                        "register_frozen_balance": 0
                    },
                    "pledge_registration": {}
                }
            },
            "roll_records": {
                "basic_information_roster": {
                    "register_product_ref": "B5F3D6C472C5F27F2CFDDAB40A88EC25",
                    "register_subject_ref": "A56BEE2E257BF5C451E073DABEDC611B"
                },
                "fund_investors": null,
                "register_creditors": null,
                "register_shareholders": null
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
	fmt.Println(jsonStr)
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
