package utils

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
)

type UploadTool struct {
}

func (ut *UploadTool) UploadImage(imagePath string) (string, error) {
	uuidTool := UUIDTool{}
	accessKey := viper.GetString("upload.qiniu.accessKey")
	secretKey := viper.GetString("upload.qiniu.secretKey")
	bucket := viper.GetString("upload.qiniu.bucket")
	key := "imageRecord" + "_" + uuidTool.GenerateUUID()
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:     &storage.ZoneHuanan,
		UseHTTPS: false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "测试文件",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, imagePath, &putExtra)
	if err != nil {
		return "", err
	}
	domain := viper.GetString("upload.qiniu.domain")
	url := fmt.Sprintf("http://%s/%s", domain, key)
	/*err = os.Remove(imagePath)
	if err != nil {
		return "", err
	}*/
	return url, nil
}

func (ut *UploadTool) UploadVideo(videoPath string) (string, error) {
	uuidTool := UUIDTool{}
	accessKey := viper.GetString("upload.qiniu.accessKey")
	secretKey := viper.GetString("upload.qiniu.secretKey")
	bucket := viper.GetString("upload.qiniu.bucket")
	key := "videoRecord" + "_" + uuidTool.GenerateUUID()
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:     &storage.ZoneHuanan,
		UseHTTPS: false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "测试文件",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, videoPath, &putExtra)
	if err != nil {
		return "", err
	}
	domain := viper.GetString("upload.qiniu.domain")
	url := fmt.Sprintf("http://%s/%s", domain, key)
	/*err = os.Remove()
	if err != nil {
		return "", err
	}*/
	return url, nil
}
