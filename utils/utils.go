package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strconv"
)

func Getenv(path, defVal string) string {
	result := os.Getenv(path)
	if len(result) == 0 {
		return defVal
	}
	return result
}

// 辅助函数
func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// string转int
func StringToInt(str string) int {
	if str == "" {
		return 0
	}
	return int(StringToInt32(str))
}

// string转int32
func StringToInt32(str string) int32 {
	if str == "" {
		return 0
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	t := int32(i)
	return t
}

// string转int64
func StringToInt64(str string) int64 {
	if str == "" {
		return 0
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	t := int64(i)
	return t
}

func StringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return f
}
