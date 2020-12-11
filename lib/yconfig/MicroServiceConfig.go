package yconfig

type (
	MicroServiceConfig struct {
		//微服務tag,想再的微服務tag相同
		Tag string

		//启动的本地商品
		Port int
		//启动的本地地址
		Host string

		//向etcd注册的商品，,在docker中时需要区分
		RegPort int
		//向etcd注册的host,在docker中时需要区分
		RegHost     string
		DispatchUrl string
		DispatchPwd string
	}
)
