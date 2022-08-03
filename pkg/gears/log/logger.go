package log

import (
	"github.com/golang/glog"
)

func Debugf(format string, args ...any) {
	glog.Infof(format, args...)
}

func Infof(format string, args ...any) {
	glog.Infof(format, args...)
}

func Warningf(format string, args ...any) {
	glog.Warningf(format, args...)
}

func Errorf(format string, args ...any) {
	glog.Errorf(format, args...)
}

func Fatalf(format string, args ...any) {
	glog.Fatalf(format, args...)
}
