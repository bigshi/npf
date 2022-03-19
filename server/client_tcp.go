/**
 * Create Time:2022/2/28
 * User: luchao
 * Email: luc@shinemo.com
 */
package server

import (
	"../glog"
	"net"
	"strings"
	"time"
)

var whiteIps = []string{"39.170.35.150", "112.16.91.49", "127.0.0.1"}

// 监听客户端连接
func ListenClientTcp(port int) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: port})
	if nil != err {
		glog.Errorf("ListenClientTcp listen fail - port:%d, err:%v", port, err)
		return
	}
	defer listener.Close()
	glog.Infof("ListenClientTcp listen success - port:%d", port)
	for {
		glog.Infof("ListenClientTcp accept pending - port:%d", port)
		conn, err := listener.AcceptTCP()
		if nil != err {
			glog.Errorf("ListenClientTcp accept fail - port:%d, err:%v", port, err)
			continue
		}
		remoteAddr := conn.RemoteAddr().String()
		var isOk bool
		for _, s := range whiteIps {
			if strings.Contains(remoteAddr, s) {
				isOk = true
			}
		}
		if !isOk {
			glog.Errorf("ListenClientTcp accept fail - msg:no auth ip,localAddr:%v, remoteAddr:%v", conn.LocalAddr(), remoteAddr)
			continue
		}
		glog.Infof("ListenClientTcp accept success - localAddr:%v, remoteAddr:%v", conn.LocalAddr(), remoteAddr)
		go handleClientConn(conn)
	}
}

// 处理客户端连接
func handleClientConn(conn *net.TCPConn) {
	localAddr := conn.LocalAddr()
	remoteAddr := conn.RemoteAddr()
	key := remoteAddr.String()
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	receiveMsg := make([]byte, 1024)
	n, err := conn.Read(receiveMsg)
	if err != nil {
		glog.Errorf("handleClientConn fail - localAddr:%v, remoteAddr:%v, err:%v", localAddr, remoteAddr, err)
		conn.Close()
		return
	}
	msg := string(receiveMsg[:n])
	if msg != "I'm Client" {
		glog.Errorf("handleClientConn fail - msg:receiveMsg is error, localAddr:%v, remoteAddr:%v, receiveMsg:%s", localAddr, remoteAddr, err, msg)
		conn.Close()
		return
	}
	conn.SetKeepAlive(true)
	oldConn, isOk := clientConnCache.Add(key, conn)
	if !isOk {
		oldConn.Close()
	}
	clientConnEqualizer.Add(key)
	glog.Infof("handleClientConn success - localAddr:%v, remoteAddr:%v", localAddr, remoteAddr)
}
