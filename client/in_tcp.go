/**
 * Create Time:2022/3/6
 * User: luchao
 * Email: luc@shinemo.com
 */
package client

import (
	"../glog"
	"net"
	"time"
)

var inConn *net.TCPConn

// 通知
var InConnChannel = make(chan bool, 2)

// 建立内网业务连接
func makeInConn(ip string, port string) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ip+":"+port)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		glog.Errorf("MakeInConn fail - ip:%s, port:%s, err:%v", ip, port, err)
		InConnChannel <- true
		return
	}
	inConn = conn
	//conn.SetDeadline(time.Now().Add(3 * time.Second))
	conn.SetKeepAlive(true)
	connChannel <- true
	glog.Infof("MakeInConn success -  localAddr:%v, remoteAddr:%v", conn.LocalAddr(), conn.RemoteAddr())
}

func KeepInConn(ip string, port string) {
	for {
		<-InConnChannel
		time.Sleep(3 * time.Second)
		makeInConn(ip, port)
	}
}
