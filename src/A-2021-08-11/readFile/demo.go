package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

var (
	httpClient *http.Client
	sysType    = runtime.GOOS
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

func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

func getDataRepairPath() string {
	path := GetProjectPath()
	if sysType == "windows" {
		path = path + "\\dataRepair\\"
	} else if sysType == "linux" {
		path = path + "/dataRepair/"
	}
	return path
}

func makeHasFixedDataPath() string {
	path := GetProjectPath()
	if sysType == "windows" {
		path = path + "\\dataRepair\\hasFix\\"
	} else if sysType == "linux" {
		path = path + "/dataRepair/hasFix/"
	}
	return path
}

func makeErrFixedDataPath() string {
	path := GetProjectPath()
	if sysType == "windows" {
		path = path + "\\dataRepair\\fixErr\\"
	} else if sysType == "linux" {
		path = path + "/dataRepair/fixErr/"
	}
	return path
}

func makeHasFixedFile(timeStr string) string {
	// 拼装文件路径
	path := GetProjectPath()
	if sysType == "windows" {
		path = path + "\\dataRepair\\hasFix\\" + timeStr + "-hasFix.txt"
	} else if sysType == "linux" {
		path = path + "/dataRepair/hasFix/" + timeStr + "-hasFix.txt"
	}

	// 创建文件
	f, err2 := os.Create(path)
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	return path
}

func makeErrFixedFile(timeStr string) string {
	// 拼装文件路径
	path := GetProjectPath()
	if sysType == "windows" {
		path = path + "\\dataRepair\\fixErr\\" + timeStr + "-errFix.txt"
	} else if sysType == "linux" {
		path = path + "/dataRepair/fixErr/" + timeStr + "-errFix.txt"
	}

	// 创建文件
	f, err2 := os.Create(path)
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	return path
}

// 有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
func httpDo(client *http.Client, url, data string) (string, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("NewRequest err:%v\n", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, e := client.Do(req)
	if e != nil {
		return "", fmt.Errorf("接口调用出错：%v\n", e)
	}

	defer resp.Body.Close()

	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		return "", fmt.Errorf("ioutil.ReadAll err:%v\n", er)
	}

	// fmt.Println("1===============================1===========================1")
	// var reply Reply
	// 1.json.Unmarshal(body, &reply)
	// log.Printf("reply:%+v\n", reply)
	// fmt.Println("2===============================2===========================2")

	return string(body), nil
}

// PathExists 判断目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	// 数据存放目录
	dataStoreDir := getDataRepairPath()

	exist, err := PathExists(dataStoreDir)
	if err != nil {
		log.Printf("get dir error[%v]\n", err)
		return
	}

	if exist {
		log.Printf("has dir[%v]\n", dataStoreDir)
	} else {
		log.Printf("no dir[%v]\n", dataStoreDir)
		// 创建文件夹
		err := os.Mkdir(dataStoreDir, os.ModePerm)
		if err != nil {
			log.Printf("mkdir failed[%v]\n", err)
		} else {
			log.Printf("mkdir success\n")
		}
	}

	// 新建已经修复数据的文件夹
	hasFixedDataPath := makeHasFixedDataPath()
	isExist, _ := PathExists(hasFixedDataPath)
	if !isExist {
		os.Mkdir(hasFixedDataPath, os.ModePerm)
	}

	// 存放修复失败文件的文件夹
	errFixedDataPath := makeErrFixedDataPath()
	exists, _ := PathExists(errFixedDataPath)
	if !exists {
		os.Mkdir(errFixedDataPath, os.ModePerm)
	}

	// 初始化httpClient
	httpClient = createHTTPClient()
}

func writeResult(fileName, data string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		// 打开文件失败处理
		log.Fatalln("########", err)
		return err
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)
		data = data + "\n"

		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(data), n)
	}

	defer f.Close()
	return nil
}

func main() {
	// 当前时间字符串
	timeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02-15-04-05")
	// 新建已经修复成功的数据文件
	hasFixedFile := makeHasFixedFile(timeStr)

	// 新建已经修复成功的数据文件
	hasErrFixFile := makeErrFixedFile(timeStr)

	// http 要请求的url路径
	var url string
	var filePath string

	var defaultFilePath = ""
	if sysType == "windows" {
		defaultFilePath = "/fileData.txt"
	} else {
		defaultFilePath = "./fileData.txt"
	}
	flag.StringVar(&filePath, "path", defaultFilePath, "处理文件存放路径")
	flag.StringVar(&url, "url", "http://localhost:9100/equity/object/repair", "http请求路径")
	flag.Parse()

	// 读取文件，然后进行数据处理
	fi, err := os.Open(filePath)
	if err != nil {
		log.Printf("打开文件出错,Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		sLine := string(a)
		// 将字符串转换成对象
		obj := convertToObj(sLine)

		objBytes, _ := json.Marshal(obj)
		// 调用http请求
		httpResponse, err := httpDo(httpClient, url, string(objBytes))
		if err != nil {
			sLine = sLine + fmt.Sprintf("%s%s%s", " [", err, "]")
			writeResult(hasErrFixFile, sLine)
		} else {
			// http 返回值
			var reply Reply
			err := json.Unmarshal([]byte(httpResponse), &reply)
			if err != nil {
				sLine = sLine + " [" + fmt.Sprintf("JSON发序列化出错:%s", err) + "]"
				writeResult(hasErrFixFile, sLine)
			} else {
				if reply.State != http.StatusOK {
					// 失败  400
					sLine = sLine + " [" + reply.Message + "]"
					writeResult(hasErrFixFile, sLine)
				} else {
					// 修复成功的数据
					sLine = sLine + fmt.Sprintf("%s%s%s", " [", reply.Data, "]")
					writeResult(hasFixedFile, sLine)
				}
			}
		}
	}
}

func convertToObj(sLine string) ObjectRepair {
	// 1,864310308f36426783d84afb1bfbae98,0,registration
	split := strings.Split(sLine, ",")
	indexInt, _ := strconv.Atoi(split[0])
	versionInt, _ := strconv.Atoi(split[2])

	return ObjectRepair{
		Index:   indexInt,
		Id:      split[1],
		Version: int64(versionInt),
		Type:    split[3],
	}

}

type ObjectRepair struct {
	Index   int
	Id      string // 对象标识
	Version int64  // 对象版本
	Type    string // 对象类型（subject、registration……）
}

type Reply struct {
	State   int         `1.json:"state"`
	Message string      `1.json:"message"`
	Data    interface{} `1.json:"data"`
}
