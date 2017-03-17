package utils

import (
	"os"
	"log"
)

func AppPort() string {
	return setIfNotEmpty("APP_PORT")
}

func InfluxAddr() string {
	return setIfNotEmpty("INFLUX_ADDR")
}

func Database() string {
	return setIfNotEmpty("DATABASE")
}

func Content() string {
	if os.Getenv("CONTENT") != "" {
		return os.Getenv("CONTENT")
	} else {
		return "public"
	}
}

func DumpAllEnvVars() {
	vars := map[string]string{
		"APP_PORT": setIfNotEmpty("APP_PORT"),
		"INFLUX_ADDR": setIfNotEmpty("INFLUX_ADDR"),
		"DATABASE": setIfNotEmpty("DATABASE"),
		"CONTENT": setIfNotEmpty("CONTENT"),
	}
	for k, v := range vars {
		log.Printf("%s=%s", k, v)
	}
}

func setIfNotEmpty(env string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	} else {
		panic("You have to set up: " + env)
	}
}