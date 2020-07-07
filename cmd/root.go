/**
 * @Author sgfoot.com
 * @date 2020/7/7 17:42
 * @Project_name yezihack
 */
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// 根
var rootCmd = &cobra.Command{
	Use: "saber",
	// Example: "saber fs /data/logs/ --port 7000",
	Short: "Saber toolkit",
	Long:  `Saber is an integrated toolbox`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(FileSystem)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		/*		fmt.Println(err)*/
		os.Exit(1)
	}
}
