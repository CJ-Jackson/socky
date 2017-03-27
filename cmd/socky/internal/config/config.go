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
	Path string
	Protocol string
	Address string
}

type SocketList []SocketValue

const filePath = "~/.config/socky/config.json"

func GetConfig() Config {
	configVar := Config{}

	file, err := os.Open(os.Getenv("HOME") + filePath[1:])
	if nil != err {
		log.Panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&configVar)
	if nil != err {
		log.Panic(err)
	}

	return configVar
}