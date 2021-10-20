package utils

import (
	"os"
	"path/filepath"
)

//GetPid 获取服务进程id
func GetPid() int {
	pid := os.Getpid()
	return pid
}

//GetDirFiles 获取指定目录下的文件
//eg: ./resource/errno  .toml
//return [resource\errno\common.toml,resource\errno\live\create\create.toml,resource\errno\live\live.toml]
func GetDirFiles(rootPath, ext string) []string {
	files := make([]string, 0)
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return files
}
