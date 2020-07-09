/**
 * @date 2020/7/9 19:02
 * @Project_name yezihack
 */
package util

import "os"

// 判断是否是目录
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

// HumaneFileSize calculates the file size and generate user-friendly string.
func HumaneFileSize(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1024, sizes)
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
