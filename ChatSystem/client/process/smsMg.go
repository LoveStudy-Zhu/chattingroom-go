package process

import (
	common "ChatSystem/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *common.Message){
	//1.反序列化mes.Data
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal err =",err.Error())
		return
	}
	info := fmt.Sprintf("用户id :\t%d对大家说\t%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
