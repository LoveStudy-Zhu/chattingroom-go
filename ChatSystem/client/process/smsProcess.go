package process

import (
	"ChatSystem/client/utils"
	common "ChatSystem/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}
//发送群聊消息
func (this *SmsProcess) SendGroupMes(content string)(err error) {
	//1.创建一个mes
	var mes common.Message
	mes.Type = common.SmsMesType

	var smsMes common.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data ,err := json.Marshal(smsMes)
	if err != nil{
		fmt.Println("sendGroupMes json.Marshal fail =",err.Error())
		return
	}
	mes.Data = string(data)

	data ,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("sendGroupMes json.Marshal fail =",err.Error())
		return
	}
	//将序列化后的mes发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("sendGroupMes err =",err.Error())
		return
	}
	return
}