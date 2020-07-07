/**
 * @Author sgfoot
 * @date 2020/7/7 19:38
 * @Project_name yezihack
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yezihack/saber/config"
)

// 版本
var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Version: config.Version,
	Short:   "Print the version number of Saber",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Saber version is " + config.Version)
	},
}
