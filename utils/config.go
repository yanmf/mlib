package utils

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var (
	gConfig     interface{} // 全局配置变量
	gConfigPath = "../config.toml"
)

// LoadConfig 加载配置文件的通用函数
func LoadConfig[T any](configPath string) (*T, error) {
	config := new(T)
	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("read config file failed: %w", err)
	}
	// 解析TOML
	if err := toml.Unmarshal(data, config); err != nil {
		return config, fmt.Errorf("parse toml config failed: %w", err)
	}

	gConfig = &config
	gConfigPath = configPath

	return config, nil
}

func SaveConfig[T any](cfg T) error {
	cfgBody, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}
	gConfig = cfg
	os.Rename(gConfigPath, fmt.Sprint(gConfigPath, ".", Second()))

	os.WriteFile(gConfigPath, cfgBody, 0755)

	return nil
}
