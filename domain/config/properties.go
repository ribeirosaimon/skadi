package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magiconair/properties"
)

func GetProperties(arg string) properties.Properties {
	dir, _ := FindRootDir()
	config := fmt.Sprintf("config.%s.properties", arg)
	return *properties.MustLoadFile(fmt.Sprintf("%s/%s", dir, config), properties.UTF8)
}

func FindRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("project root not found")
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)

		if parentDir == currentDir {
			break
		}

		currentDir = parentDir
	}

	return "", fmt.Errorf("project root not found")
}
