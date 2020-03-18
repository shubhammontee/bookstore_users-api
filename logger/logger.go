package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	//Log will be available throughout the application
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},                  //log to standard output
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), //level is set to info
		Encoding:    "json",                              //all logs generated will be in json
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level", //so the log generated will be like Info:-->msg
			TimeKey:      "time",  //log will show time like time:-->the time
			MessageKey:   "msg",   //message with key as msg:-->the msg we give
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error

	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// func GetLogger() log{
// 	return log
// }

//Info ...
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

//Error ...
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	log.Sync()
}
