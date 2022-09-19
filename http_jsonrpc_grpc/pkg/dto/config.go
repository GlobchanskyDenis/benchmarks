package dto

import (
)

type StorageConfig struct {
	PathPrefix string `conf:"PathPrefix"`
}

type ServerConfig struct {
	Ip		 string `conf:"Ip"`
	Port  	 string `conf:"Port"`
	Endpoint string `conf:"Endpoint"`
}
