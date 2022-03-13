/**
 * Create Time:2022/3/7
 * User: luchao
 * Email: luc@shinemo.com
 */
package util

import (
	"../glog"
	"io"
	"net"
)

func Bridge(conn1 *net.TCPConn, conn2 *net.TCPConn) {
	defer conn1.Close()
	defer conn2.Close()
	// 使用io.Copy传输两个tcp连接，
	_, err := io.Copy(conn1, conn2)
	if err != nil {
		glog.Errorf("Bridge fail - conn1:%v, conn2:%v, err:%v", conn1.RemoteAddr(), conn2.RemoteAddr(), err)
		return
	}
}
