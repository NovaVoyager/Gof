package utils

import (
	"fmt"
	"math/rand"
	"time"
)

//GenerateTraceId 生成traceId
func GenerateTraceId() string {
	random := GetRandRang(10000, 99999)
	timeNow := time.Now().Unix()
	pid := GetPid()
	return fmt.Sprintf("%d-%d-%d", timeNow, random, pid)
}

//GetRandRang 获取指定范围内的随机数
func GetRandRang(min, max int) int {
	if min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(max-min) + min
	return res
}
