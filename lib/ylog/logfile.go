package ylog

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func MkdirValidate(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return err
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			return errors.New(fmt.Sprint("创建目录失败", err))
		}
	}
	return nil
}

// GetCurFilename
// Get current file name, without Suffix
func FileNameValidate(path, fileName, fileSuffix, ext string) (string, error) {
	full := ""
	if err := MkdirValidate(path); err != nil {
		return "", err
	}

	name := fileName + fileSuffix + ext
	full = filepath.Join(path, name)
	return full, nil
}

//自动绑定文件和log
func NewBindLogger(path, fileName, fileSuffix, ext string) (*log.Logger, *os.File, error) {
	full, err := FileNameValidate(path, fileName, fileSuffix, ext)
	logFile, err := os.OpenFile(full, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("open file error=%s\r\n", err.Error())
		return nil, nil, err
	}
	//defer logFile.Close()
	writers := []io.Writer{
		logFile, os.Stdout}

	//if levelN() <= 1 {
	//	writers = append(writers, os.Stdout)
	//}

	multiWriter := io.MultiWriter(writers...)

	//logger := log.New(logFile, "\r\n", log.Ldate|log.Ltime)
	logger := log.New(multiWriter, "", log.Ldate|log.Ltime)
	logger.SetFlags(log.LstdFlags)

	return logger, logFile, nil
}

func MoveTo(fileName, srcPath, dstPath string) error {
	log.Println("enter MoveTo")
	log.Println(fileName, ",", srcPath, ",", dstPath)

	//
	if err := MkdirValidate(dstPath); err != nil {
		return err
	}
	//
	src := filepath.Join(srcPath, fileName)
	dst := GetUniqueName(fileName, dstPath)

	if err := os.Rename(src, dst); err != nil {
		return err
	}

	return nil
}

func GetUniqueName(fileName, path string) string {
	ext := filepath.Ext(fileName)
	prefix := fileName[:len(fileName)-len(ext)]
	log.Println("prefix:", prefix)

	full := filepath.Join(path, fileName)
	if !FileExists(full) {
		return full
	}

	//
	for i := 1; i < 1000; i++ {
		full = filepath.Join(path, fmt.Sprint(prefix, "_", i, ext))
		if !FileExists(full) {
			return full
		}
	}

	full = filepath.Join(path, fileName)
	return full
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)

	}

	return true
}
