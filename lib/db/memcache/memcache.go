package memcache

import "github.com/bradfitz/gomemcache/memcache"

var Memcache *memcache.Client

// Connect
// @Description: 连接memcache
// @param server
// @return *memcache.Client
func Connect(server []string) *memcache.Client {
	Memcache = memcache.New(server...)
	return Memcache
}
