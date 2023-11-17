package main

import (
	"errors"
	"fmt"
	"go-pdirs/internal"
	"go-pdirs/pkg/resolver"
	"os"
	"strings"
)

func main() {
	environ := internal.Environ()
	files := internal.NewFiles(environ)

	contents, err := files.LoadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	projects, err := resolver.ParseProjects(contents, environ)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	names := []string{}
	for i := 1; i < len(os.Args); i++ {
		current := os.Args[i]
		switch current {
		case "--help":
			printHelp()
			return
		case "--list":
			printProjects(projects)
			return
		default:
			names = append(names, current)
		}
	}

	if len(names) == 0 {
		printHelp()
		os.Exit(1)
	}

	if err := printProjectDirs(names, projects); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func printProjects(projects resolver.Projects) {
	if len(projects) == 0 {
		fmt.Println("No projects")
		return
	}

	fmt.Fprintf(os.Stderr, "Projects:\n")
	for name, path := range projects {
		fmt.Fprintf(os.Stderr, "    %s: %s\n", name, path)
	}
}

func printProjectDirs(names []string, projects resolver.Projects) error {
	if len(names) == 0 {
		return errors.New("Missing project name")
	}

	results := []string{}
	for _, name := range names {
		if len(name) == 0 {
			continue
		}

		path, exists := projects[name]
		if !exists {
			return errors.New("Invalid project: " + name)
		}
		results = append(results, path)
	}

	if len(results) == 0 {
		return errors.New("No projects found")
	}

	for _, project := range results {
		fmt.Println(project)
	}
	return nil
}

func printHelp() {
	program := os.Args[0]

	usage := strings.Builder{}
	usage.WriteString("Usage:\n")
	usage.WriteString(fmt.Sprintf("    %s [options] <project>\n", program))
	usage.WriteString("Options:\n")
	usage.WriteString("    --help                   View help\n")
	usage.WriteString("    --list                   List all projects\n")

	fmt.Fprintf(os.Stderr, usage.String())
}
