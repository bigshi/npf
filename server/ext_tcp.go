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
)

// 监听外部请求
func ListenExtTcp(port int) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("0.0.0.0"), Port: port})
	if nil != err {
		glog.Errorf("ListenExtTcp listen fail - port:%d, err:%v", port, err)
		return
	}
	defer listener.Close()
	glog.Infof("ListenExtTcp listen success - port:%d", port)
	for {
		glog.Infof("ListenExtTcp accept pending - port:%d", port)
		conn, err := listener.AcceptTCP()
		if nil != err {
			glog.Errorf("ListenExtTcp accept fail - port:%d, err:%v", port, err)
			continue
		}
		glog.Infof("ListenExtTcp accept success - localAddr:%v, remoteAddr:%v", conn.LocalAddr(), conn.RemoteAddr())
		go handleExtConn(conn)
	}
}

// 处理外部连接
func handleExtConn(conn *net.TCPConn) {
	//conn.SetDeadline(time.Now().Add(3 * time.Second))
	conn.SetKeepAlive(true)
	remoteAddr := conn.RemoteAddr()
	clientConn, isOk := clientConnCache.Get(clientConnEqualizer.GetNextKey())
	if !isOk {
		glog.Errorf("handleExtConn fail - extConn:%v, msg:clientConn is nil", remoteAddr)
		conn.Close()
		return
	}
	go util.Bridge(conn, clientConn)
	go util.Bridge(clientConn, conn)
	glog.Infof("handleExtConn success - extConn:%v, clientConn:%v", remoteAddr, clientConn.RemoteAddr())
}
