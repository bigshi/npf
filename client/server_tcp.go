/**
 * Create Time:2022/2/28
 * User: luchao
 * Email: luc@shinemo.com
 */
package client

import (
	"../glog"
	"net"
	"time"
)

var serverConn *net.TCPConn

// 通知
var ServerConnChannel = make(chan bool, 2)

// 建立服务端连接
func makeServerConn(ip string, port string) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ip+":"+port)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		glog.Errorf("MakeServerConn fail - ip:%s, port:%s, err:%v", ip, port, err)
		ServerConnChannel <- true
		return
	}
	serverConn = conn
	conn.SetKeepAlive(true)
	connChannel <- true
	glog.Infof("MakeServerConn success -  localAddr:%v, remoteAddr:%v", conn.LocalAddr(), conn.RemoteAddr())
}

func KeepServerConn(ip string, port string) {
	for {
		<-ServerConnChannel
		time.Sleep(3 * time.Second)
		makeServerConn(ip, port)
	}
}
