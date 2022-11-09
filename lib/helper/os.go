// Package api
// @Author fuzengyao
// @Date 2022-11-09 11:18:11
package helper

import (
	"path"
	"runtime"
	"strings"
)

// 当前文件名称
func CurrentFileName() string {
	_, fullFilename, _, _ := runtime.Caller(1)
	return strings.Replace(path.Base(fullFilename), ".go", "", -1)
}
