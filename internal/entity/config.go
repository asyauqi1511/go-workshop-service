package entity

type AppConfig struct {
	DB    DBConfig    `yaml:"db"`
	Redis RedisConfig `yaml:"redis"`
}

type DBConfig struct {
	Hostname     string `yaml:"hostname"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"database_name"`
}

type RedisConfig struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}
