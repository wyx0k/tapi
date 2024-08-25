package logger

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var DefaultLogger = logger{}

type logger struct {
}

func (l logger) Errorf(s string, i ...interface{}) {
	runtime.LogErrorf(context.Background(), s, i...)
}

func (l logger) Warningf(s string, i ...interface{}) {
	runtime.LogWarningf(context.Background(), s, i...)
}

func (l logger) Infof(s string, i ...interface{}) {
	runtime.LogInfof(context.Background(), s, i...)
}

func (l logger) Debugf(s string, i ...interface{}) {
	runtime.LogDebugf(context.Background(), s, i...)
}

// funcs begin-------------------------------------------
func Fatal(i ...interface{}) {
	runtime.LogFatal(context.Background(), fmt.Sprint(i...))
}

func Error(i ...interface{}) {
	runtime.LogError(context.Background(), fmt.Sprint(i...))
}

func Warning(i ...interface{}) {
	runtime.LogWarning(context.Background(), fmt.Sprint(i...))
}

func Info(i ...interface{}) {
	runtime.LogInfo(context.Background(), fmt.Sprint(i...))
}

func Debug(i ...interface{}) {
	runtime.LogDebug(context.Background(), fmt.Sprint(i...))
}

func Fatalf(s string, i ...interface{}) {
	runtime.LogFatalf(context.Background(), s, i...)
}

func Errorf(s string, i ...interface{}) {
	runtime.LogErrorf(context.Background(), s, i...)
}

func Warningf(s string, i ...interface{}) {
	runtime.LogWarningf(context.Background(), s, i...)
}

func Infof(s string, i ...interface{}) {
	runtime.LogInfof(context.Background(), s, i...)
}

func Debugf(s string, i ...interface{}) {
	runtime.LogDebugf(context.Background(), s, i...)
}
