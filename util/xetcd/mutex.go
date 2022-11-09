// Package xetcd
// @Author fuzengyao
// @Date 2022-11-09 11:17:43
package xetcd

import (
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"sync"
)

func Locker(key string) sync.Locker {
	session, err := concurrency.NewSession(Client)
	if err != nil {
		log.Println("session err :", err)
	}

	locker := concurrency.NewLocker(session, Pfx+key)

	return locker
}
