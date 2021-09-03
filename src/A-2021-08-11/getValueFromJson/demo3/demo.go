package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var json1 = `
    "body": {
        "subject_information": {
            "basic_information_subject": {
                "general_information_subject": {
                    "subject_create_time": "2020-03-31T00:00:00+08:00",
                    "subject_main_administrative_region": 1,
                    "subject_type": 1
                },
                "subject_qualification_information": [
                    {
                        "investor_suitability_information": [],
                        "qualification_authentication_information": [
                            {
                                "subject_qualification_status": false
                            }
                        ]
                    }
                ]
            },
            "personal_subject_information": {
                "personal_subject_basic_information": {
                    "subject_birthday": "1949-09-03",
                    "subject_cellphone_number": "13704814500",
                    "subject_contact_address": "黑龙江省哈尔滨市香坊区珠江路106号313室",
                    "subject_contact_number": "0451-57775513",
                    "subject_gender": 2,
                    "subject_id_number": "230107194909030019",
                    "subject_id_type": 1,
                    "subject_investor_name": "王文学"
                }
            }
        }
    },
{
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
}
`

func main() {

	set, err := sjson.Set(json1, "body.subject_information.basic_information_subject.general_information_subject.subject_type", 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(set)

	subjectInvestorName := gjson.Get(json1, "body.subject_information.personal_subject_information.personal_subject_basic_information.subject_investor_name")
	fmt.Println("=====>", subjectInvestorName.String())
	fmt.Println("=====>", subjectInvestorName)
	subjectType := gjson.Get(json1, "body.subject_information.basic_information_subject.general_information_subject.subject_type")
	fmt.Printf("%T,%d", subjectType.Type, subjectType.Int())

	// value1 := gjson.Get(json1, "body.registration_information.registration_rights.basic_information_rights.basic_information_description.register_product_ref")
	// fmt.Println("value--->", value1.String(), "<---")

	// replace := strings.Replace(json1, "register_basic_infomation", "register_basic_information", 1)
	// fmt.Println(replace)

	// value2 := gjson.Get(replace, "body.registration_information.register_basic_information")
	// // value2 := gjson.Get(replace, "body.registration_information")
	// fmt.Println("value2--->", value2.String(), "<---")
	// fmt.Println(value2.String() == "{}")
	//
	// var s1 = `"register_basic_information":{
	// 			"register_object_type":1,
	// 			"register_event_type":1
	//           },
	// 		`
	// replace2 := strings.Replace(replace, "\"register_basic_information\": {},", s1, 1)
	// fmt.Println(replace2)

	// result := gjson.Parse(json1)
	// fmt.Println("result:", result.Map())

}
