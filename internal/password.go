/**
 * @date 2020/7/9 19:19
 * @Project_name yezihack
 */
package internal

import (
	"math/rand"
	"time"
)

var (
	numbers    = "0123456789"
	chars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charNumber = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	mix        = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+=-{}|'',./"
)

func RandomStrs(length, count int, cate string) []string {
	result := make([]string, count)
	for i := 0; i < count; i++ {
		result[i] = RandomStr(length, cate)
	}
	return result
}

func RandomStr(length int, cate string) string {
	var (
		bytes  = make([]byte, 0)
		result = make([]byte, length)
		r      = rand.New(rand.NewSource(time.Now().UnixNano()))
	)
	switch cate {
	case "num":
		bytes = append(bytes, []byte(numbers)...)
	case "char":
		bytes = append(bytes, []byte(chars)...)
	case "num-char":
		bytes = append(bytes, []byte(charNumber)...)
	default:
		bytes = append(bytes, []byte(mix)...)
	}
	for i := 0; i < length; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}
