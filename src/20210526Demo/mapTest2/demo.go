package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{
    "success": true,
    "qqsj": "2020-03-19 11:21:15:642",
    "page": 0,
    "size": 0,
    "qtime": 0,
    "records": 1,
    "total": 0,
    "cxbh": null,
    "jlid": null,
    "statusCode": 0,
    "message": null,
    "data": {
        "IsSuccessful": "查询成功",
        "fwlist": [
            {
                "xh": 1,
                "fttdmj": "24.04",
                "bz": "",
                "dytdmj": "",
                "qllx": "国有建设用地使用权/房屋所有权",
                "qszt": "现势",
                "sfzh": "",
                "yyxxlist": [],
                "qlrxxlist": [
                    {
                        "dbsj": "2015-12-29",
                        "bdcqzh": "房产证号:2015003236、",
                        "qlrzjh": "321084199104******",
                        "fj": "",
                        "qlrid": "QLR-180313151524-BKGWJRVZWK_201511209959",
                        "gyqk": "共同共有",
                        "gyfe": "",
                        "fzrq": "",
                        "ybdcqzh": "",
                        "qlr": "樊浩",
                        "qlbl": ""
                    },
                    {
                        "dbsj": "2015-12-29",
                        "bdcqzh": "房产证号:2015003236、",
                        "qlrzjh": "32108419910*******",
                        "fj": "",
                        "qlrid": "QLR-180313151524-GNBSQ3L787_201511209959",
                        "gyqk": "共同共有",
                        "gyfe": "",
                        "fzrq": "",
                        "ybdcqzh": "",
                        "qlr": "靳慧",
                        "qlbl": ""
                    }
                ],
                "cdsj": "2020-03-19 11:21:15:642",
                "xzzt": "抵押(限制状态以登记查询窗口为准)",
                "szc": "3",
                "dyaqxxlist": [
                    {
                        "dbsj": "2018-05-08",
                        "zwlxqx": "2014-01-28起2044-01-28止",
                        "dyaqlr": "中国农业银行股份有限公司高邮市支行",
                        "fzrq": "2018-05-08",
                        "zgzqse": "520000",
                        "qszt": "现势",
                        "dyfs": "最高额抵押",
                        "dysw": "1",
                        "bdbzqse": "52",
                        "dymj": "135.5",
                        "dyfj": "共有人：靳慧。\n该房已抵押给中国农业银行股份有限公司\n高邮市支行。",
                        "djyy": "房屋抵押补录",
                        "djzmh": "房产他项权证号：2015003219"
                    }
                ],
                "jzmj": "135.5",
                "qlxz": "市场化商品房",
                "qlr": "",
                "bdclx": "土地、房屋",
                "mj": "宗地面积2217.23㎡/房屋建筑面积135.50㎡",
                "fwzl": "盐河西路 水岸英伦花园16幢302室",
                "hth": "",
                "qlbsm": "",
                "dyqxxlist": [],
                "zyzxsj": "",
                "fjh": "302",
                "fwfzxxlist": [],
                "tdsyqx": "2046-06-01止",
                "cfzt": "0",
                "cfxxlist": [],
                "dyzt": "1",
                "cs": "12/3",
                "jznd": "",
                "jyjg": "",
                "djsj": "2015-12-29",
                "fwyt": "其他商服用地/住宅",
                "fwjg": "混合结构",
                "jgsj": "",
                "sdxxlist": [],
                "bdcdyh": "321084120211GB00086F00160302",
                "ygxxlist": [],
                "zcs": "12"
            }
        ],
        "records": 1,
        "IsABnormal": "1",
        "YWBH": "公开2020区号0700074367"
    }
}`

	var resp struct {
		Data struct {
			Fwlist json.RawMessage
		}
	}

	json.Unmarshal([]byte(s), &resp)
	fmt.Println("==>", string(resp.Data.Fwlist))

	// bytes, _ := json.Marshal(resp)
	// m := make(map[string]interface{})
	// json.Unmarshal(bytes, &m)
	// fmt.Println("------------")
	// fmt.Printf("%+v\n",m)

}
