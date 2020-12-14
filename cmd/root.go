/**
 * @Author sgfoot.com
 * @date 2020/7/7 17:42
 * @Project_name yezihack
 */
package cmd

import (
	"log"
	"os"

	"github.com/yezihack/saber/entity"
	"github.com/yezihack/saber/internal"

	"github.com/spf13/cobra"
)

// 根
var (
	// 配置文件
	conf    = new(entity.Config)
	rootCmd = &cobra.Command{
		Use:   "saber",
		Short: "Saber toolkit",
		Long:  `Saber is an integrated toolbox`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&conf.Debug, "debug", "d", false, "print verbose log info")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(FileSystem)
	rootCmd.AddCommand(TcpProxy)
	rootCmd.AddCommand(Password)
}

func GetRoot() *cobra.Command {
	return rootCmd
}

// 执行
func Execute() {
	// 初使 配置文件
	internal.InitInternal(conf)
	if conf.Debug {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	} else {
		log.SetFlags(log.LstdFlags)
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}
