package core

import (
	"github.com/miaogu-go/Gof/core/initialize"
)

func Bootstrap() {
	initialize.LoadConfig()
	initialize.LoadLog()
	initialize.LoadDb()
}
