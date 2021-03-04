package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Dsn     string `yaml:"dsn"`
		SqlType string `yaml:"sqlType"`
	}
	Server struct {
		Port uint32 `yaml:"port"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	}
}

func ParseConfig(filepath string) *Config {
	config := Config{}
	t, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("failed to parse config")
		os.Exit(0)
	}
	//fmt.Println(yamlFile)
	err = yaml.UnmarshalStrict(t, &config)

	if err != nil {
		log.Fatalln(err)
	}

	return &config
}
