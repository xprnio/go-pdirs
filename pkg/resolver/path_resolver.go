package resolver

import (
	"errors"
	"go-pdirs/pkg/utils"
	"regexp"
	"strings"
)

const EnvPattern = "\\$[A-Za-z_][A-Za-z0-9_]*"

func ResolvePath(path string, env utils.Environment) (string, error) {
	re := regexp.MustCompile(EnvPattern)
	envs := re.FindAllString(path, -1)

	for _, rawEnv := range envs {
		e := rawEnv[1:]
		value, exists := env.LookupEnv(e)
		if !exists {
			return "", errors.New("Missing environment variable: " + e)
		}
		path = strings.ReplaceAll(path, rawEnv, value)
	}

	return path, nil
}
