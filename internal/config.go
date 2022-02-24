package internal

import "github.com/joeshaw/envdecode"

var c Config

type Config struct {
	Tipo     string `env:"TIPO"`
	Suc      string `env:"SUC"`
	Numero   string `env:"NUMERO"`
	Interval int    `env:"INTERVAL,default=3600"` // this is in seconds
	Token    string `env:"TG_TOKEN"`
}

func GetConfig() (*Config, error) {
	err := envdecode.Decode(&c)
	return &c, err
}
