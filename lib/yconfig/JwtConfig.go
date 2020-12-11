package yconfig

type (
	JwtConfig struct {
		//安全key,所有微服务上必须相同
		SecretKey string `json:"secretKey,omitempty"`
		// 过期时间，计量单位:小時
		Expired int `json:"expired,omitempty"`
	}
)
