package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	SocketList SocketList
}

type SocketValue struct {
	OriginUrl      string
	OriginProtocol string
	OriginOrigin   string
	ListenType     string
	ListenAddress  string
}

type SocketList []SocketValue

const filePath = "~/.config/socky-reverse/config.json"

func GetConfig() Config {
	file, err := os.Open(os.Getenv("HOME") + filePath[1:])
	if nil != err {
		log.Panic(err)
	}
	defer file.Close()

	configVar := Config{}
	err = json.NewDecoder(file).Decode(&configVar)
	if nil != err {
		log.Panic(err)
	}

	return configVar
}
