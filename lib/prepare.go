package lib

import (
	"github.com/vhaoran/vchat/common/ytime"
	_ "github.com/vhaoran/vchat/common/ytime"
	"github.com/vhaoran/vchat/lib/yconfig"
	"github.com/vhaoran/vchat/lib/yes"
	"github.com/vhaoran/vchat/lib/yetcd"
	"github.com/vhaoran/vchat/lib/yjwt"
	"github.com/vhaoran/vchat/lib/ylog"
	"github.com/vhaoran/vchat/lib/ymongo"
	"github.com/vhaoran/vchat/lib/ymq"
	"github.com/vhaoran/vchat/lib/ymqa"
	"github.com/vhaoran/vchat/lib/ypg"
	"github.com/vhaoran/vchat/lib/yqiniu"
	"github.com/vhaoran/vchat/lib/yred"
	"github.com/vhaoran/vchat/lib/yredis"
)

type LoadOption struct {
	LoadMicroService bool //0
	LoadEtcd         bool //1
	LoadPg           bool //2
	//load redis cluster
	LoadRedis bool //3
	//load redis single server
	LoadRed      bool
	LoadMongo    bool //4
	LoadMq       bool //5
	LoadRabbitMq bool //6
	LoadJwt      bool //7
	LoadES       bool //8
	LoadQiniu    bool //9

}

func InitModulesOfAll() (*yconfig.YmlConfig, error) {
	cfg := LoadOption{
		LoadMicroService: true,
		LoadEtcd:         true,
		LoadPg:           true,
		LoadRedis:        true,
		LoadRed:          true,
		LoadMongo:        true,
		LoadMq:           false,
		LoadRabbitMq:     true,
		LoadJwt:          true,
		LoadES:           true,
		LoadQiniu:        true,
	}

	return InitModulesOfOptions(&cfg)
}

func InitModulesOfOptions(opt *LoadOption) (*yconfig.YmlConfig, error) {
	var (
		cfg *yconfig.YmlConfig
		err error
	)
	//initOthers()

	if cfg, err = yconfig.GetYmlConfig(); err != nil {
		return nil, err
	}

	if err = ylog.InitLog(cfg.Log); err != nil {
		ylog.DebugDump(cfg.Log)
		return nil, err
	}

	//--------etcd -----------------------------
	// 微服务注册地址设置set XEtcdConfig
	if opt.LoadEtcd {
		ylog.Debug("etcd config init...", cfg.Etcd.Hosts)
		if err := yetcd.InitETCD(cfg.Etcd); err != nil {
			return nil, err
		}
		ylog.Debug("etcd config init ok")
	}

	//-------- postgres sql -----------------------------
	//postgres 数据库配置参数设置 X
	if opt.LoadPg {
		ylog.Debug("postgres connecting...", cfg.Postgres.URL)
		debug := (cfg.Log.LogLevel == "debug" || len(cfg.Log.LogLevel) == 0)
		if err := ypg.InitPG(cfg.Postgres, debug); err != nil {
			return nil, err
		}
		ylog.Debug("postgres connected ok")
	}

	//--------load redis -cluster----------------------------
	//redis cluster连接设置 xred
	if opt.LoadRedis {
		//set X
		ylog.Debug("redis cluster connecting...", cfg.Redis.Addrs)
		if err := yredis.InitRedis(cfg.Redis); err != nil {
			return nil, err
		}
		ylog.Debug("redis cluster  connected ok")
	}

	//--------load redis single server-----------------------------
	//redis cluster连接设置 xred
	if opt.LoadRed {
		//set X
		ylog.Debug("redis single server connecting...", cfg.Redis.Addrs)
		if err := yred.InitRedis(cfg.Redis); err != nil {
			return nil, err
		}
		ylog.Debug("redis single connected ok")
	}

	//--------load emq -----------------------------

	//--------load mongo -----------------------------
	if opt.LoadMongo {
		ylog.Debug("mongo connecting...", cfg.Mongo.URL)
		if err := ymongo.InitMongo(cfg.Mongo); err != nil {
			return nil, err
		}
		ylog.Debug("mongo connected ok")
	}

	if opt.LoadJwt {
		if err := yjwt.InitJwt(cfg.Jwt); err != nil {
			return nil, err
		}
	}

	if opt.LoadMq {
		if &opt.LoadQiniu != nil {
			ylog.Debug("emqx connecting...", cfg.Emq.Url, "  ;  ", cfg.Emq.Host)
			if err := ymq.InitMq(cfg.Emq); err != nil {
				return nil, err
			}
			ylog.Debug("emqx connected ok")
		}
	}

	if opt.LoadRabbitMq {
		ylog.Debug("rabbit connecting...", cfg.Emq.Url, "  ;  ", cfg.Emq.Host)
		if err := ymqa.InitRabbit(cfg.RabbitMq); err != nil {
			return nil, err
		}
		ylog.Debug("rabbitmq connected ok")
	}

	if opt.LoadES {
		ylog.Debug("elasticsearch connecting...", cfg.ES)
		if err := yes.InitES(cfg.ES); err != nil {
			return nil, err
		}
		ylog.Debug("elasticsearch connected ok")
	}

	if opt.LoadQiniu {
		yqiniu.InitQiniu(cfg.Qiniu)
	}

	return cfg, nil
}

func InitModules(loadEtcd, loadPostgres, loadRedis, loadEMQ, loadMongo bool) error {
	opt := &LoadOption{
		LoadEtcd:  loadEtcd,
		LoadPg:    loadPostgres,
		LoadRedis: loadRedis,
		LoadMongo: loadMongo,
		LoadMq:    loadEMQ,
		LoadJwt:   true,
	}

	_, err := InitModulesOfOptions(opt)
	return err
}

//初始化其它数据
func initOthers() {
	ytime.SetTimeZone()
}
