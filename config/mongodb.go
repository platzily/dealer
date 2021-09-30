package config

import (
	"github.com/platzily/dealer/utils/constants"
)

type MongoDBConfig struct {
	URL      string
	Database string
}

func ReadMongoDBConfig() *MongoDBConfig {

	urlValue := getEnvVariableAsString(constants.EnvironmentVariables.MONGO_URL)
	databaseName := getEnvVariableAsString(constants.EnvironmentVariables.MONGO_DATABASE)
	return &MongoDBConfig{
		URL:      urlValue,
		Database: databaseName,
	}
}
