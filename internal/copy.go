package internal

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/unknwon/com"
	"github.com/yezihack/saber/util"
)

type copyFiles struct {
	index      int    // 序号
	sourceFile string // 源文件
	targetFile string // 目标文件
}

// 复制目录下的所有文件
func Copy(source, dest, regular string) (ok bool, err error) {
	// 创建目标目录
	if !com.IsExist(dest) {
		err = os.MkdirAll(dest, os.ModePerm)
		if err != nil {
			return
		}
	}
	// 过滤后结果集
	fileArray := make(map[string]*copyFiles, 0)
	// 正则表达，如检索mp3文件，采用： --reg "mp3", 如果即想要 mp3 和 mp4 文件，采用： --reg ".(mp3|mp4)"
	reg := regexp.MustCompile(regular)
	i := 0
	filepath.Walk(source, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && reg.MatchString(path) {
			i++
			fileArray[path] = &copyFiles{
				index:      i,
				sourceFile: path,
				targetFile: util.AddSuffix(dest, "/") + info.Name(),
			}

			// fmt.Println("name:", path, info.Size())
		}
		return nil
	})

	// 复制文件

	success := 0 // 复制成功多少个
	for _, item := range fileArray {
		err = com.Copy(item.sourceFile, item.targetFile)
		if err == nil {
			success++
		} else {
			log.Printf("copy failed. index: %d, source-files: %s,  err:%v\n", item.index, item.sourceFile, err)
		}
	}
	fmt.Printf("共:%d 个文件，成功复制:%d 个，失败:%d个!", len(fileArray), success, len(fileArray)-int(success))
	return true, nil
}
