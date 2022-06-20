package log

import (
	"github.com/golang/glog"
)

func Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func Warningf(format string, args ...interface{}) {
	glog.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
}
