/**
 * @date 2020/7/9 19:28
 * @Project_name yezihack
 */
package internal

import "testing"

func TestRandomStr(t *testing.T) {
	str := RandomStr(12)
	println(str)
}
