package golog

import (
	"mgo/tools/golog/conf"
	"mgo/tools/golog/plugins/zaplog"
)

//默认
var l ILog = zaplog.New()

type backend uint8

const (
	ZAPLOG backend = iota
)

//设置
func SetLogger(b backend, opts ...conf.Option) {
	switch b {
	case ZAPLOG:
		l = zaplog.New(opts...)
	}
}

//目前只有zap生效
func SetLogLevel(level conf.Level) {
	l.SetLogLevel(level)
}

//目前只有zap生效
func Sync() error {
	return l.Sync()
}

//普通日志
//func Debug(args ...interface{}) {
//	l.Debug(args...)
//}
//func Info(args ...interface{}) {
//	l.Info(args...)
//}
//func Warn(args ...interface{}) {
//	l.Warn(args...)
//}
//func Error(args ...interface{}) {
//	l.Error(args...)
//}
//func Panic(args ...interface{}) {
//	l.Panic(args...)
//}
//func Fatal(args ...interface{}) {
//	l.Fatal(args...)
//}

//需要格式化日志dlog.error("err",err)
func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

//key value
func Debug(msg string, keysAndValues ...interface{}) {
	l.Debugw(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
	l.Infow(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	l.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	l.Errorw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	l.Panicw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	l.Fatalw(msg, keysAndValues...)
}
