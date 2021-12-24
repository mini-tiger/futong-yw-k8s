package cfg

import (
	"os"
	"time"

	"ftk8s/base/enum"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Zlog *zap.Logger
	Mlog *zap.SugaredLogger
)

// Log config info
type LogConf struct {
	LogLevel string
	LogFile  string
}

func InitZapLogger(appRunMode string, logLevelString string, logPath string) {
	Zlog, Mlog = NewZapLogger(appRunMode, logLevelString, logPath)
}

// Create Zap logger object
func NewZapLogger(appRunMode string, logLevelString string,
	logPath string) (zlog *zap.Logger, mlog *zap.SugaredLogger) {
	// Configure log level
	logLevel := convertZapLogLevel(logLevelString)

	// Configure log files
	syncFile := zapcore.AddSync(&lumberjack.Logger{
		Filename: logPath,
		MaxSize:  100,
		// MaxAge:     365,
		// MaxBackups: 100,
		LocalTime: true,
		Compress:  true,
	})

	var coreObj zapcore.Core
	switch appRunMode {
	case enum.AppRunModeRelease:
		coreObj = zapcore.NewCore(
			zapcore.NewJSONEncoder(newEncoderConfigFile()), // Log serialization
			zapcore.NewMultiWriteSyncer(syncFile),          // Log output location
			logLevel,                                       // Log level
		)
	default:
		coreObj = zapcore.NewCore(
			zapcore.NewConsoleEncoder(newEncoderConfigConsole()),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			logLevel,
		)
	}

	zlog = zap.New(coreObj, zap.AddCaller())
	mlog = zlog.Sugar()

	zlog.Info("successfully to create Zap logger object")

	return zlog, mlog
}

// convertZapLogLevel convert level of string to zapcore.Level type.
func convertZapLogLevel(levelString string) zapcore.Level {
	switch levelString {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

// Log encoding configuration of file
func newEncoderConfigFile() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "zaplogger",
		CallerKey:      "func",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

// Log encoding configuration of console
func newEncoderConfigConsole() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "zaplogger",
		CallerKey:      "func",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

// Set time format
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
