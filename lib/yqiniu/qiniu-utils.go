package yqiniu

import (
	"context"
	"strings"
	"time"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"

	"github.com/vhaoran/vchat/lib/yconfig"
	"github.com/vhaoran/vchat/lib/ylog"
)

var (
	accessKey      string = "gEpp05"
	secretKey      string = "EkZHh2"
	buck_permanent string = "permanent-wlkj"
	buck_temp      string = "temporary-wlkj"
)

func InitQiniu(cfg yconfig.QiniuConfig) {
	accessKey = cfg.AccessKey
	secretKey = cfg.SecretKey
	buck_permanent = cfg.BuckPermanent
	buck_temp = cfg.BuckTemp
}

func GetToken(hours int64) (string, error) {
	bucket := buck_permanent
	n := time.Now().Unix() + hours*3600

	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: uint64(n * 3600),
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken, nil
}

func GetTokenTemp(hours int64) (string, error) {
	bucket := buck_temp
	//- 持续化存储空间名: permanent-wlkj
	//--->  域名地址:p.0755yicai.com
	//- 临时存储空间名(7天):
	//--->  域名地址:t.0755yicai.com

	n := time.Now().Unix() + hours*3600

	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: uint64(n),
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken, nil
}

func GetVisitURL(key string, hours ...int64) string {
	mac := qbox.NewMac(accessKey, secretKey)
	//- 持续化存储空间名: permanent-wlkj
	//--->  域名地址:p.0755yicai.com
	//- 临时存储空间名(7天): temporary-wlkj
	//--->  域名地址:t.0755yicai.com

	//domain := "q52as9ix7.bkt.clouddn.com"
	domain := "p.0755yicai.com"
	n := int64(1)
	if len(hours) > 0 {
		n = hours[0]
	}

	deadline := time.Now().Add(time.Hour * time.Duration(n)).Unix()
	//n小时有效期

	//storage.MakePublicURL()
	url := storage.MakePrivateURL(mac, domain, key, deadline)
	//
	if strings.Index(url, "http") == -1 {
		url = "http://" + strings.TrimSpace(url)
	}
	return url
}

func UploadFileNoToken(localFile, keyForIdentify string) error {
	bucket := buck_permanent
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	//-------upload-------------------
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutFile(
		context.Background(),
		&ret,
		upToken,
		keyForIdentify,
		localFile,
		nil)
	if err != nil {
		ylog.Error("qiniu-utils.go->", err)
		return err
	}
	return nil
}
