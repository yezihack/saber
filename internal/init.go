/**
 * @date 2020/7/9 16:09
 * @Project_name yezihack
 */
package internal

import (
	"github.com/yezihack/saber/entity"
)

var (
	__conf *entity.Config
)

func InitInternal(conf *entity.Config) {
	__conf = conf
}
