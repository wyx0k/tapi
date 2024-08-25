package logger

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var DefaultLogger = logger{}
var _ctx context.Context

func SetupLogger(ctx context.Context) {
	_ctx = ctx
}

type logger struct {
}

func (l logger) Errorf(s string, i ...interface{}) {
	runtime.LogErrorf(_ctx, s, i...)
}

func (l logger) Warningf(s string, i ...interface{}) {
	runtime.LogWarningf(_ctx, s, i...)
}

func (l logger) Infof(s string, i ...interface{}) {
	runtime.LogInfof(_ctx, s, i...)
}

func (l logger) Debugf(s string, i ...interface{}) {
	runtime.LogDebugf(_ctx, s, i...)
}

// funcs begin-------------------------------------------
func Fatal(i ...interface{}) {
	runtime.LogFatal(_ctx, fmt.Sprint(i...))
}

func Error(i ...interface{}) {
	runtime.LogError(_ctx, fmt.Sprint(i...))
}

func Warning(i ...interface{}) {
	runtime.LogWarning(_ctx, fmt.Sprint(i...))
}

func Info(i ...interface{}) {
	runtime.LogInfo(_ctx, fmt.Sprint(i...))
}

func Debug(i ...interface{}) {
	runtime.LogDebug(_ctx, fmt.Sprint(i...))
}

func Fatalf(s string, i ...interface{}) {
	runtime.LogFatalf(_ctx, s, i...)
}

func Errorf(s string, i ...interface{}) {
	runtime.LogErrorf(_ctx, s, i...)
}

func Warningf(s string, i ...interface{}) {
	runtime.LogWarningf(_ctx, s, i...)
}

func Infof(s string, i ...interface{}) {
	runtime.LogInfof(_ctx, s, i...)
}

func Debugf(s string, i ...interface{}) {
	runtime.LogDebugf(_ctx, s, i...)
}
