package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)
import "github.com/xeipuuv/gojsonschema"

func main() {
	// 	s := `{
	//     "body": {
	//         "subject_information": {
	//             "basic_information_subject": {
	//                 "general_information_subject": {
	//                     "subject_create_time": "2013-08-07",
	//                     "subject_main_administrative_region": 1,
	//                     "subject_type": 1
	//                 },
	//                 "subject_qualification_information": [
	//                     {
	//                         "investor_suitability_information": [
	//                             {}
	//                         ],
	//                         "qualification_authentication_information": [
	//                             {
	//                                 "subject_qualification_status": true,
	//                                 "subject_review_time": "2013-08-07"
	//                             }
	//                         ],
	//                         "subject_market_roles_type": 1,
	//                         "subject_qualification_category": 1
	//                     },
	//                     {
	//                         "investor_suitability_information": [
	//                             {}
	//                         ],
	//                         "qualification_authentication_information": [
	//                             {
	//                                 "subject_qualification_status": true,
	//                                 "subject_review_time": "2013-12-26"
	//                             }
	//                         ],
	//                         "subject_market_roles_type": 3,
	//                         "subject_qualification_category": 1
	//                     }
	//                 ]
	//             },
	//             "organization_subject_information": {
	//                 "basic_information_of_enterprise": {
	//                     "basic_information_description": {
	//                         "subject_business_scope": "上海隆光蜃景光电科技有限公司成立于2005年，注册于上海市闵行区，毗邻上海虹桥交通枢纽；是一家专注于LED照明产品的研发、制造、营销和城市空间照明设计为一体的高新技术企业；在研发大功率LED芯片的集成封装、LED灯具的配光、结构、散热、外观等方面，获得了十多项具有独立自主知识产权的专利技术；创建了“隆光”和“蜃景”两大品牌，其产品涵盖了LED功能照明和LED夜景照明两大类，广泛应用于室内公共照明、户外道路、隧道、园林、城市景观等领域的功能性照明和夜景照明。公司曾参与2010上海世博会的LED灯光建设项目，夜景照明产品在2010年上海世博会建设中被大量采用。“隆光”牌大功率LED路灯和LED筒灯（天花灯）分别被认定为“上海市自主创新产品”和2011年度上海“品牌产品”。2006年公司已通过了ISO9001：2000质量体系认证，2008年并被认定为上海市高新技术企业；同时公司于2010年获得闵行区“创新型科普示范企业”；2012年度获得“中国LED照明行业最具创新力企业”荣誉称号；被上海市政府列为上海市“专精特新”重点培育企业；并入围2010年中国（上海）半导体照明灯具推荐产品大赛优秀产品名单，承接了多个上海市",
	//                         "subject_city": "310100",
	//                         "subject_company_english_name": "",
	//                         "subject_company_name": "上海隆光蜃景光电科技有限公司",
	//                         "subject_company_profile": "",
	//                         "subject_company_short_name": "蜃景光电",
	//                         "subject_company_type": 1,
	//                         "subject_contact_address": "",
	//                         "subject_contact_number": "021-13764848486",
	//                         "subject_district": "310112",
	//                         "subject_document_information": [
	//                             {
	//                                 "code": "310112000560684",
	//                                 "type": 1
	//                             }
	//                         ],
	//                         "subject_fax": "",
	//                         "subject_internet_address": "",
	//                         "subject_legal_type": 3,
	//                         "subject_mail_box": "012345678@sjled.cn",
	//                         "subject_office_address": "",
	//                         "subject_organization_nature": 1,
	//                         "subject_paid_in_capital": 11000000,
	//                         "subject_paid_in_capital_currency": "156",
	//                         "subject_postal_code": "",
	//                         "subject_province": "310000",
	//                         "subject_registered_address": "上海市闵行区中春路7001号第1幢10楼1002室",
	//                         "subject_registered_capital": 11000000,
	//                         "subject_registered_capital_currency": "156",
	//                         "subject_scale_type": 0,
	//                         "subject_shareholders_number": 0
	//                     },
	//                     "leading_member_information": [
	//                         {
	//                             "subject_document_type": 5,
	//                             "subject_key_personnel_id": "",
	//                             "subject_key_personnel_name": "谢金崇",
	//                             "subject_key_personnel_type": 3
	//                         }
	//                     ]
	//                 }
	//             }
	//         }
	//     },
	//     "header": {
	//         "content": {
	//             "object_id": "212EC15142807BD8DBC50141D2514E7B",
	//             "operation": "update",
	//             "timestamp": 1625249405,
	//             "type": "account",
	//             "version": 3
	//         },
	//         "metadata": null,
	//         "model": {
	//             "protocol": "区域性股权市场跨链业务数据模型",
	//             "version": "2.0.0"
	//         }
	//     }
	// }`
	s := `{
    "body": {
        "subject_information": {
            "basic_information_subject": {
                "general_information_subject": {
                    "subject_create_time": "2017-03-31T00:00:00+08:00",
                    "subject_main_administrative_region": 1,
                    "subject_type": 2
                },
                "subject_qualification_information": [
                    {
                        "investor_suitability_information": [
                            {
                                "subject_investor_qualification": 2,
                                "subject_investor_qualification_certificate_time": "2017-10-11T00:00:00+08:00",
                                "subject_investor_qualification_status": false
                            },
                            {
                                "subject_investor_qualification": 1,
                                "subject_investor_qualification_certificate_time": "2017-10-11T00:00:00+08:00",
                                "subject_investor_qualification_status": true
                            }
                        ],
                        "qualification_authentication_information": [
                            {
                                "subject_qualification_status": true,
                                "subject_review_time": "2017-10-11T00:00:00+08:00"
                            }
                        ]
                    }
                ]
            },
            "personal_subject_information": {
                "personal_subject_basic_information": {
                    "subject_id_number": "91310115051261120U",
                    "subject_id_type": 0,
                    "subject_investor_name": "上海瑞紫投资管理有限公司"
                }
            }
        }
    },
    "header": {
        "content": {
            "object_id": "FB80D6C961BBE0E54204768F79DBF09B",
            "operation": "update",
            "timestamp": 1489449600,
            "type": "subject",
            "version": 1
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

var schemaLoader = gojsonschema.NewReferenceLoader("file://E:/20.06.16Project/GOExample/src/20210716Demo/jsonSchema/Schema-Alpha-100.json")

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
