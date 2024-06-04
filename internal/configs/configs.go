package configs

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	DBConnectionString string `json:"db_connection_string"`
	WebServerPort      int    `json:"web_server_port"`
}

func New() (*Config, error) {
	var config *Config

	file, err := os.Open("../internal/configs/config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
