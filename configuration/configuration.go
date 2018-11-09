package configuration

type appConfiguration struct {
	MongoDB configuration
}

type configuration interface {
	Load()
}

var conf *appConfiguration

func init() {
	conf = new(appConfiguration)
	conf.MongoDB.Load()
}
