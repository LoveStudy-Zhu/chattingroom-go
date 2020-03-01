package process

import (
	"ChatSystem/client/model"
	common "ChatSystem/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsers map[int] *common.User =make(map[int]*common.User,10)
var CurUser model.CurUser



//显示当前在线的用户
func outputOnlineUser(){
	fmt.Println("当前在线用户列表")
	for id ,_ := range  onlineUsers{
		//如果不显示自己
		fmt.Println("用户id\t",id)
	}
}

func updateUserStatus(notifyUserStatusMes *common.NotifyUserStatusMes){
	user,ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok{
		user = &common.User{
			UserId:     notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId]=user
	outputOnlineUser()
}


