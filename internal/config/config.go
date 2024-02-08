package config

import (
	"os"
)

const (
	DEVELOPER    = "developer"
	HOMOLOGATION = "homologation"
	PRODUCTION   = "production"
)

type Config struct {
	PORT string `json:"port"`
	Mode string `json:"mode"`

	*PGSQLConfig
}

type PGSQLConfig struct {
	DB_DRIVE                  string `json:"db_drive"`
	DB_HOST                   string `json:"db_host"`
	DB_PORT                   string `json:"db_port"`
	DB_USER                   string `json:"db_user"`
	DB_PASS                   string `json:"db_pass"`
	DB_NAME                   string `json:"db_name"`
	DB_DSN                    string `json:"-"`
	DB_SET_MAX_OPEN_CONNS     int    `json:"db_set_max_open_conns"`
	DB_SET_MAX_IDLE_CONNS     int    `json:"db_set_max_idle_conns"`
	DB_SET_CONN_MAX_LIFE_TIME int    `json:"db_set_conn_max_life_time"`
	SRV_DB_SSL_MODE           bool   `json:"srv_db_ssl_mode"`
}

func NewConfig() *Config {
	conf := defaultConf()

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.PORT = SRV_PORT
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	return conf
}

func defaultConf() *Config {
	default_conf := Config{
		PORT: "8080",

		Mode: DEVELOPER,

		PGSQLConfig: &PGSQLConfig{
			DB_DRIVE: "postgres",
			DB_HOST:  "localhost",
			DB_PORT:  "5432",
			DB_USER:  "postgres",
			DB_PASS:  "supersenha",
			DB_NAME:  "fazpay_db_dev",
		},
	}

	return &default_conf
}
