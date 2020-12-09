package tools

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log zapLog
var Log *zap.Logger

// Log cutting settings
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "api.log", // 日志文件位置
		MaxSize:    10,        // 日志文件最大大小(MB)
		MaxBackups: 5,         // 保留旧文件最大数量
		MaxAge:     30,        // 保留旧文件最长天数
		Compress:   false,     // 是否压缩旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

// log encoder
func getEncoder() zapcore.Encoder {
	// Use the default JSON encoding
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// InitLogger init log
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	Log = zap.New(core, zap.AddCaller())
}
