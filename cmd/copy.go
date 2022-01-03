package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yezihack/saber/internal"
)

// command
var (
	source      string // 源目录
	target      string // 目标目录
	regular     string // 正则内容
	CopyCommand = &cobra.Command{
		Use:     "cp",
		Example: "saber cp --src /data-src --target /data-new --reg *.mp3",
		Short:   "检索您的目标文件",
		Long:    `将检索到的文件复制到指定目录`,
		Args: func(cmd *cobra.Command, args []string) error { // 检查 arguments 参数的
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) { // 这个 args 就是 arguments
			ok, err := internal.Copy(source, target, regular)
			if err != nil {
				fmt.Println("copy command is err:", err)
				os.Exit(1)
			}
			if !ok {
				fmt.Println("copy is failed!")
				os.Exit(1)
			}
		},
	}
)

func init() {
	CopyCommand.PersistentFlags().StringVar(&source, "src", ".", "源目录地址")
	CopyCommand.PersistentFlags().StringVar(&target, "desc", "./desc", "目标目录地址")
	CopyCommand.PersistentFlags().StringVar(&regular, "reg", "", "正则表达内容，如: log 或 (mp3|mp4)")
}
