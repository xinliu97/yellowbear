package utils

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Mongo struct {
			Uri           string `yaml:"uri"`
			Sample1quizFp string `yaml:"sample_1quiz_fp"`
			Sample2quizFp string `yaml:"sample_2quiz_fp"`
			SampleNquizFp string `yaml:"sample_Nquiz_fp"`
		} `yaml:"mongo"`
	} `yaml:"database"`
}

func ReadConfigYaml() (*Config, error) {
	projRootPath := os.Getenv("YELLOWBEAR_ROOT")
	if projRootPath == "" {
		fmt.Println("Please set YELLOWBEAR_ROOT environment variable to YellowBear's root path.")
		return nil, errors.New("env not found")
	}
	configFp := filepath.Join(projRootPath, "config.yaml")
	yamlFile, err := os.ReadFile(configFp)
	if err != nil {
		fmt.Println("[ReadConfigYaml]", err)
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("[ReadConfigYaml]", err)
		return nil, err
	}

	config.Database.Mongo.Sample1quizFp = filepath.Join(projRootPath, config.Database.Mongo.Sample1quizFp)
	config.Database.Mongo.Sample2quizFp = filepath.Join(projRootPath, config.Database.Mongo.Sample2quizFp)
	config.Database.Mongo.SampleNquizFp = filepath.Join(projRootPath, config.Database.Mongo.SampleNquizFp)

	return &config, nil
}
