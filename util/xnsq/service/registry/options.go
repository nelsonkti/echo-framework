// Package registry
// @Author fuzengyao
// @Date 2022-11-10 11:07:10
package registry

type Options struct {
	NsqAddress      string
	NSQConsumers    []string
	NSQAdminAddress string
	Env             string
	LocalAddress    string
}
