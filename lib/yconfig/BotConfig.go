package yconfig

type BotConfig struct {
	//多具bottoken
	Tokens []string `json:"tokens,omitempty"`
	//角色，craw(爬早)/trade(交易)/
	Role string `json:"role,omitempty"`
}
