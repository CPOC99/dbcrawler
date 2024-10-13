package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// TOMLExample represents the entire TOML structure.
// Database represents database configuration.

type Database struct {
	ExecPath      string `toml:"exec_path"`
	ExecQueryPath string `toml:"query_path"`
}

type Config struct {
	Postgres Postgres `toml:"postgres"`
}

type Postgres struct {
	Database
}

func NewConfig(filePath string) *Config {

	config := &Config{}

	// Open the TOML file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := toml.NewDecoder(file).Decode(&config); err != nil {
		return config
	}

	return config
}
