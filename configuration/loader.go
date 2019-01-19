package configuration

import "os"

//Configuration defines all the configuration this app needs to run
type Configuration struct {
	MongoDB *MongoDBConfiguration
}

//Loader defines how a configuration loader should behave
type Loader interface {
	LoadConfiguration() (*Configuration, error)
}

//EnvironmentLoader is the environment variable loader
type EnvironmentLoader struct{}

//LoadConfiguration loads configuration from environment variables
func (el *EnvironmentLoader) LoadConfiguration() (*Configuration, error) {

	c := new(Configuration)

	c.MongoDB = new(MongoDBConfiguration)
	c.MongoDB.Host = os.Getenv("MONGODB_HOST")
	c.MongoDB.Port = os.Getenv("MONGODB_PORT")
	c.MongoDB.Database = os.Getenv("MONGODB_DATABASE")
	c.MongoDB.Collection = os.Getenv("MONGODB_COLLECTION")

	return c, nil
}
