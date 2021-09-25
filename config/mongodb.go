package config

import (
	"github.com/platzily/dealer/utils/constants"
)

type MongoDBConfig struct {
	URL string
}

func ReadMongoDBConfig() *MongoDBConfig {

	urlValue := getEnvVariableAsString(constants.MongoURL)
	return &MongoDBConfig{
		URL: urlValue,
	}
}
