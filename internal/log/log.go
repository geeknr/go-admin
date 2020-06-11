package log

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go-admin-svr/internal/util/pathutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"path/filepath"
	"time"
)

var (
	Zap *zap.SugaredLogger
)

type Options struct {
	TimeFormat string
	LogName    string
	MaxSizeMB  int
	MaxBackups int
	MaxAgeDay  int
	Compress   bool // 压缩日志.(分割时)
}

func Init(customLog string) error {
	var err error
	if customLog == "" {
		customLog = filepath.Join(pathutil.WorkDir(), "logs", "app.log")
	} else {
		customLog = filepath.Join(pathutil.WorkDir(), "logs", customLog)
	}

	options := &Options{
		TimeFormat: "2006-01-02 15:04:05",
		LogName:    customLog,
		MaxAgeDay:  7,
		MaxSizeMB:  50, //50M
		MaxBackups: 3,
		Compress:   true,
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.TimeKey = "time"
	encoder.CallerKey = "file"
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { //时间格式编码器
		enc.AppendString(t.Format(options.TimeFormat))
	}

	level := zapcore.FatalLevel
	if viper.GetString("mode") == "debug" {
		level = zapcore.DebugLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(writer(options)), level),
	)

	logger := zap.New(core, zap.AddCaller())
	if logger == nil {
		return errors.Wrap(err, "init logger")
	}
	Zap = logger.Sugar()
	return nil
}

func writer(option *Options) io.Writer {
	return &lumberjack.Logger{
		Filename:   option.LogName,
		MaxSize:    option.MaxSizeMB,
		MaxBackups: option.MaxBackups,
		MaxAge:     option.MaxAgeDay,
		Compress:   option.Compress,
	}
}
