package utils

import (
	"fmt"
	"time"
)

//TimeCost 耗时统计
func TimeCost() func() string {
	start := time.Now()
	return func() string {
		tc := time.Since(start)
		return fmt.Sprintf("%s", tc)
	}
}
