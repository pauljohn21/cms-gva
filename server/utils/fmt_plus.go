package utils

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func Pointer[T any](in T) (out *T) {
	return &in
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// MaheHump 将字符串转换为驼峰命名
func MaheHump(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

// 随机字符串
func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(0, len(letters))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func ConvertToChinese(amount string) string {
	digits := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	units := []string{"元", "角", "分"}

	// 解析金额字符串
	parsedAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "无效金额：" + amount
	}

	// 分解整数部分和小数部分
	intPart := math.Floor(parsedAmount)
	fracPart := parsedAmount - intPart

	// 转换整数部分
	intStr := strconv.FormatFloat(intPart, 'f', -1, 64)
	intResult := convertDigitsToInt(intStr, digits)

	// 转换小数部分
	fracStr := strconv.FormatFloat(fracPart*100, 'f', 2, 64)
	fracStr = strings.TrimPrefix(fracStr, "0.")

	// 转换小数部分为中文数字
	fracResult := convertDigitsToInt(fracStr, digits)

	// 合并结果
	result := intResult
	if len(fracResult) > 0 {
		result += units[len(fracResult)] + fracResult
	} else {
		result += "整"
	}

	return result
}

// convertDigitsToInt converts a string of digits to its Chinese character representation.
func convertDigitsToInt(digits string, digitsMap []string) string {
	var result strings.Builder

	for _, char := range digits {
		digit, _ := strconv.Atoi(string(char))
		if digit > 0 {
			result.WriteString(digitsMap[digit])
		} else if result.Len() > 0 {
			result.WriteString(digitsMap[digit])
		}
	}

	return result.String()
}
