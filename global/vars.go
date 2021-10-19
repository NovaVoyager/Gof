package global

import "time"

var (
	nowTime = func() time.Time {
		return time.Now().In(time.FixedZone("CST", 8*3600))
	}
)
