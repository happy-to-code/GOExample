package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

const json1 = `{
  "body": {
    "registration_information": {
      "register_basic_infomation": {},
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

func main() {
	value1 := gjson.Get(json1, "body.registration_information.registration_rights.basic_information_rights.basic_information_description.register_product_ref")
	fmt.Println("value--->", value1.String(), "<---")

	replace := strings.Replace(json1, "register_basic_infomation", "register_basic_information", 1)
	// fmt.Println(replace)

	value2 := gjson.Get(replace, "body.registration_information.register_basic_information")
	fmt.Println("value--->", value2.String(), "<---")
	result := gjson.Parse(json1)
	fmt.Println("result:", result.Map())

}
