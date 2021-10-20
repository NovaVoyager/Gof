package errno

import (
	"github.com/miaogu-go/Gof/utils/response"
)

var (
	NotFountErr = response.RegisterErrno("10001", []string{"not found"})
)
