package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	module      string
	name        string
	owner       string
	repo        string
	description string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "init template: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	config := parseFlags()
	if config.needsPrompt() {
		if err := config.promptMissing(os.Stdin, os.Stdout); err != nil {
			return err
		}
	}
	if err := config.validate(); err != nil {
		return err
	}

	replacements := []replacement{
		{old: "github.com/example/your-cli", new: config.module},
		{old: "your-cli", new: config.name},
		{old: "example", new: config.owner},
		{old: "your-cli", new: config.repo},
		{old: "Your CLI description", new: config.description},
	}

	if err := filepath.WalkDir(".", func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if shouldSkipDir(path) {
				return filepath.SkipDir
			}
			return nil
		}
		if shouldSkipFile(path) {
			return nil
		}
		return replaceInFile(path, replacements)
	}); err != nil {
		return err
	}

	if _, err := os.Stat("README.template.md"); err == nil {
		if err := os.Remove("README.md"); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
		return os.Rename("README.template.md", "README.md")
	} else if errors.Is(err, os.ErrNotExist) {
		return nil
	} else {
		return err
	}
}

func parseFlags() config {
	config := config{}
	flag.StringVar(&config.module, "module", "", "Go module path, for example github.com/you/your-cli")
	flag.StringVar(&config.name, "name", "", "CLI binary name")
	flag.StringVar(&config.owner, "owner", "", "GitHub owner or organization")
	flag.StringVar(&config.repo, "repo", "", "GitHub repository name")
	flag.StringVar(&config.description, "description", "", "Project description")
	flag.Parse()
	return config
}

func (config config) validate() error {
	if config.module == "" {
		return errors.New("--module is required")
	}
	if config.name == "" {
		return errors.New("--name is required")
	}
	if config.repo == "" {
		return errors.New("--repo is required")
	}
	if config.owner == "" {
		return errors.New("--owner is required")
	}
	if config.description == "" {
		return errors.New("--description is required")
	}
	if strings.ContainsAny(config.name, " /\\") {
		return errors.New("--name must be a binary-friendly name without spaces or slashes")
	}
	return nil
}

func (config config) needsPrompt() bool {
	return config.module == "" || config.name == "" || config.owner == "" || config.repo == "" || config.description == ""
}

func (config *config) promptMissing(input *os.File, output *os.File) error {
	reader := bufio.NewReader(input)

	var err error
	config.module, err = prompt(reader, output, "Go module path", config.module)
	if err != nil {
		return err
	}

	defaultName := config.name
	if defaultName == "" {
		defaultName = moduleName(config.module)
	}
	config.name, err = prompt(reader, output, "Binary name", defaultName)
	if err != nil {
		return err
	}

	defaultOwner := config.owner
	if defaultOwner == "" {
		defaultOwner = moduleOwner(config.module)
	}
	config.owner, err = prompt(reader, output, "GitHub owner", defaultOwner)
	if err != nil {
		return err
	}

	defaultRepo := config.repo
	if defaultRepo == "" {
		defaultRepo = config.name
	}
	config.repo, err = prompt(reader, output, "GitHub repo", defaultRepo)
	if err != nil {
		return err
	}

	config.description, err = prompt(reader, output, "Description", config.description)
	if err != nil {
		return err
	}

	return nil
}

func prompt(reader *bufio.Reader, output *os.File, label string, defaultValue string) (string, error) {
	if defaultValue == "" {
		fmt.Fprintf(output, "%s: ", label)
	} else {
		fmt.Fprintf(output, "%s [%s]: ", label, defaultValue)
	}

	value, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	value = strings.TrimSpace(value)
	if value == "" {
		return defaultValue, nil
	}
	return value, nil
}

func moduleName(module string) string {
	module = strings.TrimSuffix(module, "/")
	if module == "" {
		return ""
	}
	index := strings.LastIndex(module, "/")
	if index < 0 {
		return module
	}
	return module[index+1:]
}

func moduleOwner(module string) string {
	parts := strings.Split(strings.Trim(module, "/"), "/")
	if len(parts) >= 2 {
		return parts[len(parts)-2]
	}
	return ""
}

func shouldSkipDir(path string) bool {
	if hasPathSegment(path, "node_modules") {
		return true
	}
	switch path {
	case ".git", "dist", "bin", "docs/.vitepress/cache", "docs/.vitepress/dist":
		return true
	default:
		return false
	}
}

func shouldSkipFile(path string) bool {
	if hasPathSegment(path, "node_modules") {
		return true
	}
	switch filepath.Base(path) {
	case "go.sum", "pnpm-lock.yaml":
		return true
	default:
		return false
	}
}

func hasPathSegment(path string, segment string) bool {
	for _, part := range strings.Split(filepath.ToSlash(path), "/") {
		if part == segment {
			return true
		}
	}
	return false
}

type replacement struct {
	old string
	new string
}

func replaceInFile(path string, replacements []replacement) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	updated := string(content)
	for _, replacement := range replacements {
		updated = strings.ReplaceAll(updated, replacement.old, replacement.new)
	}
	if updated == string(content) {
		return nil
	}

	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	return os.WriteFile(path, []byte(updated), info.Mode())
}
