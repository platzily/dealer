package constants

type environment struct {
	RABBIT_URL string
}

var EnvironmentVariables environment = environment{
	RABBIT_URL: "RABBIT_URL",
}
