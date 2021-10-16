// Package zlog log包
package zlog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 的对外接口
type Logger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
}

type logger struct {
	log *zap.SugaredLogger
}

type Option interface {
	apply(*logger)
}

type optionFunc func(*logger)

func (f optionFunc) apply(l *logger) {
	f(l)
}

func NewLogger(opts ...Option) Logger {
	instance := &logger{}
	for _, opt := range opts {
		opt.apply(instance)
	}
	if instance.log == nil {
		instance.log = newLogger()
	}
	return instance
}

func newLogger() *zap.SugaredLogger {
	// 获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        // 指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 按级别显示不同颜色，不需要的话用zapcore.CapitalLevelEncoder
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder       // 显示完整路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig) // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(os.Stdout), zapcore.DebugLevel)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar() // zap.AddCaller()为显示文件名和行号，可省略, zap.AddCallerSkip是为了显示日志行号时，显示争取的位置而不是下面的问题
}

func (l *logger) Info(args ...interface{}) {
	l.log.Info(args)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.log.Infof(template, args...)
}
func (l *logger) Debug(args ...interface{}) {
	l.log.Debug(args)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.log.Debugf(template, args...)
}
func (l *logger) Warn(args ...interface{}) {
	l.log.Warn(args)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.log.Warnf(template, args...)
}
func (l *logger) Error(args ...interface{}) {
	l.log.Error(args)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.log.Errorf(template, args...)
}
