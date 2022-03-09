package main

import (
	"context"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
)

type Config struct {
	Port     string `config:"port"`
	ApiKey   string `config:"ApiKey"`
	City     string `config:"City"`
	Unit     string `config:"Unit"`
	Language string `config:"Language"`
}

func Load() *Config {
	loaders := []backend.Backend{
		file.NewOptionalBackend("/etc/openweather_exporter.yaml"),
		flags.NewBackend(),
	}

	loader := confita.NewLoader(loaders...)

	cfg := Config{
		Port:     ":9091",
		ApiKey:   "",
		City:     "Europe/Athens",
		Unit:     "C",
		Language: "EN",
	}

	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
