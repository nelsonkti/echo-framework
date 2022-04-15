/**
** @创建时间 : 2022/3/22 17:38
** @作者 : fzy
 */
package producer

import (
	"github.com/nelsonkti/echo-framework/lib/db"
	string2 "github.com/nelsonkti/echo-framework/util/xnsq/util/string"
)

// set ongoing ip
func SetOngoingIp(ip string) {
	key := "yimLive:server:ongoing:ip:addr"
	db.RedisDefault().HSet(key, ip, ip)
}

// get ongoing ip
func GetOngoingIp() []string {
	key := "yimLive:server:ongoing:ip:addr"
	return db.RedisDefault().HVals(key).Val()
}

// del ongoing ip
func DelOngoingIp(ip string) {
	key := "yimLive:server:ongoing:ip:addr"
	db.RedisDefault().HDel(key, ip)
}

// set ended ip
func SetEndedIp(ip string) {
	key := "yimLive:server:ended:ip:addr"
	db.RedisDefault().HSet(key, ip, ip)
}

// get ended ip
func GetEndedIp() []string {
	key := "yimLive:server:ended:ip:addr"
	return db.RedisDefault().HVals(key).Val()
}

// del ended ip
func DelEndedIp(ip string) {
	key := "yimLive:server:ended:ip:addr"
	db.RedisDefault().HDel(key, ip)
}

// 获取差集
func GetEndedOldTp() []string {

	newIp := GetOngoingIp()
	oldIp := GetEndedIp()
	if len(newIp) == 0 || len(oldIp) == 0 {
		return nil
	}

	return string2.Difference(oldIp, newIp)
}
