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
	Paperkind string `1.json:"paperkind"`
	Paperid   string `1.json:"paperid"`
	Orgnum    string `1.json:"orgnum"`
	Username  string `1.json:"username"`
}

var qydj = `
{
    "contractAddress": "516c0ab15163cd816a76dac49c0f976404d11ff3e16aa3bd9ec849ce8f28ee85",
    "equityCode": "%s",
    "totalShares": 5000000,
    "subjectCreateTime": 1448841600,
    "productCreateTime": 1448841600,
    "enterpriseSubjectInfo": {
        "subject_object_id": "RTR004",
        "subject_main_administrative_region": 1,
        "subject_type": 1,
        "subject_create_time": "2015-12-17",
        "subject_qualification_information": [],
        "subject_company_name": "上海联讯科技信息股份有限公司",
        "subject_company_short_name": "联讯科技",
        "subject_organization_nature": 2,
        "subject_legal_type": 0,
        "subject_economic_type": 0,
        "subject_company_type": 0,
        "subject_scale_type": 4,
        "subject_high_technology_enterprise": 0,
        "subject_document_information": [
            {
                "type": 1,
                "code": "310115000447454"
            }
        ],
        "subject_registry_date": "2021-07-16",
        "subject_business_license": [],
        "subject_business_scope": "电脑及配件、电子元器件、五金交电的销售，计算机机房装修，计算机网络系统设计、开发，大楼智能系统设计，计算机及其网络系统的四技服务，计算机信息服务。【依法须经批准的项目，经相关部门批准后方可开展经营活动】",
        "subject_industry": 0,
        "subject_company_business": "楼宇自控系统、门禁系统、能源计量智能抄表系统、基于物联网技术的智慧社区综合管理平台系统的研发和销售",
        "subject_company_profile": "公司主要从主要从事楼宇设备监控管理系统、门禁管理系统、能源计量智能抄表系统的研发、销售，主要产品为楼宇设备自控系统、智能门禁系统、智能远程抄表系统，拥有高新技术企业证书，目前拥有楼宇网络现场设备监控管理智能控制系统，等5项实用新型专利以及联讯家庭安防软件V1.0等17项软件著作权。通过掌握先进的光纤、通用总线控制技术等研发的产品先后在多项重大工程项目上得到广泛的应用。公司先后实现相宜本草弱电智能化系统设计，物联网系统设计及音视频系统设计，保利地产项目进行BA、梯控、物业管理系统的设计等项目。",
        "subject_registered_capital": 5000000,
        "subject_registered_capital_currency": "156",
        "subject_paid_in_capital": 5000000,
        "subject_paid_in_capital_currency": "156",
        "subject_registered_address": "中国（上海）自由贸易试验区巨野路93、97号6层",
        "subject_province": "310000",
        "subject_city": "310100",
        "subject_district": "310115",
        "subject_office_address": "上海市浦东新区巨野路93、97号6层",
        "subject_contact_address": "上海市浦东新区巨野路93、97号6层",
        "subject_contact_number": "021-51350988",
        "subject_fax": "021-51350988",
        "subject_postal_code": "200135",
        "subject_mail_box": "shenwt@shanghailianxun.com",
        "subject_association_articles": [],
        "subject_shareholders_number": 0,
        "subject_approval_time": "2021-07-16T00:00:00+08:00",
        "subject_insured_number": 0,
        "subject_company_status": 0,
        "subject_company_status_deregistration_date": "2021-07-16",
        "subject_company_status_windingup_date": "2021-07-16",
        "subject_name_used_before": [],
        "leading_member_information": [
            {
                "subject_key_personnel_name": "李建超",
                "subject_key_personnel_type": 1,
                "subject_key_personnel_position": 0,
                "subject_key_personnel_appointment_start": "2015-06-08",
                "subject_key_personnel_appointment_end": "2018-06-07",
                "subject_key_personnel_nationality": "",
                "subject_document_type": 1,
                "subject_key_personnel_id": "310224196606180029",
                "subject_key_personnel_address": "上海市浦东新区川沙镇沪塘街122号",
                "subject_key_personnel_shareholding_ratio": 7.00,
                "subject_key_personnel_shareholding": 350000,
                "subject_key_personnel_contact": "18918180601"
            }
        ]
    },
    "equityProductInfo": {
        "product_object_id": "%s",
        "product_trading_market_category": 1,
        "product_market_subject_ref": "D7FD749660A7F2E2F38AD05CF0DA4D6D",
        "product_market_subject_name": "上海股权托管交易中心股份有限公司",
        "service_provider_information": [],
        "product_file_information": [],
        "filing_information": [],
        "product_shares_issued_class": [],
        "product_plate_trading_name": "科创",
        "product_issuer_subject_ref": "OID123456",
        "product_issuer_name": "上海联讯科技信息股份有限公司",
        "product_code": "300018",
        "product_name": "上海联讯科技信息股份有限公司",
        "product_name_abbreviation": "联讯科技",
        "product_type_function": 0,
        "product_type": 0,
        "product_account_number_max": 200,
        "product_info_disclosure_way": 0,
        "product_scale_unit": 0,
        "product_scale_currency": "156",
        "product_scale": 5000000,
        "product_customer_browsing_right": 0,
        "product_create_time": "2021-07-16T00:00:00+08:00"
    }
}
`

var account1 = `
{
  "shareholderInfo": {
    "accountInfo": {
      "account_number": "h0123555",
      "account_type": 1,
      "account_frozen_remark": "phone0001",
      "account_create_time": "2017-12-16T13:40:00+08:00",
      "account_opening_date": "2018-10-23",
      "account_frozen_applicant_name": "name0001",
      "account_thaw_date": "2018-10-23",
      "account_depository_ref": "ADR514",
      "account_status": 1,
      "account_purpose": [
        0,
        1
      ],
      "account_thaw_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_name": "name0001",
      "account_object_id": "%s",
      "account_frozen_date": "2018-10-23",
      "account_closing_date": "2018-10-23",
      "account_closing_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_frozen_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_contact_number": "phone0001",
      "account_establish_date": "2018-10-23",
      "account_closing_agent_name": "name0001",
      "account_opening_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_subject_ref": "No9006ZYc2",
      "account_thaw_remark": "escription0001",
      "account_thaw_applicant_name": "name0001",
      "account_closing_agent_contact_number": "phone0001"
    },
    "shareholderNo": "SHNo9006ZYc21",
    "createTime": 1608062425
  },
  "createTime": 1513374000,
  "contractAddress": "516c0ab15163cd816a76dac49c0f976404d11ff3e16aa3bd9ec849ce8f28ee85",
  "clientNo": "No9006ZYc2",
  "fundInfo": {
    "accountInfo": {
      "account_number": "h0123555",
      "account_type": 2,
      "account_frozen_remark": "phone0001",
      "account_create_time": "2017-12-16T13:40:00+08:00",
      "account_opening_date": "2018-10-23",
      "account_frozen_applicant_name": "name0001",
      "account_thaw_date": "2018-10-23",
      "account_depository_ref": "ADR514",
      "account_status": 1,
      "account_purpose": [
        0,
        1
      ],
      "account_thaw_doc": [
        {
          "file_number": "2",
          "summary": "gjhfgj",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ghfgh",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_name": "name0001",
      "account_object_id": "%s",
      "account_frozen_date": "2018-10-23",
      "account_closing_date": "2018-10-23",
      "account_closing_doc": [
        {
          "file_number": "2",
          "summary": "ghgh1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "fgfdg",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_frozen_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_contact_number": "phone0001",
      "account_establish_date": "2018-10-23",
      "account_associated_account_ref": "OID88888",
      "account_closing_agent_name": "name0001",
      "account_opening_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_subject_ref": "No9006ZYc2",
      "account_thaw_remark": "escription0001",
      "account_thaw_applicant_name": "name0001",
      "account_closing_agent_contact_number": "phone0001",
      "account_association": 0
    },
    "fundNo": "fundNo9006ZYc22",
    "createTime": 1608062425
  },
  "investorInfo": {
    "subject_object_id": "No9006ZYc21",
    "subject_investor_name": "zhangsan",
    "subject_create_time": "2017-12-16T13:40:00+08:00",
    "subject_work_unit": "rtrtrt",
    "subject_personal_fax": "fgdfgdfgdfgd",
    "subject_gender": 1,
    "subject_occupation": 2,
    "subject_id_type": 0,
    "subject_cellphone_number": "ٶɋ˖ܺۅCHARACTER",
    "subject_birthday": "2018-10-23",
    "subject_contact_number": "ٶɋjϵ֧۰CHARACTER",
    "subject_city": "fgdfgdf",
    "subject_type": 2,
    "subject_id_address": "CHARACTER",
    "subject_investment_period": "12",
    "subject_native_place": "dfgdfgd",
    "subject_education": 1,
    "subject_postal_code": "Ԋ־ҠëCHARACTER",
    "subject_province": "ʡ؝CHARACTER",
    "subject_industry": 1,
    "subject_id_doc_mailbox": "ARACTER",
    "subject_main_administrative_region": 0,
    "subject_qualification_information": [
      {
        "investor_suitability_information": [
          {
            "subject_investor_qualification": 1,
            "subject_investor_qualification_certificate_time": "2017-12-16T13:40:00+08:00",
            "subject_investor_qualification_status": true,
            "subject_investor_qualification_certifier_ref": "SIQCR514",
            "subject_investor_qualification_description": "fgdfgdfgdfg",
            "subject_investor_qualification_certifier_name": "trtrtrt",
            "subject_investor_qualification_sub": "fgdf",
			"subject_investor_qualification_certificate": [
{
"file_number": "2",
"summary": "ݲ˶1",
"file_name": "1file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52221",
"url": "http://test.com/file/201/1file1.pdf",
"term_of_validity_type": "0"
},
{
"file_number": "2",
"summary": "ݲ˶2",
"file_name": "2file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52222",
"url": "http://test.com/file/201/2file1.pdf",
"term_of_validity_type": "0"
}
]
}
],
"qualification_authentication_information": [
{
"subject_qualification_reviewer": "ʳۋ׽",
"subject_certification_time": "2017-12-16T13:40:00+08:00",
"subject_qualification_authenticator": "ɏ֤׽",
"subject_review_time": "2017-12-16T13:40:00+08:00",
"subject_role_qualification_certification_doc": [
{
"file_number": "2",
"summary": "ݲ˶1",
"file_name": "1file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52221",
"url": "http://test.com/file/201/1file1.pdf",
"term_of_validity_type": "0"
},
{
"file_number": "2",
"summary": "ݲ˶2",
"file_name": "2file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52222",
"url": "http://test.com/file/201/2file1.pdf",
"term_of_validity_type": "0"
}
],
"subject_qualification_code": "؊׊պë",
"subject_qualification_status": true
}
],
"subject_market_roles_type": 1,
"subject_intermediary_qualification": [
0,
1
],
"subject_financial_qualification_type": 1,
"subject_qualification_category": 1
}
],
"subject_investment_experience": "Ͷ؊ޭzCHARACTER",
"subject_id_number":"hjhgj",
"subject_contact_address": "ɋjϵַ֘CHARACTER"
}
}
`

var account2 = `
{
  "shareholderInfo": {
    "accountInfo": {
      "account_number": "h0123555",
      "account_type": 1,
      "account_frozen_remark": "phone0001",
      "account_create_time": "2017-12-16T13:40:00+08:00",
      "account_opening_date": "2018-10-23",
      "account_frozen_applicant_name": "name0001",
      "account_thaw_date": "2018-10-23",
      "account_depository_ref": "ADR514",
      "account_status": 1,
      "account_purpose": [
        0,
        1
      ],
      "account_thaw_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_name": "name0001",
      "account_object_id": "%s",
      "account_frozen_date": "2018-10-23",
      "account_closing_date": "2018-10-23",
      "account_closing_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_frozen_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_contact_number": "phone0001",
      "account_establish_date": "2018-10-23",
      "account_closing_agent_name": "name0001",
      "account_opening_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_subject_ref": "No9006ZYc2",
      "account_thaw_remark": "escription0001",
      "account_thaw_applicant_name": "name0001",
      "account_closing_agent_contact_number": "phone0001"
    },
    "shareholderNo": "SHNo9006ZYc21",
    "createTime": 1608062425
  },
  "createTime": 1513374000,
  "contractAddress": "516c0ab15163cd816a76dac49c0f976404d11ff3e16aa3bd9ec849ce8f28ee85",
  "clientNo": "No9006ZYc2",
  "fundInfo": {
    "accountInfo": {
      "account_number": "h0123555",
      "account_type": 2,
      "account_frozen_remark": "phone0001",
      "account_create_time": "2017-12-16T13:40:00+08:00",
      "account_opening_date": "2018-10-23",
      "account_frozen_applicant_name": "name0001",
      "account_thaw_date": "2018-10-23",
      "account_depository_ref": "ADR514",
      "account_status": 1,
      "account_purpose": [
        0,
        1
      ],
      "account_thaw_doc": [
        {
          "file_number": "2",
          "summary": "gjhfgj",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ghfgh",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_name": "name0001",
      "account_object_id": "%s",
      "account_frozen_date": "2018-10-23",
      "account_closing_date": "2018-10-23",
      "account_closing_doc": [
        {
          "file_number": "2",
          "summary": "ghgh1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "fgfdg",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_frozen_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_contact_number": "phone0001",
      "account_establish_date": "2018-10-23",
      "account_associated_account_ref": "OID88888",
      "account_closing_agent_name": "name0001",
      "account_opening_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_subject_ref": "No9006ZYc2",
      "account_thaw_remark": "escription0001",
      "account_thaw_applicant_name": "name0001",
      "account_closing_agent_contact_number": "phone0001",
      "account_association": 0
    },
    "fundNo": "fundNo9006ZYc22",
    "createTime": 1608062425
  },
  "investorInfo": {
    "subject_object_id": "No9006ZYc21",
    "subject_investor_name": "zhangsan",
    "subject_create_time": "2017-12-16T13:40:00+08:00",
    "subject_work_unit": "rtrtrt",
    "subject_personal_fax": "fgdfgdfgdfgd",
    "subject_gender": 1,
    "subject_occupation": 2,
    "subject_id_type": 0,
    "subject_cellphone_number": "ٶɋ˖ܺۅCHARACTER",
    "subject_birthday": "2018-10-23",
    "subject_contact_number": "ٶɋjϵ֧۰CHARACTER",
    "subject_city": "fgdfgdf",
    "subject_type": 2,
    "subject_id_address": "CHARACTER",
    "subject_investment_period": "12",
    "subject_native_place": "dfgdfgd",
    "subject_education": 1,
    "subject_postal_code": "Ԋ־ҠëCHARACTER",
    "subject_province": "ʡ؝CHARACTER",
    "subject_industry": 1,
    "subject_id_doc_mailbox": "ARACTER",
    "subject_main_administrative_region": 0,
    "subject_qualification_information": [
      {
        "investor_suitability_information": [
          {
            "subject_investor_qualification": 1,
            "subject_investor_qualification_certificate_time": "2017-12-16T13:40:00+08:00",
            "subject_investor_qualification_status": true,
            "subject_investor_qualification_certifier_ref": "SIQCR514",
            "subject_investor_qualification_description": "fgdfgdfgdfg",
            "subject_investor_qualification_certifier_name": "trtrtrt",
            "subject_investor_qualification_sub": "fgdf",
			"subject_investor_qualification_certificate": [
{
"file_number": "2",
"summary": "ݲ˶1",
"file_name": "1file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52221",
"url": "http://test.com/file/201/1file1.pdf",
"term_of_validity_type": "0"
},
{
"file_number": "2",
"summary": "ݲ˶2",
"file_name": "2file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52222",
"url": "http://test.com/file/201/2file1.pdf",
"term_of_validity_type": "0"
}
]
}
],
"qualification_authentication_information": [
{
"subject_qualification_reviewer": "ʳۋ׽",
"subject_certification_time": "2017-12-16T13:40:00+08:00",
"subject_qualification_authenticator": "ɏ֤׽",
"subject_review_time": "2017-12-16T13:40:00+08:00",
"subject_role_qualification_certification_doc": [
{
"file_number": "2",
"summary": "ݲ˶1",
"file_name": "1file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52221",
"url": "http://test.com/file/201/1file1.pdf",
"term_of_validity_type": "0"
},
{
"file_number": "2",
"summary": "ݲ˶2",
"file_name": "2file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52222",
"url": "http://test.com/file/201/2file1.pdf",
"term_of_validity_type": "0"
}
],
"subject_qualification_code": "؊׊պë",
"subject_qualification_status": true
}
],
"subject_market_roles_type": 1,
"subject_intermediary_qualification": [
0,
1
],
"subject_financial_qualification_type": 1,
"subject_qualification_category": 1
}
],
"subject_investment_experience": "Ͷ؊ޭzCHARACTER",
"subject_id_number":"hjhgj",
"subject_contact_address": "ɋjϵַ֘CHARACTER"
}
}
`

var account3 = `
{
  "shareholderInfo": {
    "accountInfo": {
      "account_number": "h0123555",
      "account_type": 1,
      "account_frozen_remark": "phone0001",
      "account_create_time": "2017-12-16T13:40:00+08:00",
      "account_opening_date": "2018-10-23",
      "account_frozen_applicant_name": "name0001",
      "account_thaw_date": "2018-10-23",
      "account_depository_ref": "ADR514",
      "account_status": 1,
      "account_purpose": [
        0,
        1
      ],
      "account_thaw_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_name": "name0001",
      "account_object_id": "%s",
      "account_frozen_date": "2018-10-23",
      "account_closing_date": "2018-10-23",
      "account_closing_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_frozen_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_contact_number": "phone0001",
      "account_establish_date": "2018-10-23",
      "account_closing_agent_name": "name0001",
      "account_opening_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_subject_ref": "No9006ZYc2",
      "account_thaw_remark": "escription0001",
      "account_thaw_applicant_name": "name0001",
      "account_closing_agent_contact_number": "phone0001"
    },
    "shareholderNo": "SHNo9006ZYc21",
    "createTime": 1608062425
  },
  "createTime": 1513374000,
  "contractAddress": "516c0ab15163cd816a76dac49c0f976404d11ff3e16aa3bd9ec849ce8f28ee85",
  "clientNo": "No9006ZYc2",
  "fundInfo": {
    "accountInfo": {
      "account_number": "h0123555",
      "account_type": 2,
      "account_frozen_remark": "phone0001",
      "account_create_time": "2017-12-16T13:40:00+08:00",
      "account_opening_date": "2018-10-23",
      "account_frozen_applicant_name": "name0001",
      "account_thaw_date": "2018-10-23",
      "account_depository_ref": "ADR514",
      "account_status": 1,
      "account_purpose": [
        0,
        1
      ],
      "account_thaw_doc": [
        {
          "file_number": "2",
          "summary": "gjhfgj",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ghfgh",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_name": "name0001",
      "account_object_id": "%s",
      "account_frozen_date": "2018-10-23",
      "account_closing_date": "2018-10-23",
      "account_closing_doc": [
        {
          "file_number": "2",
          "summary": "ghgh1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "fgfdg",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_frozen_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_opening_agent_contact_number": "phone0001",
      "account_establish_date": "2018-10-23",
      "account_associated_account_ref": "OID88888",
      "account_closing_agent_name": "name0001",
      "account_opening_doc": [
        {
          "file_number": "2",
          "summary": "ݲ˶1",
          "file_name": "1file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52221",
          "url": "http://test.com/file/201/1file1.pdf",
          "term_of_validity_type": "0"
        },
        {
          "file_number": "2",
          "summary": "ݲ˶2",
          "file_name": "2file1.pdf",
          "term_of_validity": "2018-10-23",
          "hash": "da1234filehash52222",
          "url": "http://test.com/file/201/2file1.pdf",
          "term_of_validity_type": "0"
        }
      ],
      "account_subject_ref": "No9006ZYc2",
      "account_thaw_remark": "escription0001",
      "account_thaw_applicant_name": "name0001",
      "account_closing_agent_contact_number": "phone0001",
      "account_association": 0
    },
    "fundNo": "fundNo9006ZYc22",
    "createTime": 1608062425
  },
  "investorInfo": {
    "subject_object_id": "No9006ZYc21",
    "subject_investor_name": "zhangsan",
    "subject_create_time": "2017-12-16T13:40:00+08:00",
    "subject_work_unit": "rtrtrt",
    "subject_personal_fax": "fgdfgdfgdfgd",
    "subject_gender": 1,
    "subject_occupation": 2,
    "subject_id_type": 0,
    "subject_cellphone_number": "ٶɋ˖ܺۅCHARACTER",
    "subject_birthday": "2018-10-23",
    "subject_contact_number": "ٶɋjϵ֧۰CHARACTER",
    "subject_city": "fgdfgdf",
    "subject_type": 2,
    "subject_id_address": "CHARACTER",
    "subject_investment_period": "12",
    "subject_native_place": "dfgdfgd",
    "subject_education": 1,
    "subject_postal_code": "Ԋ־ҠëCHARACTER",
    "subject_province": "ʡ؝CHARACTER",
    "subject_industry": 1,
    "subject_id_doc_mailbox": "ARACTER",
    "subject_main_administrative_region": 0,
    "subject_qualification_information": [
      {
        "investor_suitability_information": [
          {
            "subject_investor_qualification": 1,
            "subject_investor_qualification_certificate_time": "2017-12-16T13:40:00+08:00",
            "subject_investor_qualification_status": true,
            "subject_investor_qualification_certifier_ref": "SIQCR514",
            "subject_investor_qualification_description": "fgdfgdfgdfg",
            "subject_investor_qualification_certifier_name": "trtrtrt",
            "subject_investor_qualification_sub": "fgdf",
			"subject_investor_qualification_certificate": [
{
"file_number": "2",
"summary": "ݲ˶1",
"file_name": "1file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52221",
"url": "http://test.com/file/201/1file1.pdf",
"term_of_validity_type": "0"
},
{
"file_number": "2",
"summary": "ݲ˶2",
"file_name": "2file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52222",
"url": "http://test.com/file/201/2file1.pdf",
"term_of_validity_type": "0"
}
]
}
],
"qualification_authentication_information": [
{
"subject_qualification_reviewer": "ʳۋ׽",
"subject_certification_time": "2017-12-16T13:40:00+08:00",
"subject_qualification_authenticator": "ɏ֤׽",
"subject_review_time": "2017-12-16T13:40:00+08:00",
"subject_role_qualification_certification_doc": [
{
"file_number": "2",
"summary": "ݲ˶1",
"file_name": "1file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52221",
"url": "http://test.com/file/201/1file1.pdf",
"term_of_validity_type": "0"
},
{
"file_number": "2",
"summary": "ݲ˶2",
"file_name": "2file1.pdf",
"term_of_validity": "2018-10-23",
"hash": "da1234filehash52222",
"url": "http://test.com/file/201/2file1.pdf",
"term_of_validity_type": "0"
}
],
"subject_qualification_code": "؊׊պë",
"subject_qualification_status": true
}
],
"subject_market_roles_type": 1,
"subject_intermediary_qualification": [
0,
1
],
"subject_financial_qualification_type": 1,
"subject_qualification_category": 1
}
],
"subject_investment_experience": "Ͷ؊ޭzCHARACTER",
"subject_id_number":"hjhgj",
"subject_contact_address": "ɋjϵַ֘CHARACTER"
}
}
`

var zr = `{
    "keyId": "%s",
    "UniqueId": "%s",
    "fromAddress": "%s",
    "amount": 10,
    "toAddress": "%s",
    "shareProperty": 0,
    "equityCode": "%s",
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

var csdj = `
{
    "platformKeyId": "c3lp9vd2uehckgv6812g",
    "equityCode": "%s",
    "CreateTime": 156,
    "temStore": false,
    "contractAddress": "516c0ab15163cd816a76dac49c0f976404d11ff3e16aa3bd9ec849ce8f28ee85",
    "UniqueId": "%s",
    "shareList": [
        {
            "address": "%s",
            "amount": 2000,
            "shareProperty": 0,
            "registerInformation": {
                "register_account_obj_id": "No000guAXq2lab9",
                "register_asset_type": 1,
                "register_thaw_description": "冻结/解冻说明信息TEXT",
                "register_registration_object_id": "%s",
                "register_pledge_balance_change": 1000,
                "register_shareholders": [
                    {
                        "register_equity_subject_type": 1,
                        "register_equity_capital": 200,
                        "register_equity_number": 20,
                        "register_equity_subject_ref": "RSR004",
                        "register_equity_type": 1,
                        "register_equity_shareholding": 10,
                        "register_equity_capital_paidin": 30000
                    }
                ],
                "register_subject_account_ref": "RSR004",
                "register_product_description": "产品描33述",
                "register_right_recognition_subject_name": "确权方主体名称CHARACTER",
                "register_nature_of_shares": 1,
                "register_subject_ref": "RSR004",
                "register_source_type": 0,
                "register_frozen_balance": 1000,
                "register_right_recognition_description": "确权描述信息CHARACTER",
                "register_asset_currency": "156",
                "register_asset_holding_status": 1,
                "register_authentic_right_recognition_status": 1,
                "register_creditors": [
                    {
                        "register_debt_holder_subscription_quantity": 1000,
                        "register_debt_holder_ref": "RSR004",
                        "register_debt_holder_subscription_price": 10000,
                        "register_debt_holder_contact_number": "债权人联系方式CHARACTER",
                        "register_debt_holder_type": 1
                    }
                ],
                "register_subject_type": 1,
                "register_pledge_balance_after": 5000,
                "register_asset_note": "登记说明TEXT",
                "register_frozen_balance_change": 500,
                "register_transaction_ref": "RTR004",
                "register_right_recognition_agent_subject_name": "确权代理方主体名称CHARACTER",
                "register_asset_unit": 1,
                "register_right_recognition_subject_ref": "RTR004",
                "register_serial_number": "登记流水号CHARACTER",
                "register_time": "2020-12-16T12:00:25+08:00",
                "register_right_recognition_mode": 1,
                "register_list_date": "2025-12-16",
                "register_thaw_doc": [
                    {
                        "file_number": "2",
                        "summary": "简述1",
                        "file_name": "1file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52221",
                        "url": "http://test.com/file/201/1file1.pdf",
                        "term_of_validity_type": "0"
                    },
                    {
                        "file_number": "2",
                        "summary": "简述2",
                        "file_name": "2file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52222",
                        "url": "http://test.com/file/201/2file1.pdf",
                        "term_of_validity_type": "0"
                    }
                ],
                "register_pledge_balance_before": 1000,
                "register_asset_holding_status_description": "持有状态说明TEXT",
                "register_event_type": 1,
                "register_product_ref": "RTR004",
                "register_right_recognition_doc": [
                    {
                        "file_number": "2",
                        "summary": "简述1",
                        "file_name": "1file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52221",
                        "url": "http://test.com/file/201/1file1.pdf",
                        "term_of_validity_type": "0"
                    },
                    {
                        "file_number": "2",
                        "summary": "简述2",
                        "file_name": "2file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52222",
                        "url": "http://test.com/file/201/2file1.pdf",
                        "term_of_validity_type": "0"
                    }
                ],
                "register_authentic_right_recognition_date": "2025-12-16",
                "register_right_recognition_agent_subject_ref": "RSR004",
                "register_asset_balance_after": 3000,
                "register_asset_equity_type": 1,
                "register_product_name": "产品名称",
                "register_object_type": 0,
                "register_create_time": "2020-12-16T12:00:25+08:00",
                "fund_investors": [
                    {
                        "register_subscription_amount": 1000,
                        "register_subscription_number": 2000,
                        "register_investor_subject_ref": "RSR004",
                        "register_fund_investors_classification": 1,
                        "register_investor_name": "投资人主体名称CHARACTER"
                    }
                ],
                "register_asset_verification_doc": [
                    {
                        "file_number": "2",
                        "summary": "简述1",
                        "file_name": "1file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52221",
                        "url": "http://test.com/file/201/1file1.pdf",
                        "term_of_validity_type": "0"
                    },
                    {
                        "file_number": "2",
                        "summary": "简述2",
                        "file_name": "2file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52222",
                        "url": "http://test.com/file/201/2file1.pdf",
                        "term_of_validity_type": "0"
                    }
                ],
                "register_asset_holding_nature": 1,
                "register_asset_balance_change": 1000,
                "register_list_asset_type": 1,
                "register_frozen_balance_after": 5000,
                "register_description": "登记描述信息CHARACTER",
                "register_asset_balance_before": 2000
            }
        },
        {
            "address": "%s",
            "amount": 2000,
            "shareProperty": 0,
            "registerInformation": {
                "register_account_obj_id": "No000gu5AXq2lab9",
                "register_asset_type": 1,
                "register_thaw_description": "冻结/解冻说明信息TEXT",
                "register_registration_object_id": "%s",
                "register_pledge_balance_change": 1000,
                "register_shareholders": [
                    {
                        "register_equity_subject_type": 1,
                        "register_equity_capital": 200,
                        "register_equity_number": 20,
                        "register_equity_subject_ref": "RSR004",
                        "register_equity_type": 1,
                        "register_equity_shareholding": 10,
                        "register_equity_capital_paidin": 30000
                    }
                ],
                "register_subject_account_ref": "RSR004",
                "register_product_description": "产品描33述",
                "register_right_recognition_subject_name": "确权方主体名称CHARACTER",
                "register_nature_of_shares": 1,
                "register_subject_ref": "RSR004",
                "register_source_type": 0,
                "register_frozen_balance": 1000,
                "register_right_recognition_description": "确权描述信息CHARACTER",
                "register_asset_currency": "156",
                "register_asset_holding_status": 1,
                "register_authentic_right_recognition_status": 1,
                "register_creditors": [
                    {
                        "register_debt_holder_subscription_quantity": 1000,
                        "register_debt_holder_ref": "RSR004",
                        "register_debt_holder_subscription_price": 10000,
                        "register_debt_holder_contact_number": "债权人联系方式CHARACTER",
                        "register_debt_holder_type": 1
                    }
                ],
                "register_subject_type": 1,
                "register_pledge_balance_after": 5000,
                "register_asset_note": "登记说明TEXT",
                "register_frozen_balance_change": 500,
                "register_transaction_ref": "CCFC673754ED3579E33506104C8C8E414",
                "register_right_recognition_agent_subject_name": "确权代理方主体名称CHARACTER",
                "register_asset_unit": 1,
                "register_right_recognition_subject_ref": "RTR004",
                "register_serial_number": "登记流水号CHARACTER",
                "register_time": "2020-12-16T12:00:25+08:00",
                "register_right_recognition_mode": 1,
                "register_list_date": "2025-12-16",
                "register_thaw_doc": [
                    {
                        "file_number": "2",
                        "summary": "简述1",
                        "file_name": "1file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52221",
                        "url": "http://test.com/file/201/1file1.pdf",
                        "term_of_validity_type": "0"
                    },
                    {
                        "file_number": "2",
                        "summary": "简述2",
                        "file_name": "2file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52222",
                        "url": "http://test.com/file/201/2file1.pdf",
                        "term_of_validity_type": "0"
                    }
                ],
                "register_pledge_balance_before": 1000,
                "register_asset_holding_status_description": "持有状态说明TEXT",
                "register_event_type": 1,
                "register_product_ref": "RTR004",
                "register_right_recognition_doc": [
                    {
                        "file_number": "2",
                        "summary": "简述1",
                        "file_name": "1file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52221",
                        "url": "http://test.com/file/201/1file1.pdf",
                        "term_of_validity_type": "0"
                    },
                    {
                        "file_number": "2",
                        "summary": "简述2",
                        "file_name": "2file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52222",
                        "url": "http://test.com/file/201/2file1.pdf",
                        "term_of_validity_type": "0"
                    }
                ],
                "register_authentic_right_recognition_date": "2025-12-16",
                "register_right_recognition_agent_subject_ref": "RSR004",
                "register_asset_balance_after": 3000,
                "register_asset_equity_type": 1,
                "register_product_name": "产品名称",
                "register_object_type": 0,
                "register_create_time": "2020-12-16T12:00:25+08:00",
                "fund_investors": [
                    {
                        "register_subscription_amount": 1000,
                        "register_subscription_number": 2000,
                        "register_investor_subject_ref": "RSR004",
                        "register_fund_investors_classification": 1,
                        "register_investor_name": "投资人主体名称CHARACTER"
                    }
                ],
                "register_asset_verification_doc": [
                    {
                        "file_number": "2",
                        "summary": "简述1",
                        "file_name": "1file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52221",
                        "url": "http://test.com/file/201/1file1.pdf",
                        "term_of_validity_type": "0"
                    },
                    {
                        "file_number": "2",
                        "summary": "简述2",
                        "file_name": "2file1.pdf",
                        "term_of_validity": "2020/04/18",
                        "hash": "da1234filehash52222",
                        "url": "http://test.com/file/201/2file1.pdf",
                        "term_of_validity_type": "0"
                    }
                ],
                "register_asset_holding_nature": 1,
                "register_asset_balance_change": 1000,
                "register_list_asset_type": 1,
                "register_frozen_balance_after": 5000,
                "register_description": "登记描述信息CHARACTER",
                "register_asset_balance_before": 2000
            }
        }
    ]
}
`

var xzbg = `
{
    "platformkeyId": "c3lp9vd2uehckgv6812g",
    "address": "%s",
    "equityCode": "%s",
    "amount": 1000,
    "oldShareProperty": 0,
    "newShareProperty": 1,
    "UpdateTime":156,
    "UniqueId":"%s",
    "registerInformationList": [
        {
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
        {
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
        }
    ]
}
`

// var equity = "SH2194"
// var equity string

func main() {
	var contentType = "application/1.json"
	for i := 0; i < 1; i++ {
		equity := createUUID()
		fmt.Println("===============iii=================", i, "  ", equity)

		// 增发
		// api1 := "http://10.1.3.150:9100/equity/share/increase"
		// httpDo(httpClient, "POST", api1, contentType, zfs)

		// 企业登记
		qydjUrl := "http://10.1.3.150:9100/equity/enterprise/issue"
		qydjs := fmt.Sprintf(qydj, equity, createUUID())
		httpDo(httpClient, "POST", qydjUrl, contentType, qydjs)

		// 创建账号
		accountUrl := "http://10.1.3.150:9100/equity/account/create"
		account1Str := fmt.Sprintf(account1, createUUID(), createUUID())
		accountRes1 := httpDo(httpClient, "POST", accountUrl, contentType, account1Str)
		accountStruct1 := change2Account(accountRes1)
		fmt.Println("accountStruct1:", accountStruct1)

		account2Str := fmt.Sprintf(account2, createUUID(), createUUID())
		accountRes2 := httpDo(httpClient, "POST", accountUrl, contentType, account2Str)
		accountStruct2 := change2Account(accountRes2)
		fmt.Println("accountStruct2:", accountStruct2)

		account3Str := fmt.Sprintf(account3, createUUID(), createUUID())
		accountRes3 := httpDo(httpClient, "POST", accountUrl, contentType, account3Str)
		accountStruct3 := change2Account(accountRes3)
		fmt.Println("accountStruct3:", accountStruct3)

		// 初始登记
		csdjUrl := "http://10.1.3.150:9100/equity/share/issue"
		csdjs := fmt.Sprintf(csdj, equity, createUUID(), accountStruct1.Data.AccountList.Address, createUUID(),
			accountStruct2.Data.AccountList.Address, createUUID())
		httpDo(httpClient, "POST", csdjUrl, contentType, csdjs)

		shareholderList := fmt.Sprintf(`{"equityCode":"%s"}`, equity)
		fmt.Println("shareholderList===>", shareholderList)
		// time.Sleep(time.Second * 3)
		// // 股东列表
		// api := "http://10.1.3.150:9100/equity/share/shareholder/list"
		// holderListStr := httpDo(httpClient, "POST", api, contentType, shareholderList)
		// fmt.Println("股东列表：", holderListStr)

		time.Sleep(time.Second * 4)
		// 股份性质变更
		gfxzbgUrl := "http://10.1.3.150:9100/equity/share/change"
		xzbgs := fmt.Sprintf(xzbg, accountStruct1.Data.AccountList.Address, equity, createUUID(), createUUID())
		xzbgRes := httpDo(httpClient, "POST", gfxzbgUrl, contentType, xzbgs)
		fmt.Println("xzbgRes===>", xzbgRes)

		// 股东列表
		api := "http://10.1.3.150:9100/equity/share/shareholder/list"
		holderListStr2 := httpDo(httpClient, "POST", api, contentType, shareholderList)
		fmt.Println("股东列表2：", holderListStr2)
		m2 := change2Map(holderListStr2)

		time.Sleep(time.Second * 3)
		// 转让
		api2 := "http://10.1.3.150:9100/equity/share/transfer"
		zrs := fmt.Sprintf(zr, accountStruct1.Data.AccountList.KeyId, createUUID(), accountStruct1.Data.AccountList.Address,
			accountStruct3.Data.AccountList.Address, equity, createUUID())
		transRes := httpDo(httpClient, "POST", api2, contentType, zrs)
		fmt.Println("转让：", transRes)

		// time.Sleep(time.Second * 5)
		// 股东列表
		api3 := "http://10.1.3.150:9100/equity/share/shareholder/list"
		holderListStr3 := httpDo(httpClient, "POST", api3, contentType, shareholderList)
		fmt.Println("股东列表3：", holderListStr3)
		m3 := change2Map(holderListStr3)
		if len(m3)-len(m2) != 1 {
			fmt.Println("--------------------出错--------------------出错")
			return
		}
		fmt.Println("----------------------------------------------------------------------------end")
	}
	// change2Map(holderListStr)
	//
	// fmt.Println("===========================")
	// // 转让
	// api2 := "http://10.1.3.150:9100/equity/share/transfer"
	// httpDo(httpClient, "POST", api2, contentType, zrs)
	//
	// holderListStr2 := httpDo(httpClient, "POST", api, contentType, shareholderList)
	// change2Map(holderListStr2)
}

func change2Map(holderListStr string) map[string]int {
	var res Res
	json.Unmarshal([]byte(holderListStr), &res)

	m := make(map[string]int)
	data := res.Data
	for _, datum := range data {
		m[datum.Address] = datum.Amount
	}
	fmt.Printf("m==>%+v \n %d\n", m, len(m))
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
	// fmt.Println("====>httpDo返回数据：", string(body))

	return string(body)
}

func createUUID() string {
	ul := uuid.NewV4()
	return strings.Replace(ul.String(), "-", "", -1)
}

type Res struct {
	State   int    `1.json:"state"`
	Message string `1.json:"message"`
	Data    []struct {
		Amount          int    `1.json:"amount"`
		LockAmount      int    `1.json:"lockAmount"`
		ShareProperty   int    `1.json:"shareProperty"`
		SharePropertyCN string `1.json:"sharePropertyCN"`
		Address         string `1.json:"address"`
	} `1.json:"data"`
}

type AccountRes struct {
	State   int    `1.json:"state"`
	Message string `1.json:"message"`
	Data    struct {
		AccountList struct {
			KeyId         string `1.json:"keyId"`
			Address       string `1.json:"address"`
			ClientNo      string `1.json:"clientNo"`
			ShareholderNo string `1.json:"shareholderNo"`
		} `1.json:"accountList"`
		TxId string `1.json:"txId"`
	} `1.json:"data"`
}

func change2Account(res string) AccountRes {
	var account AccountRes
	json.Unmarshal([]byte(res), &account)
	return account
}
