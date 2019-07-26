package golog

import (
	"testing"
	"time"

	"mgo/tools/golog/conf"
)

func Test_logge(t *testing.T) {
	SetLogger(ZAPLOG,
		conf.WithLogType(conf.LogJsontype),
		conf.WithProjectName("go_xxx"),
		conf.WithFilename("log.txt"),
	)

	SetLogLevel(conf.ErrorLevel)
	Debugf("this is zap")
	Debugf("this is zap")
	SetLogLevel(conf.DebugLevel)
	Debugf("this is zap")
	Debugf("this is zap")

	time.Sleep(time.Second * 5)

}
