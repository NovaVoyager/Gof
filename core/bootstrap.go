package core

import (
	"github.com/miaogu-go/Gof/core/initialize"
)

//Bootstrap 启动入口
func Bootstrap() {
	initialize.LoadConfig()
	initialize.LoadLog()
	initialize.LoadDb()
}
