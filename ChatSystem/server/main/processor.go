package main

import (
	common "ChatSystem/common/message"
	process2 "ChatSystem/server/process"
	"ChatSystem/server/utils"
	"io"
	"fmt"
	"net"
)

type Processsor struct {
	Conn net.Conn
}


//根据客户端发送消息的种类决定调用哪个函数来处理
func (this *Processsor)serverProcessMes(mes *common.Message)(err error){

	//客户端发送的消息
	fmt.Println("mes=",mes)

	switch mes.Type {
	case common.LoginMesType:
		//处理登录
		up := &process2.UserProcess{
			Conn:this.Conn,
		}

		err=up.ServerProcessLogin(mes)
	case common.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn:this.Conn,
		}

		err=up.ServerProcessRegister(mes)
	case common.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}
func (this *Processsor) startProcess() (err error){
	//循环读取客户端发送的信息
	for{

		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes,err := tf.ReadPkg()
		if err !=nil{
			if err ==io.EOF{
				fmt.Println("客户端退出")
				return err
			}
			fmt.Println("readPkg err=",err)
		}
		err = this.serverProcessMes(&mes)
		if err != nil{
			return err
		}
	}
}