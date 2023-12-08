package utils

import (
	"golang.org/x/exp/constraints"
	"path/filepath"
	"regexp"
	"strings"
)

// GetFileExt 获取文件的扩展名
func GetFileExt(filename string) string {
	ext := strings.TrimPrefix(filepath.Ext(filename), ".")
	return strings.ToLower(ext)
}

// GetMissingIds 获取缺失的 ID 列表
func GetMissingIds[T any](allIds []uint, presentIds map[uint]T) []uint {
	var missingIds []uint

	for _, id := range allIds {
		if _, ok := presentIds[id]; !ok {
			missingIds = append(missingIds, id)
		}
	}

	return missingIds
}

// Deduplicate 去重复
func Deduplicate[T constraints.Ordered](slice []T) []T {
	encountered := make(map[T]struct{})
	result := []T{}

	for _, v := range slice {
		if _, ok := encountered[v]; !ok {
			encountered[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// ToSnakeCase 驼峰转下划线
func ToSnakeCase(input string) string {
	// 使用正则表达式将驼峰命名法转换为下划线连接的小写单词
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snakeCase := re.ReplaceAllString(input, "${1}_${2}")

	// 将所有字母转换为小写
	snakeCase = strings.ToLower(snakeCase)

	return snakeCase
}

// MaskPhoneNumber 手机号码掩码
func MaskPhoneNumber(phoneNumber string) (maskedNumber string) {
	match, _ := regexp.MatchString(`^\d{11}$`, phoneNumber)
	if !match {
		return "****"
	}
	maskedNumber = phoneNumber[:3] + "****" + phoneNumber[7:]
	return maskedNumber
}
