/**
** @创建时间 : 2022/1/5 11:17
** @作者 : fzy
 */
package config

import (
	config "echo-framework/config/pb"
	_ "echo-framework/util/xencoding/json"
	_ "echo-framework/util/xencoding/yaml"
	"echo-framework/util/xfile"
	"os"
)

const (
	configFile   = "/config.yaml"
	uploadPath   = "./public/upload"
	DownloadPath = "./public/download"
)

var AppConf config.Conf

type Config struct {
	file *xfile.FileInfo
}

func init() {
	initConf(initFile()).initAppPath()
}

func initFile() *xfile.FileInfo {
	path, err := os.Getwd()

	file, err := xfile.LoadFile(path + configFile)
	if err != nil {
		panic(err)
	}

	return file
}

func initConf(file *xfile.FileInfo) *Config {
	var target map[string]interface{}

	err := unmarshal(file, &target)
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

	return &Config{}
}

func (c *Config) initAppPath() (Config *Config) {

	if path := AppConf.App.Path; path == nil {
		path, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		AppConf.App.Path = &config.App_Path{
			AppPath:      path,
			UploadPath:   uploadPath,
			DownloadPath: DownloadPath,
		}

		return
	}

	if path := AppConf.App.Path.AppPath; path == "" {
		path, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		AppConf.App.Path.AppPath = path
	}

	if path := AppConf.App.Path.UploadPath; path == "" {
		AppConf.App.Path.UploadPath = uploadPath
	}

	if path := AppConf.App.Path.DownloadPath; path == "" {
		AppConf.App.Path.DownloadPath = DownloadPath
	}

	return
}
