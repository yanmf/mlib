package utils

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var (
	gConfig interface{} // 全局配置变量
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

	return config, nil
}
