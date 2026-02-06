package logger

import (
	"fmt"
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getLogWriter(filepath string) zapcore.WriteSyncer {
	file, _ := os.Create(filepath)
	return zapcore.AddSync(file)
}

func GetLogger(encoding string) (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()

	cfg.Encoding = encoding

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("Error building logger: %v", err)
	}

	return logger.Sugar(), nil
}

func GetConsoleLogger() *zap.SugaredLogger {
	encoder := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoder, os.Stdout, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func GetFileLogger(filepath string) *zap.SugaredLogger {
	writerSyncer := getLogWriter(filepath)
	encoder := ecszap.NewDefaultEncoderConfig()

	core := ecszap.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

/*
ECS is a standard log format used by:
	ElasticSearch
	Kibana
	Filebeat
	Logstash

This makes logs:
	Queryable
	Structured
	Consistent across services

Use-case
	Docker / Kubernetes
	Logs scraped from stdout
	Elastic Stack pipelines
*/
