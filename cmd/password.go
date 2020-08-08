/**
 * 随机密码
 * @date 2020/7/9 19:13
 * @Project_name yezihack
 */
package cmd

import (
	"fmt"
	"strconv"

	"github.com/unknwon/com"

	"github.com/spf13/cobra"
	"github.com/yezihack/saber/internal"
)

var (
	length   int    // 生成密码多少长度
	typeName string // 生成密码的类型. 可以是纯数字, 字符, 字符和数字, 混合复杂的
	filename string // 定向到文件里去.
	Password = &cobra.Command{
		Use:     "pass",
		Example: "saber pass",
		Short:   "Random Password",
		Long:    `Random generate to password`,
		Args: func(cmd *cobra.Command, args []string) error { // 检查 arguments 参数的
			if len(args) > 0 {
				if _, err := strconv.Atoi(args[0]); err != nil {
					return fmt.Errorf("%s is invalid number", args[0])
				}
			}
			if typeName != "" {
				if !com.IsSliceContainsStr([]string{"num", "char", "num-char", "mix"}, typeName) {
					return fmt.Errorf("password type select:%s", "num, char, num-char, mix. num-char is default")
				}
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) { // 这个 args 就是 arguments
			count := 1
			if len(args) > 0 {
				count, _ = strconv.Atoi(args[0])
			}
			result := internal.RandomStrs(length, count, typeName)
			for i := 0; i < len(result); i++ {
				println(result[i])
			}
		},
	}
)

func init() {
	Password.PersistentFlags().IntVarP(&length, "len", "l", 12, "Password length")
	Password.PersistentFlags().StringVarP(&typeName, "type", "t", "num-char", "Password type: num, char, num-char, mix")
	Password.PersistentFlags().StringVarP(&filename, "filename", "f", "password.txt", "Output collected data to file")
}
