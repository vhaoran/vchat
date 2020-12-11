package yconfig

type (
	WSConfig struct {
		Port        int    `json:"port,omitempty"`
		HttpSendPwd string `json:"httpSendPwd,omitempty"`
		WSSendPwd   string `json:"wsSendPwd,omitempty"`
	}
)
