package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
)

// TOMLExample represents the entire TOML structure.
// Database represents database configuration.

type Database struct {
	ExecPath      string `toml:"exec_path"`
	ExecQueryPath string `toml:"query_path"`
	ExecMccsPath  string `toml:"exec_mccs_path"`
}

type Config struct {
	Postgres Postgres `toml:"postgres"`
}

type Postgres struct {
	Database `mapstructure:",squash"`
}

func NewConfig(filePath string) *Config {

	config := &Config{}

	// Open the TOML file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	toml.NewDecoder(file).Decode(&config)
	println(config.Postgres.ExecPath)
	return config
}

func (c *Config) StructToMap() map[string]interface{} {
	result := make(map[string]interface{})

	err := mapstructure.Decode(c, &result)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return result
}
