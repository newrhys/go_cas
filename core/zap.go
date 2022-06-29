package core

import (
	"fmt"
	"wave-admin/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)
var level zapcore.Level

func InitLogger() (logger *zap.Logger) {
	switch global.GnConfig.Zap.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	writeSyncer := getLogWriter()
	encoder := getEncoder()

	core := zapcore.NewCore(
		encoder,		// 编码器配置
		writeSyncer, 	// 打印到控制台和文件
		atomicLevel,	// 日志级别
	)

	development := zap.Development()
	caller := zap.AddCaller()
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(core, zap.AddStacktrace(level), development)
	} else {
		logger = zap.New(core, caller, development)
	}

	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	switch {
	case global.GnConfig.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.GnConfig.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.GnConfig.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.GnConfig.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	if global.GnConfig.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	now := time.Now()
	fileName := fmt.Sprintf("%d-%d-%d.log",now.Year(),now.Month(),now.Day())
	hook := lumberjack.Logger{
		Filename:   global.GnConfig.Zap.Director + fileName, // 日志文件路径
		MaxSize:    global.GnConfig.Zap.MaxSize,             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: global.GnConfig.Zap.MaxBackups,          // 日志文件最多保存多少个备份
		MaxAge:     global.GnConfig.Zap.MaxAge,              // 文件最多保存多少天
		Compress:   global.GnConfig.Zap.Compress,            // 是否压缩
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GnConfig.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}