package configuration

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	// given
	expected := &Configuration{
		MongoDB: &MongoDBConfiguration{
			Collection: "collection",
			Database:   "database",
			Port:       "port",
			Host:       "host",
		},
	}
	os.Setenv("MONGODB_HOST", expected.MongoDB.Host)
	os.Setenv("MONGODB_PORT", expected.MongoDB.Port)
	os.Setenv("MONGODB_DATABASE", expected.MongoDB.Database)
	os.Setenv("MONGODB_COLLECTION", expected.MongoDB.Collection)
	loader := new(EnvironmentLoader)

	// when
	c, err := loader.LoadConfiguration()

	// then
	assert.Equal(t, expected.MongoDB.Collection, c.MongoDB.Collection)
	assert.Equal(t, expected.MongoDB.Database, c.MongoDB.Database)
	assert.Equal(t, expected.MongoDB.Port, c.MongoDB.Port)
	assert.Equal(t, expected.MongoDB.Host, c.MongoDB.Host)
	assert.Nil(t, err)
}
