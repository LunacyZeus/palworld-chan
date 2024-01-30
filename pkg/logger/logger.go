package logger

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevel      zapcore.Level
	defaultLogger *zap.Logger
)

func InitLogger() {
	logLevel = zapcore.DebugLevel
	//if config.C.Logger.LogLevel == "Info" {
	//logLevel = zapcore.InfoLevel
	//}
	setLogConfig()
}

func setLogConfig() {
	// 设置日志输出格式
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	})

	// 添加日志切割归档功能

	hook := lumberjack.Logger{
		Filename:   "log.log", // 日志文件路径
		MaxSize:    512,       // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 1,         // 日志文件最多保存多少个备份
		MaxAge:     1,         // 文件最多保存多少天
		Compress:   true,      // 是否压缩
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)),
		zap.NewAtomicLevelAt(logLevel),
	)

	caller := zap.AddCaller() // 开启文件及行号
	skip := zap.AddCallerSkip(1)
	development := zap.Development()                   // 开启开发模式，堆栈跟踪
	logger := zap.New(core, caller, skip, development) // 构造日志
	zap.ReplaceGlobals(logger)                         // 将自定义的logger替换为全局的logger
	defaultLogger = zap.L()
}

func GetDefaultLogger() *zap.Logger {
	return defaultLogger
}

func ZapInfo(msg string, fields ...zapcore.Field) {
	defaultLogger.Info(msg, fields...)
}

func ZapWarn(msg string, fields ...zapcore.Field) {
	defaultLogger.Warn(msg, fields...)
}

func ZapDebug(msg string, fields ...zapcore.Field) {
	defaultLogger.Debug(msg, fields...)
}

func ZapError(msg string, fields ...zapcore.Field) {
	defaultLogger.Error(msg, fields...)
}

func ZapDPanic(msg string, fields ...zapcore.Field) {
	defaultLogger.DPanic(msg, fields...)
}

func ZapPanic(msg string, fields ...zapcore.Field) {
	defaultLogger.Panic(msg, fields...)
}

func ZapFatal(msg string, fields ...zapcore.Field) {
	defaultLogger.Fatal(msg, fields...)
}

func Info(format string, v ...interface{}) {
	if zapcore.InfoLevel < logLevel {
		return
	}
	if len(v) == 0 {
		defaultLogger.Info(format)
		return
	}
	defaultLogger.Info(fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	if zapcore.WarnLevel < logLevel {
		return
	}
	if len(v) == 0 {
		defaultLogger.Warn(format)
		return
	}
	defaultLogger.Warn(fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	if zapcore.DebugLevel < logLevel {
		return
	}
	if len(v) == 0 {
		defaultLogger.Debug(format)
		return
	}
	defaultLogger.Debug(fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	if zapcore.ErrorLevel < logLevel {
		return
	}
	if len(v) == 0 {
		defaultLogger.Error(format)
		return
	}
	defaultLogger.Error(fmt.Sprintf(format, v...))
}

func Fatal(format string, v ...interface{}) {
	if zapcore.FatalLevel < logLevel {
		return
	}
	if len(v) == 0 {
		defaultLogger.Fatal(format)
		return
	}
	defaultLogger.Fatal(fmt.Sprintf(format, v...))
}

func init() {
	InitLogger()
}
