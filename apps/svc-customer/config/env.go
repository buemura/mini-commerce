package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	PORT                       string
	GRPC_PORT                  string
	DATABASE_URL               string
	JWT_SECRET                 string
	OTEL_EXPORTER_OTLP_ENDPOINT string
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		PORT = os.Getenv("PORT")
		GRPC_PORT = os.Getenv("GRPC_PORT")
		DATABASE_URL = os.Getenv("DATABASE_URL")
		JWT_SECRET = os.Getenv("JWT_SECRET")
		OTEL_EXPORTER_OTLP_ENDPOINT = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
		return
	}

	PORT = viper.GetString("PORT")
	GRPC_PORT = viper.GetString("GRPC_PORT")
	DATABASE_URL = viper.GetString("DATABASE_URL")
	JWT_SECRET = viper.GetString("JWT_SECRET")
	OTEL_EXPORTER_OTLP_ENDPOINT = viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT")
}
