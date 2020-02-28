package main

import (
	common "ChatSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//完成登录
func login(userId int,passWord string) (err error) {
	//1.连接到服务器
	conn,err :=net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("net.dial err=",err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.通过conn发送消息给服务
	var mes common.Message
	mes.Type = common.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = passWord

	//4.将loginMes序列化
	data,err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes序列化
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}
	//7.data是要发送的数据
	//7.1先把data的长度发给服务器
	//7.2先获取fata的长度-》转成一个表示一个表示长度的切片
	var pkgLen uint32
	pkgLen = uint32((len(data)))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4],pkgLen)
	//发送长度
	n,err := conn.Write(bytes[:4])
	if n != 4||err !=nil{
		fmt.Println("conn.Write(bytes) failed",err)
		return
	}
	//发送消息本身
	_, err = conn.Write(data)
	if err !=nil{
		fmt.Println("conn.Write(data) fail",err)
		return
	}

	mes,err =readPkg(conn)
	//将mes的Data部分反序列话为LoginReMes
	if err !=nil {
		fmt.Println("readPkg err",err)
		return
	}
	//将mes的Data怒分反序列化为LoginResMes
	var loginResMes common.LoginReMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code==200{
		fmt.Println("登录成功")
	}else if loginResMes.Code ==500{
		fmt.Println(loginResMes.Error)
	}

	return
}