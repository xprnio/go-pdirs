package resolver

import (
	"errors"
	"fmt"
	"go-pdirs/pkg/utils"
	"strings"
)

type Projects map[string]string

func ParseProjects(contents string, environ utils.Environment) (Projects, error) {
	projects := make(Projects)
	for index, line := range strings.Split(contents, "\n") {
		lineIndex := index + 1
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			message := fmt.Sprintf("Invalid configuration at line %d (%s)", lineIndex, line)
			return nil, errors.New(message)
		}

		project := parts[0]
		path, err := ResolvePath(parts[1], environ)
		if err != nil {
			return nil, err
		}

		projects[project] = path
	}

	return projects, nil
}
