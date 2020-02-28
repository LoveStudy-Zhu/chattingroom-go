package main

import (
	common "ChatSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func readPkg(conn net.Conn)(mes common.Message,err error) {
	//这里将读取的数据包直接封装成一个函数readPkg(),
	buf:= make([]byte,8096)

	//conn.Read在coon没有被关闭的情况下，才会阻塞
	_,err = conn.Read(buf[:4])
	if err != nil{
		fmt.Println("conn.Read err=",err)
		return
	}
	//根据buf[0:4]转成uint32类型
	var pkgLen uint32
	pkgLen=binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen读取消息内容
	n,err :=conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil{
		fmt.Printf("conn.Read fail err=%v\n",err)
		return
	}

	//把pkgLen反序列化为Message
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil{
		fmt.Println("json.Unmarsha err=",err)
		return
	}
	return

}

func writePkg(conn net.Conn,data []byte) (err error){
	//先发送一个长度给对方，
	var pkgLen uint32
	pkgLen = uint32((len(data)))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4],pkgLen)
	//发送长度
	n,err := conn.Write(bytes[:4])
	if n != 4||err !=nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	//发送data本身
	n,err = conn.Write(data)
	if n != int(pkgLen)||err !=nil{
		fmt.Printf("conn.Write(bytes) fail=%v,n=%v,pkgLen=%v\n",err,n,pkgLen)
		return
	}
	return
}

func serverProcessLogin(conn net.Conn,mes *common.Message)(err error){
	var loginMes common.LoginMes
	err =json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}


	//先声明一个resMes
	var resMes common.Message
	resMes.Type = common.LoginResMesType

	//再声明一个LoginResMes
	var loginResMes common.LoginReMes

	if loginMes.UserId ==100 && loginMes.UserPwd =="123456"{
		loginResMes.Code = 200
	}else{
		loginResMes.Code =500
		loginResMes.Error = "该用户不存在，请注册"
	}
	//将loginResMes序列化
	data,err := json.Marshal(loginResMes)
	if err !=nil{
		fmt.Println("json.Marshal fail",err)
		return
	}

	resMes.Data =string(data)
	//将resMes序列化,准备发送
	data,err = json.Marshal(resMes)
	if err !=nil{
		fmt.Println("json.Marshal fail",err)
		return
	}
	err =writePkg(conn,data)
	return
}

//根据客户端发送消息的种类决定调用哪个函数来处理
func serverProcessMes(conn net.Conn,mes *common.Message)(err error){
	switch mes.Type {
		case common.LoginMesType:
			//处理登录
			err=serverProcessLogin(conn,mes)
		case common.RegisterMesType:
			//处理注册
		default:
			fmt.Println("消息类型不存在，无法处理")

	}
	return
}


func process(conn net.Conn)  {
	//这里需要延时关闭conn
	defer conn.Close()
	//循环读取客户端发送的信息
	for{
		mes,err := readPkg(conn)
		if err !=nil{
			if err ==io.EOF{
				fmt.Println("客户端退出")
				return
			}
			fmt.Println("readPkg err=",err)
		}
		err = serverProcessMes(conn,&mes)
		if err != nil{
			return
		}
	}
}
func main()  {
	fmt.Println("服务器在8889端口监听")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil{
		fmt.Println("net.listen err=",err)
		return
	}
	for {
		fmt.Println("等待客户端连接服务器")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=",err)
		}
		go process(conn)

	}
}
