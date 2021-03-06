/**
 * Create Time:2022/3/12
 * User: luchao
 * Email: luc@shinemo.com
 */
package client

import (
	"../glog"
	"../util"
)

// 通知
var connChannel = make(chan bool, 3)

func Bridge() {
	for {
		<-connChannel
		<-connChannel
		go util.Bridge(inConn, serverConn)
		glog.Infof("Bridge success - conn1:%v, conn2:%v", serverConn.RemoteAddr(), inConn.RemoteAddr())
		util.Bridge(serverConn, inConn)
		InConnChannel <- true
		ServerConnChannel <- true
	}
}
