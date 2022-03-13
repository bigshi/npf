/**
 * Create Time:2022/3/12
 * User: luchao
 * Email: luc@shinemo.com
 */
package server

import (
	"sync"
)

// 连接负载均衡器
type ClientConnEqualizer struct {
	ConnList  []string // 数组 远端地址
	NextIndex int      // 数组下标 下一个连接
	Lock      sync.RWMutex
}

// 实例化
var clientConnEqualizer = ClientConnEqualizer{
	ConnList:  make([]string, 0),
	NextIndex: 0,
	Lock:      sync.RWMutex{},
}

// 添加轮询数组
func (equalizer *ClientConnEqualizer) Add(key string) {
	equalizer.Lock.Lock()
	defer equalizer.Lock.Unlock()
	equalizer.ConnList = append(equalizer.ConnList, key)
}

// 删除轮询数组
func (equalizer *ClientConnEqualizer) Del(key string) {
	equalizer.Lock.Lock()
	defer equalizer.Lock.Unlock()
	for i, k := range equalizer.ConnList {
		if k == key {
			equalizer.ConnList[i] = ""
			break
		}
	}
}

// 轮询获取下一个连接地址
func (equalizer *ClientConnEqualizer) GetNextKey() string {
	equalizer.Lock.RLock()
	defer equalizer.Lock.RUnlock()
	var key string
	// 取出长度 用于计算是否全部遍历过
	length := len(equalizer.ConnList)
	if length == 0 {
		return key
	}
	size := length
	for {
		key = equalizer.ConnList[equalizer.NextIndex]
		// 遍历过
		length--
		// 下一个节点
		if equalizer.NextIndex == size-1 {
			equalizer.NextIndex = 0
		} else {
			equalizer.NextIndex++
		}

		if key != "" {
			break
		}
		// 全部遍历过
		if length == 0 {
			return key
		}
	}
	return key
}
