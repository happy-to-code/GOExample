package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
	"time"
)

// UserWater 用户用水情况
type UserWater struct {
	HZ                string `json:"hz"`       // 户主
	YHDZ              string `json:"yhdz"`     // 用户地址
	HH                string `json:"hh"`       // 户号
	HZZJLX            string `json:"hzzjlx"`   // 户主证件类型
	HZZJHM            string `json:"hzzjhm"`   // 户主证件号码
	LX                string `json:"lx"`       // 类型
	SD                string `json:"sd"`       // 时段
	SYL               string `json:"syl"`      // 使用量
	JLDW              string `json:"jldw"`     // 计量单位
	XXYTBMQC          string `json:"xxytbmqc"` // 信息源头部门全称
	TGDWQC            string `json:"tgdwqc"`   // 提供单位全称
	TGRQ              string `json:"tgrq"`     // 提供日期
	S_LAST_UPDATETIME string `json:"S_LAST_UPDATETIME"`
	I_TIME            string `json:"I_TIME"`
}

// 数据库配置
const (
	userName = "renminyinhang"
	password = "renminyinhang"
	ip       = "172.21.47.17"
	port     = "3306"
	dbName   = "renminyinhang"
)

// DB Db数据库连接池
var DB *sql.DB

type User struct {
	id    int64
	name  string
	age   int8
	sex   int8
	phone string
}

// InitDB 注意方法名大写，就是public
func InitDB() {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
	// 打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	// 设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	// 验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

// Query 查询操作
func Query(sql string) ([]map[string]string, error) {
	// defer DB.Close()

	rows, e := DB.Query(sql)
	if e != nil {
		return nil, e
	}
	columns, _ := rows.Columns()            // 获取列的信息
	count := len(columns)                   // 列的数量
	var values = make([]interface{}, count) // 创建一个与列的数量相当的空接口
	for i := range values {
		var ii interface{} // 为空接口分配内存
		values[i] = &ii    // 取得这些内存的指针，因后继的Scan函数只接受指针
	}

	// 创建返回值：不定长的map类型切片
	ret := make([]map[string]string, 0)

	for rows.Next() {
		err := rows.Scan(values...) // 开始读行，Scan函数只接受指针变量
		if err != nil {
			return nil, err
		}

		m := make(map[string]string) // 用于存放1列的 [键/值] 对
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{})) // 读出raw数据，类型为byte
			typeStr := fmt.Sprintf("%s", reflect.TypeOf(raw_value))
			if typeStr == "time.Time" {
				t := raw_value.(time.Time)
				m[colName] = t.Format("2006-01-02 15:04:05") // colName是键，v是值
			} else {
				b, _ := raw_value.([]byte)
				v := string(b) // 将raw数据转换成字符串
				m[colName] = v // colName是键，v是值
			}
		}
		ret = append(ret, m) // 将单行所有列的键值对附加在总的返回值上（以行为单位）
	}
	rows.Close()
	fmt.Printf("======================>%+v\n", ret)
	return ret, nil
}

func DeleteUser(user User) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	// 准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	// 设置参数以及执行sql语句
	res, err := stmt.Exec(user.id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	// 提交事务
	tx.Commit()
	// 获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

func InsertUser(user User) bool {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	// 准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`name`, `phone`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	// 将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.name, user.phone)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	// 将事务提交
	tx.Commit()
	// 获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func main() {
	InitDB()

	sql := "select * from ysxx limit 2"
	ret, err := Query(sql)
	if err != nil {
		panic(err)
	}
	var userWaters []UserWater
	for _, v := range ret {
		var marshal, _ = json.Marshal(v)
		var uw UserWater
		err := json.Unmarshal(marshal, &uw)
		if err != nil {
			continue
		}
		userWaters = append(userWaters, uw)
	}
	fmt.Printf("222=========userWaters=============>%+v\n", userWaters)
	sql = "select * from ysxx limit 3"
	rets, err := Query(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println("rets==?", rets)
}
