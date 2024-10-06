package logrus

import (
	"github.com/r1nb0/food-app/pkg/logger"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

var _ logger.Logger = (*Logger)(nil)

type Logger struct {
	logger *logrus.Logger
	config logger.Config
	mu     *sync.Mutex
}

func DefaultLogger() *Logger {
	l := Logger{
		config: logger.Config{
			Level:      logger.InfoLevel,
			TimeFormat: logger.DefaultTimeFormat,
		},
	}

	lgr := logrus.New()
	lgr.SetFormatter(&logrus.TextFormatter{
		DisableColors:   !l.config.UseColor,
		TimestampFormat: l.config.TimeFormat,
	})
	setLevel(lgr, l.config.Level)

	l.logger = lgr
	return &l
}

func New(config *logger.Config) (*Logger, error) {
	if config == nil {
		config = &logger.Config{
			Level:      logger.InfoLevel,
			TimeFormat: logger.DefaultTimeFormat,
		}
	}

	if config.TimeFormat == "" {
		config.TimeFormat = logger.DefaultTimeFormat
	}

	lgr, err := newLogger(config)
	if err != nil {
		return nil, err
	}
	l := Logger{
		logger: lgr,
		config: *config,
	}
	return &l, nil
}

func newLogger(config *logger.Config) (*logrus.Logger, error) {
	lgr := logrus.New()

	if config.LogFile != "" {
		file, err := logger.CreateLogFile(config.LogFile)
		if err != nil {
			return nil, err
		}

		writers := io.MultiWriter(os.Stderr, file)
		lgr.SetOutput(writers)
	}

	if config.UseJSON {
		lgr.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: config.TimeFormat,
		})
	} else {
		lgr.SetFormatter(&logrus.TextFormatter{
			ForceColors:     config.UseColor,
			DisableColors:   !config.UseColor,
			TimestampFormat: config.TimeFormat,
		})
	}

	lgr.SetReportCaller(config.Caller)

	setLevel(lgr, config.Level)

	return lgr, nil
}

func (l *Logger) SetConfig(config *logger.Config) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if config == nil {
		return nil
	}

	log, err := newLogger(config)
	if err != nil {
		return err
	}

	l.logger = log
	return nil
}

func setLevel(lgr *logrus.Logger, level logger.Level) {
	switch level {
	case logger.DebugLevel:
		lgr.SetLevel(logrus.DebugLevel)
	case logger.InfoLevel:
		lgr.SetLevel(logrus.InfoLevel)
	case logger.WarnLevel:
		lgr.SetLevel(logrus.WarnLevel)
	case logger.ErrorLevel:
		lgr.SetLevel(logrus.ErrorLevel)
	case logger.FatalLevel:
		lgr.SetLevel(logrus.FatalLevel)
	default:
		lgr.SetLevel(logrus.InfoLevel)
	}
}

func (l *Logger) SetLevel(level logger.Level) error {
	if level < logger.DebugLevel || level > logger.FatalLevel {
		level = logger.InfoLevel
	}

	if level != l.config.Level {
		setLevel(l.logger, level)
		l.config.Level = level
	}

	return nil
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

func (l *Logger) Debugw(message string, fields logger.KV) {
	l.logger.WithFields(logrus.Fields(fields)).Debugln(message)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *Logger) Infow(message string, fields logger.KV) {
	l.logger.WithFields(logrus.Fields(fields)).Infoln(message)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

func (l *Logger) Warnw(message string, fields logger.KV) {
	l.logger.WithFields(logrus.Fields(fields)).Warnln(message)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *Logger) Errorw(message string, fields logger.KV) {
	l.logger.WithFields(logrus.Fields(fields)).Errorln(message)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

func (l *Logger) Fatalw(message string, fields logger.KV) {
	l.logger.WithFields(logrus.Fields(fields)).Fatalln(message)
}
