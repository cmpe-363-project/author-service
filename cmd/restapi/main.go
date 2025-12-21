package main

import (
	hardcodedrepository "author-service/internal/repository/hardcoded_adapter"
	"author-service/internal/restapi"
	"author-service/pkg/logger/slog"

	"github.com/caarlos0/env/v11"
)

type EnvVars struct {
	Host string `env:"HOST,required"`
	Port int    `env:"PORT,required"`
}

func main() {
	logger := slog.NewLogger(slog.NewLoggerArgs{
		LogFormat: "json",
	})

	envVars, err := env.ParseAs[EnvVars]()
	if err != nil {
		panic(err)
	}

	repo := hardcodedrepository.NewHardcodedRepository()

	app := &restapi.App{
		Version:    "v0.0.2.3",
		Logger:     logger,
		Repository: repo,
		Port:       envVars.Port,
		Host:       envVars.Host,
	}
	app.SetupAndRun()
}
