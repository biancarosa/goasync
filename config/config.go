package config

//Config is the struct that defines the config.toml file structure
type Config struct {
	Title     string
	Endpoints []endpoint
}

type endpoint struct {
	URL  string
	Name string
}
