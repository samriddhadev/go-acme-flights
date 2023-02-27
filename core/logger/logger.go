package logger

import (
	"fmt"
	"log"

	"github.com/samriddhadev/go-acme-flights/core/config"
	"github.com/samriddhadev/go-acme-flights/core/environment"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const ENVIRONMENT string = "ENVIRONMENT"
const LOGGER_FILE_NAME string = "app-logger"

type LogEncoderConfig struct {
	TimeKey string `yaml:"timeKey"`
	LevelKey string `yaml:"levelKey"`
	MessageKey string `yaml:"messageKey"`
	NameKey string `yaml:"nameKey"`
	StacktraceKey string `yaml:"stacktraceKey"`
	CallerKey string `yaml:"callerKey"`
	LineEnding string `yaml:"lineEnding"`
	LevelEncoder string `yaml:"levelEncoder"`
}

type LogStructure struct {
	Level string `yaml:"level"`
	Development bool `yaml:"development"`
	DisableCaller bool `yaml:"disableCaller"`
	DisableStacktrace bool `yaml:"disableStacktrace"`
	Encoding string `yaml:"encoding"`
	OutputPaths []string `yaml:"outputPaths"`
	ErrorOutputPaths []string `yaml:"errorOutputPaths"`
	EncoderConfig LogEncoderConfig `yaml:"encoderConfig"`
}

func (ls *LogStructure) Build() *zap.Logger {
	lvl, err := zap.ParseAtomicLevel(ls.Level); if err != nil {
		log.Panicf("can't iitialize zap logger - invalid log level - : %v", err)
	}
	encoderCfg := &zapcore.EncoderConfig{
		MessageKey: ls.EncoderConfig.MessageKey,
		LevelKey: ls.EncoderConfig.LevelKey,
		TimeKey: ls.EncoderConfig.TimeKey,
		NameKey: ls.EncoderConfig.NameKey,
		CallerKey: ls.EncoderConfig.CallerKey,
		StacktraceKey: ls.EncoderConfig.StacktraceKey,
		LineEnding: zapcore.DefaultLineEnding,
		EncodeTime: zapcore.RFC3339TimeEncoder,
	}

	cfg := &zap.Config{
		Level: lvl,
		Development: ls.Development,
		DisableCaller: ls.DisableCaller,
		DisableStacktrace: ls.DisableStacktrace,
		Encoding: ls.Encoding,
		EncoderConfig: *encoderCfg,
		OutputPaths: ls.OutputPaths,
		ErrorOutputPaths: ls.ErrorOutputPaths,
	}
	logger, err := cfg.Build(); if err != nil {
		log.Panicf("can't iitialize zap logger: %v", err)
	}
	return logger
}

func NewAcmeLogger(cfg *config.Config) *AcmeLogger {
	config, err := load(); if err != nil {
		log.Panicf("can't load zap logger config: %v", err)
	}
	logger := config.Build()
	acmeLogger := &AcmeLogger{
		cfg: cfg,
	}
	acmeLogger.log = logger
	return acmeLogger
} 

func load() (*LogStructure, error) {
	viper.SetConfigName(fmt.Sprintf("%s.%s", LOGGER_FILE_NAME, environment.GetEnv(ENVIRONMENT, "local")))
	viper.AddConfigPath("./resources")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("error - reading config file - %s", err)
		return &LogStructure{}, err
	}
	config := &LogStructure{}
	err = viper.Unmarshal(&config)
	return config, err
}

type AcmeLogger struct {
	cfg *config.Config
	log *zap.Logger
}


func (lg *AcmeLogger) Info(msg string, fields ...zap.Field) {
	lg.log.Info(msg, fields...)
}

func (lg *AcmeLogger) Infof(msg string, fields ...interface{}) {
	lg.log.Sugar().Infof(msg, fields...)
}

func (lg *AcmeLogger) Debug(msg string, fields ...zap.Field) {
	lg.log.Debug(msg, fields...)
}

func (lg *AcmeLogger) Debugf(msg string, fields ...interface{}) {
	lg.log.Sugar().Debugf(msg, fields...)
}

func (lg *AcmeLogger) Warn(msg string, fields ...zap.Field) {
	lg.log.Warn(msg, fields...)
}

func (lg *AcmeLogger) Warnf(msg string, fields ...interface{}) {
	lg.log.Sugar().Warnf(msg, fields...)
}

func (lg *AcmeLogger) Error(msg string, fields ...zap.Field) {
	lg.log.Error(msg, fields...)
}

func (lg *AcmeLogger) Errorf(msg string, fields ...interface{}) {
	lg.log.Sugar().Errorf(msg, fields)
}