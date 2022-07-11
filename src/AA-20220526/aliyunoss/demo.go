package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	// endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// // 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	// accessKeyId := "LTAIiAmZWkjVbaEG"
	// accessKeySecret := "l9LSX6VTGjwAxF84ILyffFVU7SQK96"
	// bucketName := "yida-test-pic"
	// // <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// objectName := "20211118144931.jpg"
	// // <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	// localFileName := "D:\\A-MyFile\\20211118144931.jpg"
	// // 创建OSSClient实例。
	// client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	// if err != nil {
	// 	handleError(err)
	// }
	// // 获取存储空间。
	// bucket, err := client.Bucket(bucketName)
	// if err != nil {
	// 	handleError(err)
	// }
	// // 上传文件。
	// err = bucket.PutObjectFromFile(objectName, localFileName)
	// if err != nil {
	// 	handleError(err)
	// }

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", "LTAIiAmZWkjVbaEG", "l9LSX6VTGjwAxF84ILyffFVU7SQK96")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("yida-test-pic")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile("10086.jpg", "D:\\A-MyFile\\20211118144931.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

}
