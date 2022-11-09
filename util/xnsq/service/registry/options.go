// Package api
// @Author fuzengyao
// @Date 2022-11-09 11:18:11
package registry

type Options struct {
	NsqAddress      string
	NSQConsumers    []string
	NSQAdminAddress string
	Env             string
	LocalAddress    string
}
