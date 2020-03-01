package process2

import (
	common "ChatSystem/common/message"
	"ChatSystem/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {

}
//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *common.Message){
	//遍历服务端的在线用户
	//将消息转发出去

	//取出mes额内容SmsMes
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal err =",err)
	}
	data ,err := json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err ",err)
		return
	}
	for id,up := range userMgr.onlineUsers{
		if id == smsMes.UserId{
			continue
		}
		this.SengMesToEachOnlineUser(data,up.Conn)
	}
}

func (this *SmsProcess) SengMesToEachOnlineUser (data []byte,conn net.Conn){
	//创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err !=nil{
		fmt.Println("转发消息失败,err=",err)
	}

}