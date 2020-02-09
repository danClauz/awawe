package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func loadJSONEnvPathOrPanic(envPath string, config interface{}) {
	if err := loadJSONEnvPath(envPath, config); err != nil {
		panic(fmt.Errorf("failed to load config file. Error: %s", err.Error()))
	}
}

func loadJSONEnvPath(envPath string, config interface{}) error {
	if config == nil {
		return errors.New("object config empty")
	}

	filename := os.Getenv(envPath)
	if filename == "" {
		return fmt.Errorf("env var is empty: %s", envPath)
	}

	envByte, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(envByte, config)
}
