package log

import "go.uber.org/zap"

// IZap represents an interface of zap.Logger.
type IZap interface {
	DPanic(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Sync() error
}
