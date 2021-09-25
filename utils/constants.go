package constants

type environment struct {
	RABBIT_URL string
	MONGO_URL  string
}

var EnvironmentVariables environment = environment{
	RABBIT_URL: "RABBIT_URL",
	MONGO_URL:  "MONGO_URL",
}
