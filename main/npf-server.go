/**
 * Create Time:2021/3/31
 * User: luchao
 * Email: luc@shinemo.com
 */
// 账号中心-核心层
package main

import (
	"../glog"
	"../server"
	"time"
)

// 基础中心 - 核心层
func main() {
	appName := "npf-server"

	glog.Launch("/data/logs", appName)
	glog.Infof("server starting - appName:%s", appName)

	go server.ListenClientTcp(8090)
	go server.ListenExtTcp(8080)

	time.Sleep(3 * time.Second)
	glog.Infof("server start finish - appName:%s", appName)
	for {
		time.Sleep(3 * time.Second)
	}

}
