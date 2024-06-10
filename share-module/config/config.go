package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host       string `mapstructure:"HOST"`
	DBPort     string `mapstructure:"PORT"`
	DBUserName string `mapstructure:"USERNAME"`
	DBPassword string `mapstructure:"PASSWORD"`
	DBName     string `mapstructure:"NAME"`
}

func (dbc DBConfig) Uri() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbc.Host, dbc.DBUserName, dbc.DBPassword, dbc.DBName, dbc.DBPort)
}

type WebConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

func (wc WebConfig) Address() string {
	return fmt.Sprintf(":%s", wc.Port)
}

type AppConfig struct {
	Env             string        `mapstructure:"ENV"`
	ShutdownTimeout time.Duration `mapstructure:"SHUTDOWN_TIMEOUT"`
	Web             WebConfig     `mapstructure:"WEB"`
	DB              DBConfig      `mapstructure:"DB"`
}

func LoadAppConfig(path string) (AppConfig, error) {

	var cfg AppConfig
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.EnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
