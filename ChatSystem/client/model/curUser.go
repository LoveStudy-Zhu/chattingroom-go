package model

import (
	common "ChatSystem/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	common.User
}
