/**
 * Create Time:2021/3/31
 * User: luchao
 * Email: luc@shinemo.com
 */
package main

import (
	"../client"
	"../glog"
)

func main() {
	appName := "npf-client"

	glog.Launch("/data/logs", appName)
	glog.Infof("server starting - appName:%s", appName)

	serverIp := "127.0.0.1"

	serverPort := "8090"
	inIp := "127.0.0.1"
	inPort := "8091"

	go client.KeepServerConn(serverIp, serverPort)
	go client.KeepInConn(inIp, inPort)

	client.ServerConnChannel <- true
	client.InConnChannel <- true

	glog.Infof("server start finish - appName:%s", appName)
	client.Bridge()
}
