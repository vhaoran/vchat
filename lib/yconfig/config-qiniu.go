package yconfig

type QiniuConfig struct {
	AccessKey     string `json:"accessKey,omitempty"`
	SecretKey     string `json:"secretKey,omitempty"`
	BuckPermanent string `json:"buckPermanent,omitempty"`
	BuckTemp      string `json:"buckTemp,omitempty"`
}
