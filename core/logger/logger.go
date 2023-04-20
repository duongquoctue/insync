package logger

import (
	"go.uber.org/zap"
)

var logger *Logger

type Logger struct {
	sugaredLogger *zap.SugaredLogger
}

func InitLogger() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.OutputPaths = []string{"stdout"}

	l, _ := cfg.Build()
	l = l.WithOptions(
		zap.AddCallerSkip(1),
	)
	defer l.Sync()

	logger = &Logger{
		sugaredLogger: l.Sugar(),
	}
}

// Debug
func (l *Logger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Debugw(msg, keysAndValues...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugaredLogger.Debugf(template, args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.sugaredLogger.Debugln(args...)
}

// Info
func (l *Logger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Infow(msg, keysAndValues...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugaredLogger.Infof(template, args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.sugaredLogger.Infoln(args...)
}

// Warn
func (l *Logger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Warnw(msg, keysAndValues...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.sugaredLogger.Warnf(template, args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.sugaredLogger.Warnln(args...)
}

// Error
func (l *Logger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Errorw(msg, keysAndValues...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugaredLogger.Errorf(template, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.sugaredLogger.Errorln(args...)
}

// Fatal
func (l *Logger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Fatalw(msg, keysAndValues...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.sugaredLogger.Fatalf(template, args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.sugaredLogger.Fatalln(args...)
}

// Panic
func (l *Logger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.sugaredLogger.Panicw(msg, keysAndValues...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.sugaredLogger.Panicf(template, args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.sugaredLogger.Panicln(args...)
}
