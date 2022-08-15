package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var s = `{
    "product_information": {
        "test": [
            {
                "mytest_ref": "cvb"
            },
            {
                "mytest_ref": "bhj"
            }
        ],
        "escrow_information": {},
        "essential_information": {
            "basic_product_information": {
                "product_account_number_max": 200,
                "product_code": "Z00049",
                "product_create_time": "2020-07-16T00:00:00+08:00",
                "product_customer_browsing_right": 0,
                "product_info_disclosure_way": 0,
                "product_issuer_name": "XXXX",
                "product_issuer_subject_ref": "FB96EC73BC12D3F5AB13F14C438B1CEE",
                "product_market_subject_name": "上海股权托管交易中心股份有限公司",
                "product_market_subject_ref": "D7FD749660A7F2E2F38AD05CF0DA4D6D",
                "product_name": "XXXX",
                "product_name_abbreviation": "XXXX",
                "product_plate_trading_name": "98",
                "product_scale": 50000,
                "product_scale_currency": "156",
                "product_scale_unit": 3,
                "product_trading_market_category": 1,
                "product_type": 3,
                "product_type_function": 1
            },
            "product_file_information": [],
            "service_provider_information": [
                {
                    "service_provider_name": "上 海股权托管交易中心股份有限公司",
                    "service_provider_type": 1
                },
                {
                    "service_provide               r_name": "上海股权托管交易中心股份有限公司",
                    "service_provider_type": 5
                },
                {
                    "service_provide r_name": "上海股权托管交易中心股份有限公司",
                    "service_provider_type": 16
                }
            ]
        },
        "                product_subject_information": {
            "business_information": {},
            "fund_use_information": {
                "product_description_fund_use": "XXX",
                "product_fund_use_type": 1
            },
            "portfolio_information": {}
        },
        "release_information": {
            "filing_information": [],
            "private_convertible_bonds": {
                "product_amount_cashed": 0,
                "product_converting_shares_condition": "XXXX",
                "product_converting_shares_price_mode": "XXXX",
                "product_converting_shares_term": "XXXX",
                "product_due_date": "2023-06-03",
                "product_guarantee_measure": "XXXX",
                "product_issue_price": 100
            },
            "product_release_basic_information": {
                "product_release_date": "2020-07-20"
            }
        },
        "service_status_information": {},
        "transaction_information": {
            "delisting_information": {},
            "listing_information": {
                "product_listing_date": "2020-07-20"
            },
            "transaction_status": {
                "product_transaction_status": 1
            }
        }
    }
}`

func main() {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", m)
	// var path = "product_information.essential_information.basic_product_information.product_issuer_subject_ref"
	var path2 = "product_information.test"
	path2SplitList := strings.Split(path2, ".")
	fmt.Println("path2SplitList:", path2SplitList)
	//
	// for _, v := range path2SplitList {
	// 	value, ok := m[v]
	// 	if ok {
	// 		subjectQualificationInfoMap, ok := value.(map[string]interface{})
	// 	}
	// }

}
