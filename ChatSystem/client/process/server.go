package process

import (
	"ChatSystem/client/utils"
	common "ChatSystem/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面
func ShowMenu() {
	fmt.Println("--------恭喜xxx登录成功-------")
	fmt.Println("--------1.显示在线列表--------")
	fmt.Println("--------2.发送消息-----------")
	fmt.Println("--------3.消息列表-----------")
	fmt.Println("--------4.退出系统-----------")
	fmt.Println("请选择(1-4):")
	var key int
	var content string //发送的内容

	smsProcess := &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:

		fmt.Print("请输入你想输入的话:)")
		fmt.Scanf("%s", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择了退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的不正确..")
	}
}

//保持和服务器通信
func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		//fmt.Printf("客户端%s正在等待读取服务器发送的消息",conn)
		fmt.Printf("客户端正在等待读取服务器发送的消息\n")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("if Readpkg err=", err)
			return
		}
		//读取成功
		switch mes.Type {
		case common.NotifyUserStatusMesType: //有人上线了
			var notifyUserStatusMes common.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		//处理
		case common.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回了未知的消息类型")
		}
	}
}
