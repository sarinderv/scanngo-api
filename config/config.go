package config

type Config struct {
	Database `yaml:"database"`
}

type Database struct {
	Port     string `yaml:"db_port" env:"PORT" env-default:""`
	DB       string `yaml:"db_name" env:"DATABASE" env-default:""`
	Host     string `yaml:"db_host" env:"HOST" env-default:""`
	User     string `yaml:"db_user" env:"" env-default:"root"`
	Password string `yaml:"db_password" env:"PASSWORD" env-default:""`
}