package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magiconair/properties"
)

func GetPropertiesFile(arg string) *properties.Properties {
	dir, _ := FindRootDir()
	config := fmt.Sprintf("config.%s.properties", arg)
	return properties.MustLoadFile(fmt.Sprintf("%s/%s", dir, config), properties.UTF8)
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

func FindModuleDir(module string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("project root not found")
	}
	index := strings.Index(currentDir, "/skadiEngine/")
	if index == -1 {
		return "", fmt.Errorf("'/skadiEngine/' not found in the path")
	}
	result := currentDir[index+len("/skadiEngine/"):]

	replace := strings.Replace(currentDir, result, module, -1)
	validatePath(replace)
	return replace, nil
}

func validatePath(directoryPath string) {
	cleanedPath := filepath.Clean(directoryPath)

	_, err := os.Stat(cleanedPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("The dir %s not exist.\n", cleanedPath)
		}
	}
}
