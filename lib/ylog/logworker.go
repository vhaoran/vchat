package ylog

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/vchat/common/ytime"
)

type (
	LogWorker struct {
		Level string
		today time.Time
		sync.Mutex

		BackupPath string

		LogPath  string
		FileName string
		Ext      string

		log     *log.Logger
		logFile *os.File
	}
)

func (r *LogWorker) GetLogger() *log.Logger {
	return r.log
}

func (r *LogWorker) Close() error {
	r.Lock()
	defer r.Unlock()
	return r.logFile.Close()
}

func (r *LogWorker) backup() error {
	//r.Lock()
	//defer r.Unlock()

	//backup
	today := ytime.Today()
	//
	if r.today != today && r.log != nil {
		//backup old
		todayStr := ytime.OfToday().ToStrDate()
		fileName := r.FileName + todayStr + r.Ext
		r.log = nil
		r.logFile.Close()

		if err := MoveTo(fileName, r.LogPath, r.BackupPath); err != nil {
			return err
		}
	}
	return nil
}

func (r *LogWorker) validateCreate() error {
	r.Lock()
	defer r.Unlock()

	r.backup()

	if r.log == nil {
		suffix := ytime.OfToday().ToStrDate()
		r.today = ytime.Today()

		l, hFile, err := NewBindLogger(
			r.LogPath,
			r.FileName,
			suffix,
			r.Ext)
		if err != nil {
			return err
		}
		r.log = l
		r.logFile = hFile
	}

	return nil
}

func (r *LogWorker) Println(a ...interface{}) {
	if r.validateCreate() == nil {
		r.log.Println(a...)
	}
}

func (r *LogWorker) PrintF(format string, v ...interface{}) {
	if r.validateCreate() == nil {
		r.log.Printf(format, v...)
	}
}

func (r *LogWorker) Dump(v interface{}) {
	if r.validateCreate() == nil {
		s := spew.Sdump(v)
		r.log.Println(s)
	}
}
