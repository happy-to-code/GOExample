package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `{"state":200,"message":"success","data":{"subjectType":"natural","subjectId":"360202198807090038","totalItems":6,"errItems":1,"Types":{"asset":{"TotalItems":4,"ErrItems":1,"Items":[{"ItemName":"房屋权属","SourceHash":"e1853ba979a42426054c5c34444628ba36bc0dc43a9cbf502167402c92672e77","ItemValue":{"isSuccessful":"查询成功","fwlist":[{"xh":1,"qlrxxlist":[{"qlrid":"9d2ee8dc15c344438fb88911f475aabb","bdcqzh":"苏(2020)扬州市不动产权第0160892号","ybdcqzh":"苏(2017)扬州市不动产权第0129169号","dbsj":"2020-12-11.txt","gyqk":"共同共有"},{"qlrid":"c239c653369c4d8984ad7315a4d90245","bdcqzh":"苏(2020)扬州市不动产权第0160892号","ybdcqzh":"苏(2017)扬州市不动产权第0129169号","dbsj":"2020-12-11.txt","gyqk":"共同共有"}],"qszt":"现势","cfzt":"0","mj":"宗地面积0.00㎡/房屋建筑面积47.10㎡","fwzl":"江阳中路56号15幢103室","bdclx":"土地、房屋","fwyt":"城镇住宅用地/住宅"},{"xh":2,"qlrxxlist":[{"qlrid":"253270a357384267abfe362a1643fdec","bdcqzh":"苏(2017)扬州市不动产权第0042729号","ybdcqzh":"维共字2011003217/维字2011006076/扬国用(2012Z)第01357号","dbsj":"2017-04-21","gyqk":"共同共有"},{"qlrid":"79361c12e0254d7ca7769bc349705bf9","bdcqzh":"苏(2017)扬州市不动产权第0042728号","ybdcqzh":"维共字2011003217/维字2011006076/扬国用(2012Z)第01357号","dbsj":"2017-04-21","gyqk":"共同共有"}],"qszt":"现势","cfzt":"0","mj":"宗地面积0.00㎡/房屋建筑面积229.04㎡","fwzl":"邗江北路88号(锦苑)天元坊8幢603室","bdclx":"土地、房屋","fwyt":"城镇住宅用地/住宅"}]}},{"ItemName":"社保缴纳","SourceHash":"3d3255f03ed984af9e33901fe96822c24101fd3b882f52582f06c2ad871651f5","ItemValue":{"Name":"曾嵩","Total_month":76,"Continuous_total_month":1,"Latest_pay_date":"2021-02-09 00:00:00","Latest_pay_month":"202102","Records":[{"Pay_date":"2020-04-14 00:00:00","Pay_month":"202004","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-05-19 00:00:00","Pay_month":"202005","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-06-03 00:00:00","Pay_month":"202006","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-06-03 00:00:00","Pay_month":"202006","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-06-03 00:00:00","Pay_month":"202006","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-07-13 00:00:00","Pay_month":"202007","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-08-12 00:00:00","Pay_month":"202008","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-09-09 00:00:00","Pay_month":"202009","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-10-15 00:00:00","Pay_month":"202010","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-11.txt-06 00:00:00","Pay_month":"202011","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2020-12-07 00:00:00","Pay_month":"202012","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-01-12 00:00:00","Pay_month":"202101","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-02-09 00:00:00","Pay_month":"202102","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-03-09 00:00:00","Pay_month":"202103","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-03-09 00:00:00","Pay_month":"202103","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-03-09 00:00:00","Pay_month":"202103","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-03-09 00:00:00","Pay_month":"202103","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-03-09 00:00:00","Pay_month":"202103","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-04-08 00:00:00","Pay_month":"202104","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-04-08 00:00:00","Pay_month":"202104","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-04-08 00:00:00","Pay_month":"202104","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-04-08 00:00:00","Pay_month":"202104","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-05-10 00:00:00","Pay_month":"202105","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"},{"Pay_date":"2021-05-10 00:00:00","Pay_month":"202105","Base":7179,"Company":"扬州市人力资源与社会保障局","Type":"410"}]}},{"ItemName":"用水信息","SourceHash":"9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0","ItemValue":[{"hz":"曾嵩","yhdz":"锦苑二期天元坊8幢603","hh":"00328972","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2021年06月","syl":"44.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2021-07-01 11.txt:10:08","s_last_updatetime":"2021-07-01 11.txt:10:55","i_time":"2021-07-01 11.txt:10:55"},{"hz":"曾嵩","yhdz":"荷花池小区15栋103","hh":"00155436","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2021年06月","syl":"0.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2021-07-01 11.txt:10:08","s_last_updatetime":"2021-07-01 11.txt:57:02","i_time":"2021-07-01 11.txt:57:02"},{"hz":"曾嵩","yhdz":"锦苑二期天元坊8幢603","hh":"00328972","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2021年04月","syl":"41.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2021-05-01 10:32:59","s_last_updatetime":"2021-05-01 10:33:45","i_time":"2021-05-01 10:33:45"},{"hz":"曾嵩","yhdz":"荷花池小区15栋103","hh":"00155436","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2021年04月","syl":"3.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2021-05-01 10:32:59","s_last_updatetime":"2021-05-01 11.txt:21:47","i_time":"2021-05-01 11.txt:21:47"},{"hz":"曾嵩","yhdz":"锦苑二期天元坊8幢603","hh":"00328972","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2021年02月","syl":"40.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2021-03-01 09:25:34","s_last_updatetime":"2021-03-01 09:26:23","i_time":"2021-03-01 09:26:23"},{"hz":"曾嵩","yhdz":"荷花池小区15栋103","hh":"00155436","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2021年02月","syl":"0.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2021-03-01 09:25:34","s_last_updatetime":"2021-03-01 10:15:28","i_time":"2021-03-01 10:15:28"},{"hz":"曾嵩","yhdz":"锦苑二期天元坊8幢603","hh":"00328972","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2020年12月","syl":"36.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2021-01-01 09:25:34","s_last_updatetime":"2021-01-01 09:26:09","i_time":"2021-01-01 09:26:09"},{"hz":"季洁","yhdz":"荷花池小区15栋103","hh":"00155436","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2020年12月","syl":"5.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2021-01-01 09:25:34","s_last_updatetime":"2021-01-01 10:02:41","i_time":"2021-01-01 10:02:41"},{"hz":"曾嵩","yhdz":"锦苑二期天元坊8幢603","hh":"00328972","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2020年10月","syl":"47.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2020-11.txt-20 09:08:59","s_last_updatetime":"2020-11.txt-20 09:07:57","i_time":"2020-11.txt-20 09:07:57"},{"hz":"曾嵩","yhdz":"锦苑二期天元坊8幢603","hh":"00328972","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2020年08月","syl":"45.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2020-09-01 21:23:24","s_last_updatetime":"2020-09-01 21:37:40","i_time":"2020-09-01 21:37:40"},{"hz":"曾嵩","yhdz":"锦苑二期天元坊8幢603","hh":"00328972","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2020年06月","syl":"48.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2020-07-03 15:31:04","s_last_updatetime":"2020-07-03 15:44:38","i_time":"2020-07-03 15:44:38"},{"hz":"曾嵩","yhdz":"锦苑二期天元坊8幢603","hh":"00328972","hzzjlx":"","hzzjhm":"360202198807090038","lx":"用水","sd":"2019年02月","syl":"47.0","jldw":"吨","xxytbmqc":"扬州市自来水公司","tgdwqc":"扬州市自来水公司","tgrq":"2020-05-19 00:00:00","s_last_updatetime":"2020-05-19 14:24:12","i_time":"2020-05-19 14:24:12"}]}]},"identity":{"TotalItems":1,"ErrItems":0,"Items":[{"ItemName":"婚姻数据","SourceHash":"09a416247d52d4b28ea94acb5d49e92c104058d67b065ceea2184bd4310a478f","ItemValue":[{"name_man":"曾嵩","folk_man":"01","nation_man":"156","cert_num_man":"360202198807090038","name_woman":"吴雨萱","folk_woman":"01","nation_woman":"156","cert_num_woman":"321002198909090023","op_type":"IA","op_date":"2012-10-12 00:00:00"}]}]},"other":{"TotalItems":0,"ErrItems":0,"Items":null},"risk":{"TotalItems":1,"ErrItems":0,"Items":[{"ItemName":"严重失信","SourceHash":"300b14e7294805b0b33f3212bfa692a564d5ed7dc8337a429ad542c7364627ef","ItemValue":[]}]}}}}`

	var temp struct {
		Data json.RawMessage `json:"data"`
	}

	err := json.Unmarshal([]byte(s), &temp)
	if err != nil {
		panic(err)
	}
	fmt.Println(temp.Data)
	fmt.Println(string(temp.Data))

	var api APIRes
	err = json.Unmarshal(temp.Data, &api)
	if err != nil {
		panic(err)
	}

	fmt.Println("-----------------")
	fmt.Println("api==>", api)
	marshal, _ := json.Marshal(api)
	fmt.Println("api string:", string(marshal))
}

type APIRes struct {
	SubjectType string `json:"subjectType"`
	SubjectId   string `json:"subjectId"`
	TotalItems  uint8  `json:"totalItems"`
	ErrItems    uint8  `json:"errItems"`
	// Identity    *APIIdentityRes
	Types map[string]*List
}
type Item struct {
	ItemName   string
	SourceHash string
	ItemValue  interface{}
}

type List struct {
	TotalItems uint8
	ErrItems   uint8
	Items      []Item
}
