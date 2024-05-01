package main

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func main() {
	accessKey := "2eAr9KFI_KIQD-QZnbbSiLZacIy-QG-tO_Yv0ONm"
	secretKey := "xl2Tj372xOOp_k5dclAKN-vlT4Pj9PuYKjM8wkR1"
	bucket := "sep-data"
	localFile := "E:\\university\\contest\\life\\Polyp-segmentation\\Polyp-PVT-main\\images\\3.png"
	key := "testPng1"

	// 创建上传策略
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 配置上传参数
	cfg := storage.Config{
		Zone:     &storage.ZoneHuadong, // 更改此处为华南的配置
		UseHTTPS: false,                // 默认使用HTTP，可以根据需要更改
	}
	cfg.Zone = &storage.ZoneHuanan
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "测试文件",
		},
	}

	// 执行文件上传
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println("上传失败:", err)
		return
	}

	// 输出上传成功的信息
	fmt.Printf("上传成功, 文件hash: %s, 文件key: %s\n", ret.Hash, ret.Key)

	// 生成文件的访问链接
	domain := "scsnadnl6.hn-bkt.clouddn.com" // 需要替换为实际的域名
	url := fmt.Sprintf("http://%s/%s", domain, key)
	fmt.Printf("文件访问链接: %s\n", url)
}
