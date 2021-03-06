package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/platzily/dealer/utils/constants"
	log "github.com/sirupsen/logrus"
)

type RabbitMQConfig struct {
	URL string
}

func ReadRabbitMQConfig() *RabbitMQConfig {

	urlValue := getEnvVariableAsString(constants.EnvironmentVariables.RABBIT_URL)
	return &RabbitMQConfig{
		URL: urlValue,
	}
}

func getEnvVariableAsString(name string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	envVar := os.Getenv(name)

	if len(envVar) == 0 {
		log.Fatalf("Environment variable %s is not set", name)
	}

	return envVar
}
