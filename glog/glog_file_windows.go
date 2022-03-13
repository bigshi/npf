// +build windows

package glog

import (
	"errors"
	"os"
)

// 创建管道文件
func createPipeFile() (*os.File, string, error) {
	return nil, "", errors.New("windows not support")
}
