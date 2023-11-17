package internal

import (
	"errors"
	"go-pdirs/pkg/utils"
	"os"
)

type Files struct {
	environment utils.Environment
}

func NewFiles(environment utils.Environment) *Files {
	return &Files{environment}
}

func (f *Files) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (f *Files) WriteFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

func (f *Files) LoadConfig() (string, error) {
	home, exists := f.environment.LookupEnv("HOME")
	if !exists {
		return "", errors.New("Missing environment variable: $HOME")
	}

	configPath := home + "/.pdirs"
	contents, err := f.ReadFile(configPath)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}
