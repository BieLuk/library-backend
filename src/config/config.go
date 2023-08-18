package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DatabaseEngine string

var (
	DatabaseEnginePostgres DatabaseEngine = "POSTGRES"
	DatabaseEngineMongo    DatabaseEngine = "MONGO"
)

type Config struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBUserName     string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	MongoDBName    string `mapstructure:"MONGODB_NAME"`
	MongoDBURI     string `mapstructure:"MONGODB_URI"`
	DatabaseEngine string `mapstructure:"DATABASE_ENGINE"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &config, nil
}
