package ylog

import (
	"fmt"
	"github.com/vhaoran/vchat/lib/yconfig"
	"log"
	"testing"
	"time"

	"github.com/vhaoran/vchat/common/g"
	"github.com/vhaoran/vchat/common/ytime"
)

func Test_log_test(t *testing.T) {
	obj := &LogWorker{
		today:      ytime.Today(),
		BackupPath: "./log/backup",
		LogPath:    "./log",
		FileName:   "vchat",
		Ext:        ".log",
	}

	h := 100000
	bean := g.NewWaitGroupN(100)

	for i := 100; i < 100+h; i++ {
		j := i
		bean.Call(func() error {
			obj.Println(j, j)
			return nil
		})
	}
	bean.Wait()
}

func Test_moveTo_(t *testing.T) {
	src := "./log/"
	dst := "./log/backup"
	//
	err := MoveTo("a.txt", src, dst)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

}

func Test_log_single(t *testing.T) {
	obj := &LogWorker{
		today:      ytime.Today(),
		BackupPath: "./log/backup",
		LogPath:    "./log",
		FileName:   "vchat",
		Ext:        ".log",
	}

	h := 10000 * 100
	t0 := time.Now()
	defer obj.Close()
	for i := 100; i < 100+h; i++ {
		j := i
		obj.Println("hello", j)
	}
	log.Println("time:", time.Since(t0))
}

func Test_debug(t *testing.T) {

	InitLog(yconfig.LogConfig{
		LogLevel:      "debug",
		LogBackupPath: "",
		LogPath:       "",
		FileName:      "",
		Ext:           "",
	})

	Debug("abc")
	Info("info")

	level = "debug"
	InitLog(yconfig.LogConfig{
		LogLevel:      "info",
		LogBackupPath: "",
		LogPath:       "",
		FileName:      "",
		Ext:           "",
	})
	Debug("abc")
	Info("info")

}
