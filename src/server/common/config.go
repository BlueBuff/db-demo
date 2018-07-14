package common

import "time"

type ApplicationContext struct {
	Server    Server     `yaml:"server"`
	Log       Log        `yaml:"log"`
	DBConfigs []DBConfig `yaml:"dbconfigs"`
	Redis     Redis      `yaml:"redis"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Server struct {
	Name    string `yaml:"name"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Version string `yaml:"version"`
}

type Log struct {
	Path string `yaml:"path"`
}

type DBConfig struct {
	Name   string `yaml:"name"`
	Config Config `yaml:"config"`
}

type Config struct {
	Mode            bool          `yaml:"mode"`
	Driver          string        `yaml:"driver"`
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	UserName        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	DataBaseName    string        `yaml:"databasename"`
	ConnMaxLifetime time.Duration `yaml:"lifetime"`
	MaxOpenNum      int           `yaml:"max-open-num"`
	MaxIdleNum      int           `yaml:"max-idle-num"`
}
