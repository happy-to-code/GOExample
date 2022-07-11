package main

import (
	"encoding/json"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"reflect"
)

func main() {
	s := `{
    "body": {
        "registration_information": {
            "register_basic_information": {
				"register_object_type":1,
				"register_event_type":1
			},
            "registration_rights": {
                "basic_information_rights": {
                    "authentic_right_record": {},
                    "available_registration": {},
                    "basic_information_description": {
                        "register_asset_currency": "156",
                        "register_asset_type": 0,
                        "register_asset_unit": 0,
                        "register_create_time": "2021-07-16T00:00:00+08:00",
                        "register_product_ref": "217201CFEE7BB52533480B816C0918B3/0",
                        "register_serial_number": "131453",
                        "register_subject_account_ref": "D910606933C72811C3B9E6AA15EF76C3/0",
                        "register_subject_ref": "C998AA18CAF818CCCCBEF1747954E9F4/10",
                        "register_subject_type": 1,
                        "register_time": "2021-07-16T00:00:00+08:00",
                        "register_transaction_ref": "3130303832313230323130373136313331343534/0"
                    },
                    "description_status_information": {
                        "register_source_type": 3
                    },
                    "freezing_registration": {
                        "register_frozen_balance": 0,
                        "register_frozen_balance_change": 0
                    },
                    "pledge_registration": {}
                }
            },
            "roll_records": {
                "basic_information_roster": {
                    "register_product_ref": "217201CFEE7BB52533480B816C0918B3/0",
                    "register_subject_ref": "C998AA18CAF818CCCCBEF1747954E9F4/10"
                },
                "fund_investors": [],
                "register_creditors": [],
                "register_shareholders": []
            }
        }
    },
    "header": {
        "content": {
            "object_id": "25c3863df99e4ae09af2d99e3d6984c4",
            "operation": "create",
            "timestamp": 1626775775,
            "type": "registration",
            "version": 0
        },
        "metadata": null,
        "model": {
            "protocol": "区域性股权市场跨链业务数据模型",
            "version": "2.0.0"
        }
    }
}`

	validJsonSchema(s)

}

var schemaLoader = gojsonschema.NewReferenceLoader("file://E:/20.06.16Project/GOExample/src/20210716Demo/jsonSchema/Schema-v2.0.0.1.json")

func validJsonSchema(jsonStr string) {
	documentLoader := gojsonschema.NewStringLoader(jsonStr)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid:%+v\n", result.Valid())
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			// fmt.Errorf("%s", desc)
			fmt.Printf("- %s\n", desc)
		}
	}
	fmt.Println("00000000000000000000")
}

func genSuperviseData(tmp []byte, src map[string]interface{}) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal(tmp, &dict); err != nil {
		return nil, err
	}
	return doGenSuperviceData(dict, src), nil
}

func doGenSuperviceData(dict, src map[string]interface{}) map[string]interface{} {
	obj := make(map[string]interface{})
	fnAssignVal := func(cols []interface{}) map[string]interface{} {
		m := make(map[string]interface{})
		for i := range cols {
			key := cols[i].(string)
			if val, has := src[key]; has {
				m[key] = val
			}
		}
		return m
	}

	for k, v := range dict {
		switch typ := v.(type) {
		case string:
			obj[k] = src[k]
		case []interface{}:
			obj[k] = fnAssignVal(v.([]interface{}))
		case map[string]interface{}:
			obj[k] = doGenSuperviceData(v.(map[string]interface{}), src)
		case map[string][]interface{}:
			obj[k] = fnAssignVal(v.([]interface{}))
		default:
			fmt.Println("unknown type:", reflect.TypeOf(typ))
		}
	}
	return obj
}
