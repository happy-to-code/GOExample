package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	_ "github.com/tidwall/sjson"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// 	var json1 = `{
	//         "body": {
	//             "registration_information": {
	//                 "register_basic_information": {
	//                     "register_object_type":1,
	//                     "register_event_type":1
	// 				},
	//                 "registration_rights": {
	//                     "basic_information_rights": {
	//                         "authentic_right_record": {},
	//                         "available_registration": {},
	//                         "basic_information_description": {
	//                             "register_asset_currency": "156",
	//                             "register_asset_type": 0,
	//                             "register_asset_unit": 0,
	//                             "register_create_time": "2021-07-16T00:00:00+08:00",
	//                             "register_product_ref": "217201CFEE7BB52533480B816C0918B3/0",
	//                             "register_serial_number": "131453",
	//                             "register_subject_account_ref": "D910606933C72811C3B9E6AA15EF76C3/0",
	//                             "register_subject_ref": "C998AA18CAF818CCCCBEF1747954E9F4/10",
	//                             "register_subject_type": 1,
	//                             "register_time": "2021-07-16T00:00:00+08:00",
	//                             "register_transaction_ref": "3130303832313230323130373136313331343534/0"
	//                         },
	//                         "description_status_information": {
	//                             "register_source_type": 3
	//                         },
	//                         "freezing_registration": {
	//                             "register_frozen_balance": 0,
	//                             "register_frozen_balance_change": 0
	//                         },
	//                         "pledge_registration": {}
	//                     }
	//                 },
	//                 "roll_records": {
	//                     "basic_information_roster": {
	//                         "register_product_ref": "217201CFEE7BB52533480B816C0918B3/0",
	//                         "register_subject_ref": "C998AA18CAF818CCCCBEF1747954E9F4/10"
	//                     },
	//                     "fund_investors": [],
	//                     "register_creditors": [],
	//                     "register_shareholders": []
	//                 }
	//             }
	//         },
	//         "header": {
	//             "content": {
	//                 "object_id": "25c3863df99e4ae09af2d99e3d6984c4",
	//                 "operation": "create",
	//                 "timestamp": 1626775775,
	//                 "type": "registration",
	//                 "version": 0
	//             },
	//             "metadata": null,
	//             "model": {
	//                 "protocol": "区域性股权市场跨链业务数据模型",
	//                 "version": "2.0.0"
	//             }
	//         }
	//     }
	// `
	var json1 = `
{
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
                    "authentic_right_record": {
                        "register_authentic_right_recognition_date": "2025-12-16",
                        "register_right_recognition_agent_subject_name": "确权代理方主体名称CHARACTER",
                        "register_right_recognition_agent_subject_ref": "RRRASR624/0",
                        "register_right_recognition_description": "确权描述信息CHARACTER",
                        "register_right_recognition_doc": [
                            {
                                "file_name": "1file1.pdf",
                                "file_number": "2",
                                "hash": "da1234filehash52221",
                                "summary": "简述1",
                                "term_of_validity": "2018-10-23",
                                "term_of_validity_type": "0",
                                "url": "http://test.com/file/201/1file1.pdf"
                            },
                            {
                                "file_name": "2file1.pdf",
                                "file_number": "2",
                                "hash": "da1234filehash52222",
                                "summary": "简述2",
                                "term_of_validity": "2018-10-23",
                                "term_of_validity_type": "0",
                                "url": "http://test.com/file/201/2file1.pdf"
                            }
                        ],
                        "register_right_recognition_mode": 1,
                        "register_right_recognition_subject_name": "确权方主体名称CHARACTER",
                        "register_right_recognition_subject_ref": "RRRSR624/0"
                    },
                    "available_registration": {
                        "register_asset_balance_after": 3000,
                        "register_asset_balance_before": 2000,
                        "register_asset_balance_change": 1000
                    },
                    "basic_information_description": {
                        "register_asset_currency": "156",
                        "register_asset_type": 1,
                        "register_asset_unit": 1,
                        "register_authentic_right_recognition_status": 1,
                        "register_create_time": "2020-12-16T12:00:25+08:00",
                        "register_description": "登记描述信息CHARACTER",
                        "register_product_ref": "prodEqp4dNCofJ3CfH/0",
                        "register_serial_number": "登记流水号CHARACTER",
                        "register_subject_account_ref": "RSAR624/0",
                        "register_subject_ref": "RSR624/0",
                        "register_subject_type": 1,
                        "register_time": "2020-12-1612:00:25+08:00"
                    },
                    "description_status_information": {
                        "register_asset_equity_type": 1,
                        "register_asset_holding_nature": 1,
                        "register_asset_holding_status": 1,
                        "register_asset_holding_status_description": "持有状态说明TEXT",
                        "register_asset_note": "登记说明TEXT",
                        "register_asset_verification_doc": [
                            {
                                "file_name": "1file1.pdf",
                                "file_number": "2",
                                "hash": "da1234filehash52221",
                                "summary": "简述1",
                                "term_of_validity": "2018-10-23",
                                "term_of_validity_type": "0",
                                "url": "http://test.com/file/201/1file1.pdf"
                            },
                            {
                                "file_name": "2file1.pdf",
                                "file_number": "2",
                                "hash": "da1234filehash52222",
                                "summary": "简述2",
                                "term_of_validity": "2018-10-23",
                                "term_of_validity_type": "0",
                                "url": "http://test.com/file/201/2file1.pdf"
                            }
                        ],
                        "register_source_type": 0
                    },
                    "freezing_registration": {
                        "register_frozen_balance": 1000,
                        "register_frozen_balance_after": 5000,
                        "register_frozen_balance_change": 500,
                        "register_thaw_description": "冻结/解冻说明信息TEXT",
                        "register_thaw_doc": [
                            {
                                "file_name": "1file1.pdf",
                                "file_number": "2",
                                "hash": "da1234filehash52221",
                                "summary": "简述1",
                                "term_of_validity": "2018-10-23",
                                "term_of_validity_type": "0",
                                "url": "http://test.com/file/201/1file1.pdf"
                            },
                            {
                                "file_name": "2file1.pdf",
                                "file_number": "2",
                                "hash": "da1234filehash52222",
                                "summary": "简述2",
                                "term_of_validity": "2018-10-23",
                                "term_of_validity_type": "0",
                                "url": "http://test.com/file/201/2file1.pdf"
                            }
                        ]
                    },
                    "pledge_registration": {
                        "register_pledge_balance_after": 5000,
                        "register_pledge_balance_before": 1000,
                        "register_pledge_balance_change": 1000
                    }
                }
            },
            "roll_records": {
                "basic_information_roster": {
                    "register_product_ref": "prodEqp4dNCofJ3CfH/0",
                    "register_subject_ref": "RSR624/0"
                }
            }
        }
    }
}
`

	registerObjectTypeResult := gjson.Get(json1, "body.registration_information.register_basic_information.register_object_type")
	fmt.Println(registerObjectTypeResult)
	fmt.Println(registerObjectTypeResult.Int() == 1)

	// set, err := sjson.Delete(json1, "body.registration_information.roll_records")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(set)

	valid, err := checkSchemaValid(json1)
	fmt.Println("valid:", valid)
	fmt.Println("err:", err)

}

func checkSchemaValid(jsonStr string) (bool, error) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://E:/20.06.16Project/GOExample/src/A-2021-08-11/jsonSchemaD/Schema-v2.0.0.json")

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
