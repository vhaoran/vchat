package yred

import (
	"fmt"
)

//支持表名+主键的hashSet
func CacheKeyTableH(tbName string) string {
	return fmt.Sprintf("/%s/cache/table/%s", RedModuleName, tbName)
}
