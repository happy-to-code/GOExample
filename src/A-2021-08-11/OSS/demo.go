package main

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

// var s = `{
//     "header": {
//         "model": {
//             "protocol": "区域性股权市场跨链业务数据模型",
//             "version": "2.0.0"
//         },
//         "content": {
//             "type": "subject",
//             "object_id": "RTR004",
//             "version": 0,
//             "operation": "create",
//             "timestamp": 1448841600
//         },
//         "metadata": null
//     },
//     "body": {
//         "subject_information": {
//             "basic_information_subject": {
//                 "general_information_subject": {
//                     "subject_create_time": "2015-12-17",
//                     "subject_main_administrative_region": 1,
//                     "subject_type": 0
//                 },
//                 "subject_qualification_information": []
//             },
//             "organization_subject_information": {
//                 "basic_information_of_enterprise": {
//                     "basic_information_description": {
//                         "subject_approval_time": "2021-07-16T00:00:00+08:00",
//                         "subject_association_articles": [],
//                         "subject_business_license": [],
//                         "subject_business_scope": "电脑及配件、电子元器件、五金交电的销售，计算机机房装修，计算机网络系统设计、开发，大楼智能系统设计，计算机及其网络系统的四技服务，计算机信息服务。【依法须经批准的项目，经相关部门批准后方可开展经营活动】",
//                         "subject_city": "310100",
//                         "subject_company_business": "楼宇自控系统、门禁系统、能源计量智能抄表系统、基于物联网技术的智慧社区综合管理平台系统的研发和销售",
//                         "subject_company_name": "上海联讯科技信息股份有限公司",
//                         "subject_company_profile": "公司主要从主要从事楼宇设备监控管理系统、门禁管理系统、能源计量智能抄表系统的研发、销售，主要产品为楼宇设备自控系统、智能门禁系统、智能远程抄表系统，拥有高新技术企业证书，目前拥有楼宇网络现场设备监控管理智能控制系统，等5项实用新型专利以及联讯家庭安防软件V1.0等17项软件著作权。通过掌握先进的光纤、通用总线控制技术等研发的产品先后在多项重大工程项目上得到广泛的应用。公司先后实现相宜本草弱电智能化系统设计，物联网系统设计及音视频系统设计，保利地产项目进行BA、梯控、物业管理系统的设计等项目。",
//                         "subject_company_short_name": "联讯科技",
//                         "subject_company_status": 0,
//                         "subject_company_status_deregistration_date": "2021-07-16",
//                         "subject_company_status_windingup_date": "2021-07-16",
//                         "subject_company_type": 0,
//                         "subject_contact_address": "上海市浦东新区巨野路93、97号6层",
//                         "subject_contact_number": "021-51350988",
//                         "subject_district": "310115",
//                         "subject_document_information": [
//                             {
//                                 "code": "310115000447454",
//                                 "type": 1
//                             }
//                         ],
//                         "subject_economic_type": 0,
//                         "subject_fax": "021-51350988",
//                         "subject_high_technology_enterprise": 0,
//                         "subject_industry": 0,
//                         "subject_insured_number": 0,
//                         "subject_legal_type": 0,
//                         "subject_mail_box": "shenwt@shanghailianxun.com",
//                         "subject_name_used_before": [],
//                         "subject_office_address": "上海市浦东新区巨野路93、97号6层",
//                         "subject_organization_nature": 2,
//                         "subject_paid_in_capital": 5000000,
//                         "subject_paid_in_capital_currency": "156",
//                         "subject_postal_code": "200135",
//                         "subject_province": "310000",
//                         "subject_registered_address": "中国（上海）自由贸易试验区巨野路93、97号6层",
//                         "subject_registered_capital": 5000000,
//                         "subject_registered_capital_currency": "156",
//                         "subject_registry_date": "2021-07-16",
//                         "subject_scale_type": 4,
//                         "subject_shareholders_number": 0
//                     },
//                     "leading_member_information": [
//                         {
//                             "subject_document_type": 1,
//                             "subject_key_personnel_address": "上海市浦东新区川沙镇沪塘街122号",
//                             "subject_key_personnel_appointment_end": "2018-06-07",
//                             "subject_key_personnel_appointment_start": "2015-06-08",
//                             "subject_key_personnel_contact": "18918180601",
//                             "subject_key_personnel_id": "310224196606180029",
//                             "subject_key_personnel_name": "李建超",
//                             "subject_key_personnel_nationality": "",
//                             "subject_key_personnel_position": 0,
//                             "subject_key_personnel_shareholding": 350000,
//                             "subject_key_personnel_shareholding_ratio": 7,
//                             "subject_key_personnel_type": 1
//                         }
//                     ]
//                 }
//             }
//         }
//     }
// }`

var s = `
{
    "body": {
        "registration_information": {
            "register_basic_infomation": {
				"register_event_type":0
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
}
`

func main() {
	// 上传一个zip文件。
	objectName := "25c3863df99e4ae09af2d99e3d6984c4/0"

	ctx := context.Background()
	endpoint := "10.1.5.168:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic(err)
	}

	// 创建一个叫mymusic的存储桶。
	bucketName := "zyf-test"

	// minioClient.
	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.PutObject(ctx, bucketName, objectName, bytes.NewReader([]byte(s)), int64(len([]byte(s))), minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s,%v\n", objectName, n)
	log.Printf("-------------------------")
}
