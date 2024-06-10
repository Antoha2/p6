package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string `env:"DB_USER" env-default:"user"`
	Password string `env:"DB_PASSWORD" env-default:"password"`
	Host     string `env:"DB_HOST" env-default:"postgres"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	Dbname   string `env:"DB_DBNAME" env-default:"postgres"`
	Sslmode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func MustLoad() *Config {

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
