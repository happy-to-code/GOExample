package main

import (
	"fmt"
)

type Data_Item int

const (
	ASSET_NA_HOUSE Data_Item = iota + 1
	ASSET_NA_VEHICLE
	ASSET_NA_SS
	ASSET_NA_WATER
	ASSET_NA_GAS
	Asset_NA_END
)

const (
	IDENT_M Data_Item = iota + 20
	IDENT_PERSONAL
)

const (
	// 严重失信
	RISK_NA_SX Data_Item = iota + 40
)

// 法人资产
const (
	LEGAL_SB       Data_Item = iota + 50 // 法人社保缴纳信息
	LEGAL_SBQJF                          // 社保欠缴费
	LEGAL_FUND                           // 法人公积金缴存（对接中）
	LEGAL_GASQJF                         // 燃气欠缴费
	LEGAL_TAX_INFO                       // 税务信息（对接中)
	LEGAL_SPAQ                           // 食品生产经营许可证信息
	LEGAL_YPJY                           // 药品生产经营许可证信息
	LEGAL_YLQC                           // 医疗器械生产许可证信息

)

type Query struct {
	policePrefix string // 公安查询前缀
	prefix       string // 社保、房屋、婚姻前缀
	fw           string // 覆盖范围
	fwgy         string // 高邮
	fwyz         string // 仪征
	hy           string // 婚姻
	sb           string // 社保
	b            string // 自然人黑名单
	w            string // 自然人白名单
	c            string // 车辆
	p            string // 人员信息

	gt   string // 个体在业
	qy   string // 企业在业
	spsc string // 食品生产
	ypjy string // 药品经营
	ylqx string // 医疗器械

	frsb    string // 社会法人社保缴费信息
	frsbqjx string // 社会法人社保欠缴信息

	assetNatural map[Data_Item]QueryFunc
	ident        map[Data_Item]QueryFunc
	risk         map[Data_Item]QueryFunc

	assetLegal map[Data_Item]QueryFunc
	identLegal map[Data_Item]QueryFunc
	riskLegal  map[Data_Item]QueryFunc
}
type Item struct {
	ItemName   string      `json:"itemName"`
	SourceHash string      `json:"sourceHash"`
	ItemValue  interface{} `json:"itemValue"`
}
type QueryFunc func(string, string) (Item, error)

func NewQueryForTest(id, name string, db bool) *Query {
	query := &Query{
		policePrefix: "http://172.21.33.7",  // 公安
		prefix:       "http://172.21.47.49", // -- 社保、婚姻、房屋

		fw:   "/gateway/api/2/fwqscx",      // 四区
		fwgy: "/gateway/api/1/gysfwqscxjk", // 高邮
		fwyz: "/gateway/api/1/yzsfwqscxjk", // 仪征

		hy:           "/gateway/api/1/hydjxxcx",        // 婚姻
		sb:           "/gateway/api/1/zrrsbjnxxcxjk",   // 自然人社保
		b:            "/gateway/api/1/zrryzsxhmdfy",    // 自然人严重失信黑名单
		w:            "/gateway/api/1/shfrcssxhmdfyfw", // 自然人守信红名单
		c:            "/gateway/api/1/cldjxx",          // 车辆登记信息查询接口
		p:            "/gateway/api/1/garyxx",          // 公安-人员信息
		gt:           "/gateway/api/1/gtzyxx",          // 个体在业
		qy:           "/gateway/api/1/QYZY",            // 企业在业
		spsc:         "/gateway/api/1/spscjyxkzcxjk",   // 食品生产
		ypjy:         "/gateway/api/1/ypjyxkzxxcx",     // 药品经营
		ylqx:         "/gateway/api/2/ylqxscxkzxx",     // 医疗器械
		assetNatural: make(map[Data_Item]QueryFunc),
		ident:        make(map[Data_Item]QueryFunc),
		risk:         make(map[Data_Item]QueryFunc),
	}

	query.register()

	// if db {
	// 	initDB()
	// }
	return query
}
func (query *Query) getWaterUse(id, _ string) (Item, error) {
	return Item{"用水信息", "", Item{}}, nil
}

func (query *Query) getGasUse(id, _ string) (Item, error) {
	return Item{"用气信息", "sourceHash", Item{}}, nil
}
func (query *Query) register() {
	query.assetNatural[ASSET_NA_WATER] = query.getWaterUse
	query.assetNatural[ASSET_NA_GAS] = query.getGasUse

}

func main() {

	query := NewQueryForTest("", "", true)
	for item, queryFunc := range query.assetNatural {
		fmt.Printf("item:%d,queryFunc:%v\n", item, &queryFunc)

	}

}
