/**
 * Create Time:2022/2/28
 * User: luchao
 * Email: luc@shinemo.com
 */
package server

import (
	"../glog"
	"../util"
	"net"
	"strings"
)

var whiteIps = []string{"127.0.0.1"}

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
				break
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
	key := util.GetIp(remoteAddr)
	conn.SetKeepAlive(true)
	oldConn, isOk := clientConnCache.Add(key, conn)
	if !isOk {
		oldConn.Close()
	}
	clientConnEqualizer.Add(key)
	glog.Infof("handleClientConn success - localAddr:%v, remoteAddr:%v", localAddr, remoteAddr)
}
