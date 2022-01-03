/**
 * @date 2020/7/9 19:02
 * @Project_name yezihack
 */
package util

import (
	"os"
)

// 判断是否是目录
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// FileMTime returns file modified time and possible error.
func FileMTime(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

// FileSize returns file size in bytes and possible error.
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

// 判断结尾处是否是以某字符结尾。
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 最后添加某字符，如何没有则添加，有则不添加。
func AddSuffix(s, suffix string) string {
	if !HasSuffix(s, suffix) {
		return s + suffix
	}
	return s
}
