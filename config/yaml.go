package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// 创建一个结构体
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Array    string `yaml:"array"`
	} `yaml:"database"`
}

func ReadYaml(path string) (*Config, error) {

	// read yaml
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
		return nil, err
	}

	// create config struct
	var config Config

	// parse yaml config
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalf("无法解析YAML: %v", err)
		return nil, err
	}

	return &config, nil
}

// 创建强制配置文件校验函数
func CheckConfig(config *Config) error {
	fmt.Printf("111")
	return nil
}
