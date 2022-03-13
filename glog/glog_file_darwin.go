// +build darwin

package glog

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

// 创建管道文件
func createPipeFile() (*os.File, string, error) {
	pipeFilePath := filepath.Join(*logDir, "pipe.fifo")
	os.Remove(pipeFilePath)
	err := syscall.Mkfifo(pipeFilePath, 0666)
	if err != nil {
		fmt.Println("create piple file fail - pipeFilePath=", pipeFilePath)
		return nil, pipeFilePath, err
	}
	file, err := os.OpenFile(pipeFilePath, os.O_RDWR, 0777)

	return file, pipeFilePath, err
}
