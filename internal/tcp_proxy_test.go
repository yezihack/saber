/**
 * @Author
 * @date 2020/7/9 16:04
 * @Project_name yezihack
 */
package internal

import (
	"testing"
)

func TestNewTcpProxy(t *testing.T) {
	NewTcpProxy("localhost:3306", "localhost:3307")
}
