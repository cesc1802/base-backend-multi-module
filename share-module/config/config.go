package config

import "time"

type DBConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	Name     string
}

func (dbc DBConfig) Uri() string {
	return ""
}

type WebConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:":8080"`
}

func (wc WebConfig) Address() string {
	return ""
}

type AppConfig struct {
	Env             string
	ShutdownTimeout time.Duration
	Web             WebConfig
	DB              DBConfig
}
