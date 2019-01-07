package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials

// Config struct for mLab database or Yelp API
type Config struct {
	Server   string
	Database string
	YelpURL  string
	YelpKey  string
}

// Read and parse the configuration file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
