package utils

import (
	"log"
	"os"
	"syscall"
)

var (
	ERR *log.Logger
	Info *log.Logger
	System *log.Logger
)



func makeBasePath(path string) bool {

	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	err = os.Mkdir(path, os.ModePerm - 022)
	if err == nil {
		return true
	}

	return false
}
func init()  {

	basepath := "log/"

	stat := makeBasePath(basepath)

	if !stat {
		log.Println("创建日志目录失败")
	}

	errfile,err := os.OpenFile(basepath + "errors.log", syscall.O_CREAT| syscall.O_RDWR|syscall.O_APPEND , 0755)
	if err != nil {
		log.Fatal(err)
	}
	ERR = log.New(errfile,"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infofile,err := os.OpenFile(basepath + "info.log", syscall.O_CREAT| syscall.O_RDWR|syscall.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	Info = log.New(infofile,"INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	systemlog,err := os.OpenFile(basepath + "system.log", syscall.O_CREAT| syscall.O_RDWR|syscall.O_APPEND , 0755)

	if err != nil {
		log.Fatal(err)
	}
	System =  log.New(systemlog,"SYSTEM: ", log.Ldate|log.Ltime|log.Lshortfile)
}
