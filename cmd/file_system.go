/**
 * @Author WANGZILIANG
 * @date 2020/7/7 19:24
 * @Project_name yezihack
 */
package cmd

import (
	"fmt"
	"os"

	"github.com/yezihack/saber/internal"

	"github.com/spf13/cobra"
	"github.com/unknwon/com"
)

// command
var (
	port       int // port
	FileSystem = &cobra.Command{
		Use:     "fs",
		Example: "saber fs /data/logs/ --port 7000",
		Short:   "File system",
		Long:    `The file system can browse your directory`,
		Args: func(cmd *cobra.Command, args []string) error { // 检查 arguments 参数的
			if len(args) < 1 {
				return fmt.Errorf("input your directory")
			}
			if !com.IsDir(args[0]) {
				return fmt.Errorf("%s is invalid directory", args[0])
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) { // 这个 args 就是 arguments
			err := internal.FileSystem(args[0], port)
			if err != nil {
				fmt.Println("file system err:", err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	FileSystem.PersistentFlags().IntVar(&port, "port", 7000, "port number")
}
