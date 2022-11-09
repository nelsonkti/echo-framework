// Package api
// @Author fuzengyao
// @Date 2022-11-09 11:18:11
package config

import (
	config "github.com/nelsonkti/echo-framework/config/pb"
	_ "github.com/nelsonkti/echo-framework/util/xencoding/json"
	_ "github.com/nelsonkti/echo-framework/util/xencoding/yaml"
	"github.com/nelsonkti/echo-framework/util/xfile"
	"os"
)

const (
	configFile   = "/config.yaml"
	uploadPath   = "./public/upload"
	DownloadPath = "./public/download"
)

type fileFuc func() *xfile.FileInfo
type fuc func()
type pathFuc func() string

func containerPath(f pathFuc) string { return f() }

func init() {
	mappingConf(getConfigFile)
}

func mappingConf(f fileFuc) {

	var target map[string]interface{}

	err := unmarshal(f(), &target)
	if err != nil {
		panic(err)
	}

	res, err := marshalJSON(target)
	if err != nil {
		panic(err)
	}

	err = unmarshalJSON(res, &AppConf)
	if err != nil {
		panic(err)
	}

	appPath(defaultPath)
}

func getConfigFile() *xfile.FileInfo {
	path, err := os.Getwd()

	file, err := xfile.LoadFile(path + configFile)
	if err != nil {
		panic(err)
	}

	return file
}

func appPath(f fuc) {
	f()

	AppConf.App.Path = &config.App_Path{
		AppPath:      containerPath(defaultAppPath),
		UploadPath:   containerPath(defaultUploadPath),
		DownloadPath: containerPath(defaultDownloadPath),
		LogPath:      containerPath(defaultLogPath),
	}
}

func defaultPath() {
	if AppConf.App.GetPath() != nil {
		return
	}
	AppConf.App.Path = &config.App_Path{}
}

func defaultAppPath() string {
	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	if path := AppConf.App.Path.AppPath; path != "" {
		return path
	}

	return path
}

func defaultUploadPath() string {
	if path := AppConf.App.Path.UploadPath; path != "" {
		return path
	}

	return uploadPath
}

func defaultDownloadPath() string {

	if path := AppConf.App.Path.DownloadPath; path != "" {
		return path
	}

	return DownloadPath
}

func defaultLogPath() string {

	if path := AppConf.App.Path.LogPath; path != "" {
		return path
	}

	return AppConf.App.Path.AppPath
}
