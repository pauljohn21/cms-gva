package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
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

// FormatTime 将给定的时间字符串格式化为 "2024年10月03日零时" 格式
func FormatTime(originalTime *time.Time) (string, error) {
	if originalTime == nil {
		return "", fmt.Errorf("originalTime is nil")
	}

	// 转换到本地时区
	localTime := originalTime.In(time.Local)

	// 获取零时的时间
	zeroTime := time.Date(localTime.Year(), localTime.Month(), localTime.Day(), 0, 0, 0, 0, localTime.Location())

	// 格式化时间
	formattedTime := zeroTime.Format("2006年01月02日零时")

	return formattedTime, nil
}

func GenerateOrderID(sequence int) string {
	// 获取当前时间
	currentTime := time.Now()

	// 获取日期部分
	datePart := currentTime.Format("20060102")

	// 获取时间部分
	timePart := currentTime.Format("150405")

	// 格式化序列号
	sequencePart := fmt.Sprintf("%06d", sequence)

	// 组合订单编号
	orderID := fmt.Sprintf("P1100C024%s%s%s", datePart, timePart, sequencePart)

	return orderID
}
