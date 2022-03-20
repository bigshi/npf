/**
 * Create Time:2022/3/12
 * User: luchao
 * Email: luc@shinemo.com
 */
package server

import (
	"net"
	"sync"
)

// 客户端连接缓存
type ClientConnCache struct {
	ConnMap map[string]*net.TCPConn // key：远端地址
	Lock    sync.RWMutex
}

// 实例化
var clientConnCache = ClientConnCache{
	ConnMap: make(map[string]*net.TCPConn, 0),
	Lock:    sync.RWMutex{},
}

// 添加客户端连接
func (cache *ClientConnCache) Add(key string, conn *net.TCPConn) (*net.TCPConn, bool) {
	cache.Lock.Lock()
	defer cache.Lock.Unlock()
	oldConn, isOk := cache.ConnMap[key]
	cache.ConnMap[key] = conn
	// 已经存在
	if isOk {
		return oldConn, false
	}
	return nil, true
}

// 删除客户端连接
func (cache *ClientConnCache) Del(key string) {
	cache.Lock.Lock()
	defer cache.Lock.Unlock()
	delete(cache.ConnMap, key)
}

// 获取客户端连接
func (cache *ClientConnCache) Get(key string) (*net.TCPConn, bool) {
	cache.Lock.RLock()
	defer cache.Lock.RUnlock()
	conn, isOk := cache.ConnMap[key]
	return conn, isOk
}

// 判断客户端连接是否存在
func (cache *ClientConnCache) Has(key string) bool {
	cache.Lock.Lock()
	defer cache.Lock.Unlock()
	_, isOk := cache.ConnMap[key]
	return isOk
}
