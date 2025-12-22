package config

type AppConfigStruct struct {
	ServerHost string
	ServerPort string
}

var AppConfig AppConfigStruct

func LoadConfig() {
	AppConfig.ServerHost = "localhost"
	AppConfig.ServerPort = "8080"
}
