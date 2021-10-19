package utils

import "os"

//GetPid 获取服务进程id
func GetPid() int {
	pid := os.Getpid()
	return pid
}
