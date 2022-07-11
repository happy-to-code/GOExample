package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

// MyPutRet 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
	Age    string
}

func main() {
	accessKey := "BQKfjq1-XwFwO5_yIMl_PLrFTSNRscc0z5nXQy3W"
	secretKey := "_9l2LNo1qP0iLp7qIHp3d2BdjiYTz07LzGwAEsm7"
	localFile := "E:\\20.06.16Project\\GOExample\\src\\A-20220516\\qiniuD\\uploadD\\506.jpg"
	bucket := "metadata-test"
	key := "yida1/506test.jpg"

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)","age":"$(x:age)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
			"x:age":  "16",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("-->%+v\n", ret)

	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, sErr := bucketManager.Stat(bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return
	}
	fmt.Println(fileInfo.String())

	// --------------------------------------------

	// 私有空间获取下载链接
	domain := "http://qn.yueda.vip"
	key = "h.txt"
	deadline := time.Now().Add(time.Second * 3600).Unix() // 1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	fmt.Println("---->", privateAccessURL)

}

func main2() {
	accessKey := "BQKfjq1-XwFwO5_yIMl_PLrFTSNRscc0z5nXQy3W"
	secretKey := "_9l2LNo1qP0iLp7qIHp3d2BdjiYTz07LzGwAEsm7"
	// localFile := "E:\\20.06.16Project\\GOExample\\src\\A-20220516\\qiniuD\\uploadD\\506.jpg"
	bucket := "test_yida"
	key := "yida1/1.txt"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = true
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	data := []byte("welcome to the hotel")
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)

}
