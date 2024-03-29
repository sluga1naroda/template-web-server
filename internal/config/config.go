package config

import (
	"sync"
)

var (
	config *Config
	once   sync.Once
	// Version is a parameter for providing version of current application.
	Version = ""
)

// Get is a function for getting config.
func Get() *Config {
	once.Do(func() {
		config = &Config{}
	})

	return config
}
