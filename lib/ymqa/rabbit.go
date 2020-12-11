package ymqa

import (
	"github.com/vhaoran/vchat/lib/yconfig"
)

var (
	X IMqA
)

func InitRabbit(cfg yconfig.RabbitConfig) error {
	obj := new(AMqWorker)
	err := obj.NewClient(cfg)
	if err != nil {
		return err
	}
	//
	X = obj
	return nil
}
