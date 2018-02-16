package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

//ParseTOML parses a config.toml file
func ParseTOML() Config {
	config := Config{}
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		print(err)
	}
	fmt.Printf("%#v", config)
	return config
}
