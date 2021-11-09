package configs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

const configFile string = "main.yml"

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"http"`
	Database struct {
		Driver     string `yaml:"driver"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Host       string `yaml:"host"`
		DockerHost string `yaml:"docker_host"`
		Port       string `yaml:"port"`
		DbName     string `yaml:"dbname"`
		SSLMode    string `yaml:"sslmode"`
		DbUrl      string `yaml:"url"`
	} `yaml:"db"`
}

func GetConfig() *Config {
	var cfg Config
	readFile(&cfg)
	return &cfg
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open(configFile)
	if err != nil {
		processError(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			logrus.Error("failed to close file: %s", err.Error())
		}
	}()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
