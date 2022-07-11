package main

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func main() {
	accessKey := "BQKfjq1-XwFwO5_yIMl_PLrFTSNRscc0z5nXQy3W"
	secretKey := "_9l2LNo1qP0iLp7qIHp3d2BdjiYTz07LzGwAEsm7"
	localFile := "E:\\20.06.16Project\\GOExample\\src\\A-20220516\\qiniuD\\uploadD\\506.jpg"
	bucket := "test_yida"
	key := "yida/506test.jpg"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 华东  storage.ZoneHuadong
	// 华北  storage.ZoneHuabei
	// 华南  storage.ZoneHuanan
	// 北美  storage.ZoneBeimei
	cfg.Zone = &storage.ZoneHuadong
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = true

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	fmt.Printf("%+v\n", ret)
}
