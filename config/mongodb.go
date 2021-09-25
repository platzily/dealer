package config

import (
	"github.com/platzily/dealer/utils/constants"
)

type MongoDBConfig struct {
	URL string
}

func ReadMongoDBConfig() *MongoDBConfig {

	urlValue := getEnvVariableAsString(constants.EnvironmentVariables.MONGO_URL)
	return &MongoDBConfig{
		URL: urlValue,
	}
}
