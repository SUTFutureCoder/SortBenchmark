/*
   该包实现基础的分级日志打印
   但其实很基础
   很多东西不支持, 如日志滚动等
   之后可以在接口不变基础上, 使用其他的包, 比如beego的日志包
*/
package logs

import (
	"log"
)

var debug bool = false

func Debug(v ...interface{}) {
	if !debug {
		return
	}
	log.SetPrefix("[Debug] ")
	log.Println(v...)
}

func Debugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	log.SetPrefix("[Debug] ")
	log.Printf(format+"\n", v...)
}

func Trace(v ...interface{}) {
	log.SetPrefix("[Trace] ")
	log.Println(v...)
}

func Tracef(format string, v ...interface{}) {
	log.SetPrefix("[Trace] ")
	log.Printf(format+"\n", v...)
}

func Warning(v ...interface{}) {
	log.SetPrefix("[Warning] ")
	log.Println(v...)
}

func Fatal(v ...interface{}) {
	log.SetPrefix("[Fatal] ")
	log.Println(v...)
}

func Fatalf(format string, v ...interface{}) {
	log.SetPrefix("[Fatal] ")
	log.Printf(format+"\n", v...)
}

func Warningf(format string, v ...interface{}) {
	log.SetPrefix("[Warning] ")
	log.Printf(format+"\n", v...)
}
