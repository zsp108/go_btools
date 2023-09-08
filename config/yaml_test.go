package config

import (
	"fmt"
	"testing"
)

var configPath = "./config.yaml"

// 创建ReadYaml的测试用例
func TestReadYaml(t *testing.T) {
	config, err := ReadYaml(configPath)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Test Config Data:%+v\n", config)
	fmt.Printf("Array : %+v \n", config.Database.Array)
	// fmt.Printf("%+v\n", config.GetServer())
}
