package main

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	httpClient *http.Client
)

// init HTTPClient
func init() {
	httpClient = createHTTPClient()
}

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

// createHTTPClient for connection re-use
func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout),
		},
	}
	return client
}

const (
	UrlPrefix = "http://172.21.33.7/"
)

type Param struct {
	Paperkind string `json:"paperkind"`
	Paperid   string `json:"paperid"`
	Orgnum    string `json:"orgnum"`
	Username  string `json:"username"`
}

var zf = `{
    "platformKeyId": "c3lp9vd2uehckgv6812g",
    "equityCode": "8008",
    "SubjectObjectId": "RTR004",
    "UniqueId": "%s",
    "shareList": [
        {
            "address": "SodNgpNjE99NxtfiUy4FwdxVspdUfPfVsRgxFVzqjqZV4jRkugQ",
            "amount":200,
            "shareProperty": 0,
            "registerInformation": {
                "register_registration_object_id": "%s",
                "register_object_information_type": 0,
                "register_registration_type": 1,
                "register_registration_serial_number": 12,
                "register_time": "2020/02/12 12:02:21",
                "register_rights_subject_ref": "OID123457",
                "register_rights_subject_type": 0,
                "register_account_obj_id": "OID123457",
                "registration_rights_type": 0,
                "registration_object_right": "oid123456",
                "register_unit": 0,
                "register_currency": 156,
                "register_rights_change_amount": 12,
                "register_available_balance": 120,
                "register_available_percentage": 0.54,
                "register_rights_pledge_change_amount": 12,
                "register_rights_pledge_balance": 41,
                "register_frozen_category": "质押冻结",
                "register_rights_frozen_change_amount": 12,
                "register_rights_frozen_balance": 13,
                "register_freeze_deadline_time": "2023/02/12 12:02:21",
                "register_holding_status": 1,
                "register_holding_attribute": [
                    "1.txt"
                ],
                "registration_source": 0,
                "register_source_type": 0,
                "register_notes": "",
                "register_verification_certificates": [
                    "1.txt"
                ],
                "transaction_type": 0,
                "register_transaction_obj_id": "%s",
                "register_roster_subject_ref": "sdfsddfgg",
                "register_rights_type": 0,
                "register_date": "2023/02/12 12:02:21",
                "register_shareholder_subject_ref": "fgrssgfs",
                "register_shareholder_subject_type": 0,
                "register_nature_of_shares": 0,
                "register_subscription_amount": 145,
                "register_paid_in_amount": 122,
                "register_shareholding_ratio": 0.12,
                "register_creditor_subject_ref": "fdfd",
                "register_creditor_type": 0,
                "register_creditor_subscription_count": 1560,
                "register_creditor_paid_in_amount": 56544,
                "register_creditor_contact_info": "dfgdfg"
            }
        },
        {
            "address": "SmtUeCCN13SASfmi3V6jKRAUgZs7K2nKSBfBPSXSrfQL9EQmB6y",
            "amount": 200,
            "shareProperty": 0,
            "registerInformation": {
                "register_registration_object_id": "%s",
                "register_object_information_type": 0,
                "register_registration_type": 1,
                "register_registration_serial_number": 12,
                "register_time": "2023/02/12 12:02:21",
                "register_rights_subject_ref": "OID123457",
                "register_rights_subject_type": 0,
                "register_account_obj_id": "OID123457",
                "registration_rights_type": 0,
                "registration_object_right": "oid123456",
                "register_unit": 0,
                "register_currency": 156,
                "register_rights_change_amount": 12,
                "register_available_balance": 120,
                "register_available_percentage": 0.54,
                "register_rights_pledge_change_amount": 12,
                "register_rights_pledge_balance": 41,
                "register_frozen_category": "质押冻结",
                "register_rights_frozen_change_amount": 12,
                "register_rights_frozen_balance": 13,
                "register_freeze_deadline_time": "2023/02/12 12:02:21",
                "register_holding_status": 1,
                "register_holding_attribute": [
                    "1.txt"
                ],
                "registration_source": 0,
                "register_source_type": 0,
                "register_notes": "",
                "register_verification_certificates": [
                    "1.txt"
                ],
                "transaction_type": 0,
                "register_transaction_obj_id": "oidjdjjd",
                "register_roster_subject_ref": "sdfsddfgg",
                "register_rights_type": 0,
                "register_date": "2023/02/12 12:02:21",
                "register_shareholder_subject_ref": "fgrssgfs",
                "register_shareholder_subject_type": 0,
                "register_nature_of_shares": 0,
                "register_subscription_amount": 145,
                "register_paid_in_amount": 122,
                "register_shareholding_ratio": 0.12,
                "register_creditor_subject_ref": "fdfd",
                "register_creditor_type": 0,
                "register_creditor_subscription_count": 1560,
                "register_creditor_paid_in_amount": 56544,
                "register_creditor_contact_info": "dfgdfg"
            }
        }
    ],
    "reason": "股份分红",
    "equityProductInfo": {
        "product_object_id": "PID345678",
        "subject_object_information_type": 0,
        "product_trading_market_category": 0,
        "product_market_subject": "OID908855",
        "product_market_subject_name": "OID908665",
        "product_plate_trading_name": "OID908899",
        "product_issuer_subject_ref": "PID345678",
        "product_issuer_name": "tongji",
        "product_code": "100005",
        "product_name": "tea111",
        "product_abbreviation": "tea111",
        "product_use": 0,
        "product_type": 0,
        "product_term": 200,
        "product_info_disclosure_way": 0,
        "product_scale_unit": 0,
        "product_scale_currency": 156,
        "product_scale": 10000,
        "product_customer_browsing_right": 0,
        "product_customer_trading_right": 0,
        "product_issuer_contact_person": "张三",
        "product_issuer_contact_info": "18966665555",
        "product_registry_subject_ref": "OID999999",
        "product_name_registration_body": "上海股交所",
        "product_trustee_subject_ref": "OID8900632",
        "product_name_trustee": "上海股交所",
        "product_underwriter_subject_ref": "OID8934354",
        "product_underwriter_name": "XXXX承销机构",
        "product_sponsor_subject": "OID90345634",
        "product_sponsor_subject_name": "XXXX保荐机构",
        "product_fund_administrator_account": "OID23433434",
        "product_fund_administrator_name": "张三",
        "product_guarantor_subject_ref": "OID34534534",
        "product_guarantor_name": "XXXX担保机构",
        "product_law_firm_subject_ref": "OID666666",
        "product_law_firm_name": "上海律师所",
        "product_accounting_firm_subject_ref": "OID555555",
        "product_accounting_firm_name": "上海会计所",
        "product_credit_enhancement_agency_subject_ref": "OID5599955",
        "product_credit_enhancement_agency_name": "XXXX增信机构",
        "product_credit_rating_agency_subject_ref": "OID5985659",
        "product_credit_rating_agency_name": "XXXX信用评级机构",
        "product_file_information": [
            {
                "product_issue_doc_num": 123,
                "product_file_name": "test",
                "product_issue_doc": "XXXX文件信息"
            }
        ],
        "product_registration_information": [
            {
                "product_license_type": 0,
                "product_license_number": 235,
                "product_license_file_name": "license",
                "product_license_documents": "XXXX产品许可文件"
            }
        ],
        "product_fund_type": [
            0,
            1
        ],
        "product_description _fund_use": "投资使用",
        "business_information": [
            {
                "product_business_purpose_name": "XXX使用",
                "product_business_purpose_details": "经营用途详情",
                "product_documents_list": [
                    "1.txt",
                    "2.pdf"
                ]
            }
        ],
        "portfolio_information": [
            {
                "product_investment_type": 0,
                "product_investment_proportion_range": "投资比例范围",
                "product_investment_details": "投资产品详情",
                "investment_documents_list": [
                    "2.doc",
                    "1.txt"
                ]
            }
        ],
        "product_issue_price": 10,
        "product_issue_price_method": "发行价格定价标准",
        "product_before_authorized_shares": 100000,
        "product_after_authorized_shares": 200000,
        "product_after_issue_market_value": 2000000,
        "product_net_profit": 150000,
        "product_annual_net_profit": 300000,
        "product_number_of_directional_issuers": 12,
        "product_issue_start_date": "203/06/05",
        "product_issue_end_date": "2023/10/05",
        "product_registration_date": "2023/06/12",
        "product_unlimited_sales_number_shares": 266000,
        "restricted_shares_number": 50000,
        "product_listing_code": "",
        "product_listing_time": "2023/01/01",
        "product_listing_status": 0,
        "product_listing_remarks": "挂牌备注信息",
        "product_delisting_time": "",
        "product_delisting_reason": ""
    }
}`

var zfs = fmt.Sprintf(zf, createUUID(), createUUID(), createUUID(), createUUID())

var zr = `{
    "keyId": "c3lp9vd2uehckgv68140",
    "UniqueId": "%s",
    "fromAddress": "SodNgpNjE99NxtfiUy4FwdxVspdUfPfVsRgxFVzqjqZV4jRkugQ",
    "amount": 10,
    "toAddress": "SoLJJvYDzqCfcvxmMVXywU1PKWNWsxye5sxBaSnsXZFBPq3wvoB",
    "shareProperty": 0,
    "equityCode": "8008",
    "txInformation": {
        "transaction_object_id": "%s",
        "transaction_object_information_type": 0,
        "transaction_traded_product_ref": "oid890032",
        "transaction_product_name": "sd",
        "transaction_type": 0,
        "transaction_description": "hj",
        "transaction_serial_num": 23,
        "transaction_close_method": 0,
        "transaction_close_currency": 156,
        "transaction_close_price": 1456,
        "transaction_close_amount": 23565,
        "transaction_close_time": "2021/02/12 12:02:21",
        "transaction_close_description": "fd",
        "transaction_Issuer_principal_subject_ref": "oid1234",
        "transaction_issuer_name": "df",
        "transaction_Investor_subject_ref": "df",
        "transaction_Investor_name": "hj",
        "transaction_original_owner_subject_ref": "hjg",
        "transaction_original_owner_name": "gh",
        "transaction_counterparty_subject_ref": "gh",
        "transaction_counterparty_name": "hgd",
        "transaction_order_verification_certificates": [
            "18.txt"
        ],
        "transaction_close_verification_certificates": [
            "1.txt"
        ],
        "transaction_intermediary_type": 0,
        "transaction_intermediary_subject_ref": "oid12421",
        "transaction_intermediary_name": "gf"
    },
    "fromRegisterInformation": {
        "register_object_type": 2,
        "register_creditors": [],
        "fund_investors": [],
        "register_shareholders": [],
        "register_registration_object_id": "rid000123",
        "register_object_information_type": 0,
        "register_registration_type": 1,
        "register_registration_serial_number": 12,
        "register_time": "2020/02/12 12:02:21",
        "register_rights_subject_ref": "OID123457",
        "register_rights_subject_type": 0,
        "register_account_obj_id": "OID123457",
        "registration_rights_type": 0,
        "registration_object_right": "oid123456",
        "register_unit": 0,
        "register_currency": 156,
        "register_rights_change_amount": 12,
        "register_available_balance": 120,
        "register_available_percentage": 0.54,
        "register_rights_pledge_change_amount": 12,
        "register_rights_pledge_balance": 41,
        "register_frozen_category": "质押冻结",
        "register_rights_frozen_change_amount": 12,
        "register_rights_frozen_balance": 13,
        "register_freeze_deadline_time": "2021/02/12 12:02:21",
        "register_holding_status": 1,
        "register_holding_attribute": [
            "1.txt"
        ],
        "registration_source": 0,
        "register_source_type": 0,
        "register_notes": "",
        "register_verification_certificates": [
            "1.txt"
        ],
        "transaction_type": 0,
        "register_transaction_obj_id": "oidjdjjd",
        "register_roster_subject_ref": "sdfsddfgg",
        "register_rights_type": 0,
        "register_date": "2019/02/12 12:02:21",
        "register_shareholder_subject_ref": "fgrssgfs",
        "register_shareholder_subject_type": 0,
        "register_nature_of_shares": 0,
        "register_subscription_amount": 145,
        "register_paid_in_amount": 122,
        "register_shareholding_ratio": 0.12,
        "register_creditor_subject_ref": "fdfd",
        "register_creditor_type": 0,
        "register_creditor_subscription_count": 1560,
        "register_creditor_paid_in_amount": 56544,
        "register_creditor_contact_info": "dfgdfg"
    },
    "toRegisterInformation": {
        "register_object_type": 2,
        "register_creditors": [],
        "fund_investors": [],
        "register_shareholders": [],
        "register_registration_object_id": "rid000123",
        "register_object_information_type": 0,
        "register_registration_type": 1,
        "register_registration_serial_number": 12,
        "register_time": "2020/02/12 12:02:21",
        "register_rights_subject_ref": "OID123457",
        "register_rights_subject_type": 0,
        "register_account_obj_id": "OID123457",
        "registration_rights_type": 0,
        "registration_object_right": "oid123456",
        "register_unit": 0,
        "register_currency": 156,
        "register_rights_change_amount": 12,
        "register_available_balance": 120,
        "register_available_percentage": 0.54,
        "register_rights_pledge_change_amount": 12,
        "register_rights_pledge_balance": 41,
        "register_frozen_category": "质押冻结",
        "register_rights_frozen_change_amount": 12,
        "register_rights_frozen_balance": 13,
        "register_freeze_deadline_time": "2021/02/12 12:02:21",
        "register_holding_status": 1,
        "register_holding_attribute": [
            "1.txt"
        ],
        "registration_source": 0,
        "register_source_type": 0,
        "register_notes": "",
        "register_verification_certificates": [
            "1.txt"
        ],
        "transaction_type": 0,
        "register_transaction_obj_id": "oidjdjjd",
        "register_roster_subject_ref": "sdfsddfgg",
        "register_rights_type": 0,
        "register_date": "2019/02/12 12:02:21",
        "register_shareholder_subject_ref": "fgrssgfs",
        "register_shareholder_subject_type": 0,
        "register_nature_of_shares": 0,
        "register_subscription_amount": 145,
        "register_paid_in_amount": 122,
        "register_shareholding_ratio": 0.12,
        "register_creditor_subject_ref": "fdfd",
        "register_creditor_type": 0,
        "register_creditor_subscription_count": 1560,
        "register_creditor_paid_in_amount": 56544,
        "register_creditor_contact_info": "dfgdfg"
    },
    "subjectObjectId": "RTR004"
}`

var zrs = fmt.Sprintf(zr, createUUID(), createUUID())

var shareholderList = `{"equityCode":"8008"}`

func main() {

	var contentType = "application/json"
	// 增发
	// api1 := "http://10.1.3.150:9100/equity/share/increase"
	// httpDo(httpClient, "POST", api1, contentType, zfs)

	// 股东列表
	api := "http://10.1.3.150:9100/equity/share/shareholder/list"
	holderListStr := httpDo(httpClient, "POST", api, contentType, shareholderList)
	change2Map(holderListStr)

	fmt.Println("===========================")
	// 转让
	api2 := "http://10.1.3.150:9100/equity/share/transfer"
	httpDo(httpClient, "POST", api2, contentType, zrs)

	holderListStr2 := httpDo(httpClient, "POST", api, contentType, shareholderList)
	change2Map(holderListStr2)
}

func change2Map(holderListStr string) map[string]int {
	var res Res
	json.Unmarshal([]byte(holderListStr), &res)

	m := make(map[string]int)
	data := res.Data
	for _, datum := range data {
		m[datum.Address] = datum.Amount
	}
	fmt.Printf("m==>%+v  %d\n", m, len(m))
	return m
}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, requestType, url, contentType, data string) string {
	req, err := http.NewRequest(requestType, url, strings.NewReader(data))
	if err != nil {
		return fmt.Sprintf("httpDo err:%v\n", err)
	}

	req.Header.Set("Content-Type", contentType)

	resp, e := client.Do(req)
	if e != nil {
		return fmt.Sprintf("接口调用出错：%v\n", e)
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		return fmt.Sprintf("httpDo er:%v\n", er)
	}
	fmt.Println("====>httpDo返回数据：", string(body))

	return string(body)
}

func createUUID() string {
	ul := uuid.NewV4()
	return strings.Replace(ul.String(), "-", "", -1)
}

type Res struct {
	State   int    `json:"state"`
	Message string `json:"message"`
	Data    []struct {
		Amount          int    `json:"amount"`
		LockAmount      int    `json:"lockAmount"`
		ShareProperty   int    `json:"shareProperty"`
		SharePropertyCN string `json:"sharePropertyCN"`
		Address         string `json:"address"`
	} `json:"data"`
}
