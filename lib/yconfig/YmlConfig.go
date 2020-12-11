package yconfig

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/vhaoran/vchat/common/g"

	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/spf13/viper"
)

/*--auth: whr  date:2019/12/511:46--------------------------
 ####请勿擅改此功能代码####
 用途：
 系统中所有配置的入口，不引用其它任何非内部模块
--------------------------------------- */
type (
	YmlConfig struct {
		ES       ESConfig     `json:"es,omitempty"`
		WS       WSConfig     `json:"ws,omitempty"`
		RabbitMq RabbitConfig `json:"rabbitMq,omitempty"`
		Gateway  GWConfig     `json:"gwConfig,omitempty"`

		//微服务配置
		MicroService MicroServiceConfig `json:"microService,omitempty"`
		//etcd配置
		Etcd     ETCDConfig  `json:"etcd,omitempty"`
		Postgres PGConfig    `json:"postgres,omitempty"`
		Redis    RedisConfig `json:"redis,omitempty"`
		Emq      MQConfig    `json:"emq,omitempty"`
		Mongo    MongoConfig `json:"mongo,omitempty"`
		Log      LogConfig   `json:"log,omitempty"`
		Jwt      JwtConfig   `json:"jwt,omitempty"`
		Qiniu    QiniuConfig `json:"qiniu,omitempty"`
	}

	ETCDConfig struct {
		//ETCD cluster hosts
		Hosts   []string             `json:"hosts,omitempty"`
		Options etcdv3.ClientOptions `json:"options,omitempty"`
	}

	LogConfig struct {
		//debug/info/warn/error
		LogLevel string

		LogBackupPath string
		//日志路径
		LogPath string
		//文件名，不含路径
		FileName string
		//扩展名
		Ext string
	}

	//postgres-sql connection param
	PGConfig struct {
		// e.g: "host=%s user=%s dbname=%s sslmode=disable password=%s"
		URL     string `json:"url,omitempty"`
		PoolMax int    `json:"poolMax,omitempty"`
		PoolMin int    `json:"poolMin,omitempty"`
	}
	RedisConfig struct {
		Addrs []string `json:"addrs,omitempty"`

		// The maximum number of retries before giving up. Command is retried
		// on network errors and MOVED/ASK redirects.
		// Default is 8 retries.
		MaxRedirects int `json:"maxRedirects,omitempty"`

		// Enables read-only commands on slave nodes.
		ReadOnly bool `json:"readOnly,omitempty"`
		// Allows routing read-only commands to the closest master or slave node.
		// It automatically enables ReadOnly.
		RouteByLatency bool `json:"routeByLatency,omitempty"`
		// Allows routing read-only commands to the random master or slave node.
		// It automatically enables ReadOnly.
		RouteRandomly bool `json:"routeRandomly,omitempty"`

		Password string `json:"password,omitempty"`

		MaxRetries      int           `json:"maxRetries,omitempty"`
		MinRetryBackoff time.Duration `json:"minRetryBackoff,omitempty"`
		MaxRetryBackoff time.Duration `json:"maxRetryBackoff,omitempty"`

		DialTimeout  time.Duration `json:"dialTimeout,omitempty"`
		ReadTimeout  time.Duration `json:"readTimeout,omitempty"`
		WriteTimeout time.Duration `json:"writeTimeout,omitempty"`

		// PoolSize applies per cluster node and not for the whole cluster.
		PoolSize           int           `json:"poolSize,omitempty"`
		MinIdleConns       int           `json:"MinIdleConns,omitempty"`
		MaxConnAge         time.Duration `json:"maxConnAge,omitempty"`
		PoolTimeout        time.Duration `json:"poolTimeout,omitempty"`
		IdleTimeout        time.Duration `json:"idleTimeout,omitempty"`
		IdleCheckFrequency time.Duration `json:"idleCheckFrequency,omitempty"`
	}

	MQConfig struct {
		Url          string `json:"url,omitempty"`
		Host         string `json:"host,omitempty"`
		TCPPort      string `json:"tcpPort,omitempty"`
		UserName     string `json:"userName,omitempty"`
		Password     string `json:"password,omitempty"`
		MinOpenConnS int    `json:"monOpenConnS,omitempty"`
		MaxOpenConnS int    `json:"maxOpenConnS,omitempty"`
	}

	MongoConfig struct {
		URL     string        `json:"url,omitempty"`
		Options *MongoOptions `json:"options,omitempty"`
	}

	MongoOptions struct {
		AppName string `json:"appName,omitempty"`
		//Auth                   *Credential
		ConnectTimeout time.Duration `json:"connectTimeout,omitempty"`
		Compressors    []string      `json:"compressors,omitempty"`
		//Dialer                 ContextDialer
		HeartbeatInterval time.Duration `json:"heartbeatInterval,omitempty"`
		Hosts             []string      `json:"hosts,omitempty"`
		LocalThreshold    time.Duration `json:"localThreshold,omitempty"`
		MaxConnIdleTime   time.Duration `json:"maxConnIdleTime,omitempty"`
		MaxPoolSize       uint64        `json:"maxPoolSize,omitempty"`
		MinPoolSize       uint64        `json:"minPoolSize,omitempty"`
		//PoolMonitor            *event.PoolMonitor
		//Monitor                *event.CommandMonitor
		//ReadConcern            *readconcern.ReadConcern
		//ReadPreference         *readpref.ReadPref
		//Registry               *bsoncodec.Registry
		ReplicaSet             string        `json:"replicaSet ,omitempty"`
		RetryWrites            bool          `json:"retryWrites ,omitempty"`
		RetryReads             bool          `json:"retryReads ,omitempty"`
		ServerSelectionTimeout time.Duration `json:"serverSelectionTimeout,omitempty"`
		Direct                 bool          `json:"direct,omitempty"`
		SocketTimeout          time.Duration `json:"socketTimeout,omitempty"`
		// TLSConfig              tls.Config
		//WriteConcern           writeconcern.WriteConcern
		ZlibLevel int `json:"zlibLevel,omitempty"`
		//err error

		// Adds an option for internal use only and should not be set. This option is deprecated and is
		// not part of the stability guarantee. It may be removed in the future.
		AuthenticateToAnything bool `json:"authenticateToAnything,omitempty"`
	}

	RabbitConfig struct {
		Url     string `json:"url,omitempty"`
		PoolMax int    `json:"poolMax,omitempty"`
		PoolMin int    `json:"poolMin,omitempty"`
	}
)

func GetYmlConfig() (*YmlConfig, error) {
	var (
		pwd      string
		execPath string
		err      error
	)

	if pwd, err = os.Getwd(); err != nil {
		return nil, errors.New("get ymlConfig err," + err.Error())
	}
	l := []string{
		"../",
		"../../",
		"../../../",
		"../../../../",
		"../../../../../"}

	vp := viper.New()
	vp.AddConfigPath(pwd)
	for _, v := range l {
		vp.AddConfigPath(path.Join(pwd, v))
	}

	//-----------add execPath-----------------
	if execPath, err = g.GetExecPath(); err == nil {
		vp.AddConfigPath(execPath)
	}

	vp.SetConfigName("config")
	if fileName := os.Getenv("vchat_yml_file"); len(fileName) > 0 {
		fmt.Println("------------vchat_yml_file hitted-----", fileName)
		//pwd = s
		vp.SetConfigName(fileName)
	} else {
		fmt.Println("used default config name config.yml")
	}

	vp.SetConfigType("yml")
	yml := &YmlConfig{}
	if err = vp.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = vp.UnmarshalExact(yml); err != nil {
		return nil, err
	}
	return yml, nil
}
