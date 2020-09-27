package config

import (
	"github.com/gojetpack/goconf"
	"log"
	"time"
)

const (
	DefaultBinarizerHost       = "localhost:50053"
	DefaultImageConversionHost = "localhost:50052"
	DefaultTimeout             = time.Hour * 10
)

type config struct {
	BinarizerHost           string
	ImageConversionHost     string
	ExternalServicesTimeout time.Duration
}

var _configCache *config

func GetConfig() *config {
	if _configCache != nil {
		return _configCache
	}
	config := config{
		BinarizerHost:           DefaultBinarizerHost,
		ImageConversionHost:     DefaultImageConversionHost,
		ExternalServicesTimeout: DefaultTimeout,
	}
	args := goconf.ExtractorArgs{
		Configs: []interface{}{
			//  Config struct | env name prefix
			&config, "",
		},
	}
	err := goconf.Extract(args)
	if err != nil {
		log.Printf("error loading the configuration: %v", err)
	}
	_configCache = &config
	return _configCache
}
