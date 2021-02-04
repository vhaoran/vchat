package ylog

import (
	"fmt"
	"log"
	"sync"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/vchat/lib/yconfig"
)

var (
	constLevelSign = map[string]int{"debug": 1, "info": 2, "warn": 3, "error": 4}
	level          string
	w              *LogWorker
	mu             sync.Mutex
)

//初始化log接口
func InitLog(cfg yconfig.LogConfig) error {
	level = cfg.LogLevel
	if len(level) == 0 {
		//if not set,then level is debug
		level = "debug"
	}

	bean := &LogWorker{
		Level:      level,
		BackupPath: cfg.LogBackupPath,
		LogPath:    cfg.LogPath,
		FileName:   cfg.FileName,
		Ext:        cfg.Ext,
	}
	if len(bean.LogPath) == 0 {
		bean.LogPath = "./logs"
	}
	if len(bean.BackupPath) == 0 {
		bean.BackupPath = "./logs/backup"
	}
	if len(bean.FileName) == 0 {
		bean.BackupPath = "log"
	}
	if len(bean.Ext) == 0 {
		bean.Ext = ".log"
	}

	mu.Lock()
	defer mu.Unlock()
	w = bean
	return nil
}

// outer called
func GetLogger() *log.Logger {
	if w != nil {
		return w.GetLogger()
	}
	return nil
}

func levelN() int {
	i, ok := constLevelSign[level]
	if !ok {
		return 1
	}
	return i
}

func ok() bool {
	if w == nil {
		return false
	}
	return true
}

func DebugDump(a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 1 {
		s := spew.Sdump(a)
		w.Println("[debug]", s)
	}
}

func Debug(a ...interface{}) {
	if !ok() {
		fmt.Println(a...)
		return
	}

	if levelN() <= 1 {
		w.Println("[debug]", a)
	}
}

func DebugF(format string, a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 1 {
		w.Println("[debug]", fmt.Sprintf(format, a...))
	}
}

func Info(a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 2 {
		w.Println("[info]", a)
	}
}
func InfoDump(a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 2 {
		s := spew.Sdump(a)
		w.Println("[info]", s)
	}
}

func InfoF(format string, a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 2 {
		w.Println("[info]", fmt.Sprintf(format, a...))
	}
}

func Warn(a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 3 {
		w.Println("[warn]", a)
	}
}

func WarnF(format string, a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 3 {
		w.Println("[warn]", fmt.Sprintf(format, a...))
	}
}

func WarnDump(a interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 3 {
		s := spew.Sdump(a)
		w.Println("[warn]", s)
	}
}

func Error(a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 4 {
		w.Println("[error]", a)
	}
}

func ErrorF(format string, a ...interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 4 {
		w.Println("[error]", fmt.Sprintf(format, a...))
	}
}

func ErrorDump(a interface{}) {
	if !ok() {
		return
	}

	if levelN() <= 4 {
		s := spew.Sdump(a)
		w.Println("[error]", s)
	}
}
