// Package helper
// @Author fuzengyao
// @Date 2022-11-10 09:39:26
package helper

import (
	"path"
	"runtime"
	"strings"
)

// CurrentFileName
// @Description: 当前文件名称
// @return string
func CurrentFileName() string {
	_, fullFilename, _, _ := runtime.Caller(1)
	return strings.Replace(path.Base(fullFilename), ".go", "", -1)
}
