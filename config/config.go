package config

import (
	"os"
	"strconv"
)

type Config struct {
	Uri          string
	ConnTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

func GetConfigFromEnvVariable() *Config {
		uri := os.Getenv("URI")

		if uri == "" {
			panic("Cannot get URI")
		}

		connTimeout, err := strconv.Atoi(os.Getenv("ConnTimeout"))

		if err != nil {
			panic("Cannot parse ConnTimeout")
		}

		readTimeout, err := strconv.Atoi(os.Getenv("ReadTimeout"))

		if err != nil {
			panic("Cannot parse ReadTimeout")
		}

		writeTimeout, err := strconv.Atoi(os.Getenv("WriteTimeout"))

		if err != nil {
			panic("Cannot parse ReadTimeout")
		}

	return &Config{
		Uri:          uri,
		ConnTimeout:  connTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}
