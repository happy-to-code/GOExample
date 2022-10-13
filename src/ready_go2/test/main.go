package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	jsonTemplate := `{
    "product_information": {
        "essential_information":{
            "basic_product_information":[
                "product_trading_market_category",
                "product_market_subject_ref",
                "product_market_subject_name",
                "product_plate_trading_name",
                "product_issuer_subject_ref",
                "product_issuer_name",
				"product_issuer_type",
                "product_code",
                "product_name",
                "product_name_abbreviation",
                "product_type_function",
                "product_type",
                "product_account_number_max",
                "product_info_disclosure_way",
                "product_scale_unit",
                "product_scale_currency",
                "product_scale",
                "product_customer_browsing_right",
                "product_issuer_contact_person",
                "product_issuer_contact_info",
				"product_remark",
                "product_create_time"
            ],
            "service_provider_information":"service_provider_information",
            "product_file_information":"product_file_information"
        },
        "product_subject_information":{
            "fund_use_information":[
                "product_fund_use_type",
				"product_fund_sub_category",
                "product_description_fund_use",
                "product_document_describing_funds"
            ],
            "business_information":[
                "product_business_purpose_name",
                "product_business_purpose_details",
                "product_business_purpose_documents"
            ],
            "portfolio_information":[
                "product_investment_products_type",
                "product_investment_proportion_range",
                "product_Investment_product_details",
                "product_detailed_description_document"
            ]
        },
        "release_information": {
			"product_release_basic_information":[
				"product_release_date"
			],
            "filing_information":"filing_information",
            "private_convertible_bonds": [
                "product_scope_issue",
                "product_bond_duration_unit",
                "product_bond_duration",
                "product_filing_amount",
                "product_by_stages",
                "product_bond_interest_rate_floor",
                "product_bond_interest_rate_cap",
                "product_staging_frequency",
                "product_initial_issue_amount",
                "product_initial_ratio",
                "product_initial_interest_rate",
                "product_interest_calculation_method",
                "product_interest_calculation_method_remarks",
                "product_payment_method",
				"product_payment_status",
                "product_payment_method_remarks",
				"product_cashed_count",
				"product_stock_count",
                "product_is_appoint_repayment_date",
                "product_appoint_repayment_date",
                "product_guarantee_measure",
                "product_converting_shares_condition",
                "product_converting_shares_price_mode",
				"product_agreed_converting_shares_price",
				"product_agreed_converting_shares_time",
                "product_converting_shares_term",
				"product_if_overdue",
				"product_overdue_info",
				"product_if_default",
				"product_default_start_time",
				"product_default_amount",
                "product_redemption",
                "product_issue_price",
                "product_face_value",
                "product_subscription_base",
                "product_successful_release_proportion",
                "product_fund_raising_conversion_condition",
                "product_is_make_over",
                "product_number_of_holders_max",
                "product_subscription_upper_limit",
                "product_subscription_lower_limit",
                "product_redemption_clause",
                "product_termination_conditions",
                "product_duration",
                "product_adjustment_change_control",
                "product_conversion_premium",
                "product_conversion_price_ref",
                "product_actual_issue_size",
                "product_raising_start_date",
                "product_raising_end_date",
                "product_start_date",
                "product_due_date",
                "product_amount_cashed",
                "product_first_interest_payment_date",
                "product_issuer_credit_rating",
                "product_credit_enhancement_agency_credit_rating",
                "product_guarantee_arrangement",
                "product_repo_arrangement",
                "product_lockup"
            ],
            "private_offering_fund": {
                "fund_basic_information":[
                    "product_raising_information_identification",
                    "product_scope_fund_raising",
                    "product_record_number",
                    "product_fund_filing_date",
                    "product_fund_type",
                    "product_foundation_date",
                    "product_escrow_bank",
                    "product_total_fund_share",
                    "product_fund_unit_holders_number",
                    "product_fund_nav",
                    "product_fund_fairvalue",
                    "product_raise_start_date",
                    "product_raise_end_date"
                ],
                "fund_sales_agency":"fund_sales_agency",
                "fund_manager_information":[
                    "product_fund_manager_name",
                    "product_fund_manager_certificate_number",
                    "product_management_style",
                    "product_funds_under_management_number",
                    "product_fund_management_scale"
                ]
            },
            "equity_issue_information": [
               "product_issue_scope",
               "product_issue_type",
               "product_shares_issued_class",
               "release_note_information",
               "product_issue_price",
               "product_issue_price_method",
               "product_before_authorized_shares",
               "product_after_authorized_shares",
               "product_after_issue_market_value",
               "product_net_profit",
               "product_annual_net_profit",
               "product_actual_raising_scale",
               "product_raising_start_time",
               "product_raising_end_time",
               "product_registered_capital_before_issuance",
               "product_registered_capital_issuance",
               "product_paid_shares",
               "product_shares_subscribed_number",
               "product_unlimited_sales_number_shares",
               "product_restricted_shares_number"
            ]


        },
        "transaction_information":{
            "transaction_status":["product_transaction_status"],
            "listing_information":[
                "product_transaction_scope",
                "product_transfer_permission_institution_to_individual",
                "product_transfer_lockup_days",
                "product_transfer_validity",
                "product_risk_level",
                "product_transaction_unit",
                "product_listing_code",
                "product_listing_date",
                "product_listing_remarks"
            ],
            "delisting_information":[
                "product_delisting_date",
                "product_delisting_type",
                "product_delisting_reason",
                "product_delisting_remarks"
            ]
        },
        "escrow_information":[
            "product_custodian_registration_date",
            "product_custodian_documents",
            "product_custodian_notes",
            "product_escrow_deregistration_date",
            "product_escrow_deregistration_document",
            "product_escrow_deregistration_remarks"
        ],
		"service_status_information":{
			"service_basic_information":[
				"product_service_status",
				"product_service_documents",
				"product_service_status_start_date",
				"product_service_status_end_date"
			],
			"product_terminate_situation":{
				"product_status_update_reason":"",
				"product_transfer_board_information":[
					"product_transfer_board_market"
				],
				"product_acquisition_information":[
					"product_acquisition_company_name",
					"product_acquisition_company_code",
					"product_acquisition_company_market"
				]
			}
		}
    }
}
`

	srcJson := `{
		"product_object_id": "29C66608B749FDF05B0885B2B3983A05",
		"product_trading_market_category": 1,
		"product_market_subject_ref": "D7FD749660A7F2E2F38AD05CF0DA4D6D",
		"product_market_subject_name": "上海股权托管交易中心股份有限公司",
		"service_provider_information": [],
		"product_file_information": [],
		"filing_information": [],
		"product_shares_issued_class": [],
		"product_plate_trading_name": "3",
		"product_issuer_subject_ref": "24E99F0CBEF42AFA16411772CFD65D9E",
		"product_issuer_name": "海南微派儿智慧城市建设有限公司",
		"product_code": "212162",
		"product_name": "海南微派儿智慧城市建设有限公司",
		"product_name_abbreviation": "微派儿",
		"product_type_function": 255,
		"product_listing_date": "2022-08-09",
		"product_type": 1,
		"product_account_number_max": 200,
		"product_info_disclosure_way": 0,
		"product_scale_unit": 1,
		"product_scale_currency": "156",
		"product_scale": 20000000,
		"product_customer_browsing_right": 0,
		"product_status_update_reason":0,
		"product_create_time": "2022-08-02T00:00:00+08:00"
	}`
	var srcMap map[string]interface{}
	err2 := json.Unmarshal([]byte(srcJson), &srcMap)
	if err2 != nil {
		panic(err2)
	}
	// t.Log(srcMap)
	result, err := genSuperviseData([]byte(jsonTemplate), srcMap)
	if err != nil {
		panic(err)
	}
	data, err3 := json.Marshal(result)
	if err3 != nil {
		panic(err3)
	}
	// t.Log("------------------------------")
	fmt.Println("------->>>>", string(data))
}

// genSuperviseData 生成监管报送数据.
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
