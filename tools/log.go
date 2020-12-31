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
		Filename:   "api.log", // Log file location
		MaxSize:    10,        // Maximum log file size(MB)
		MaxBackups: 5,         // Keep the maximum number of old files
		MaxAge:     30,        // Maximum number of days to keep old files
		Compress:   false,     // Whether to compress old files
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
