package response

import (
	"fmt"
)

type RespError struct {
	Code string
	Msg  string
}

func (this RespError) Error() string {
	return this.Msg
}

func (this *RespError) String() string {
	return fmt.Sprintf("code:%s,msg:%s", this.Code, this.Msg)
}

//RegisterErrno 注册errno
func RegisterErrno(errno string, message []string) error {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}
	return &RespError{
		Code: errno,
		Msg:  msg,
	}
}
