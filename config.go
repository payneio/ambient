package ambient

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SmartThings SmartThingsConfig
}

type SmartThingsConfig struct {
	ClientID      string
	Secret        string
	TokenFilePath string
}

func LoadConfig(filePath string) (Config, error) {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	if _, err := toml.Decode(string(data), &conf); err != nil {
		// handle error
	}
	return conf, nil
}
