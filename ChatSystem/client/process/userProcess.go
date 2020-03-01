package process

import (
	"ChatSystem/client/utils"
	common "ChatSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess) Register(userId int,
	userPwd string,userName string)(err error){
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.通过conn发送消息给服务
	var mes common.Message
	mes.Type = common.RegisterMesType
	//3.创建一个LoginMes结构体
	var registerMes common.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName =userName
	//4.将loginMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7.
	tf := &utils.Transfer{
		Conn: conn,
	}
	err =tf.WritePkg(data)
	if err !=nil{
		fmt.Println("发送信息出错err =",err)
	}
	mes, err = tf.ReadPkg()
	//将mes的Data部分反序列话为RegisterReMes
	var registerResMes common.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，重新登录")
		os.Exit(0)
	} else  {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return

}

//完成登录
func (this *UserProcess)Login(userId int,passWord string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.dial err=", err)
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
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7.data是要发送的数据
	//7.1先把data的长度发给服务器
	//7.2先获取fata的长度-》转成一个表示一个表示长度的切片
	var pkgLen uint32
	pkgLen = uint32((len(data)))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) failed", err)
		return
	}
	fmt.Printf("%v",string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	//将mes的Data部分反序列话为LoginReMes
	if err != nil {
		fmt.Println("readPkg err", err)
		return
	}
	//将mes的Data怒分反序列化为LoginResMes
	var loginResMes common.LoginReMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = common.UserOnline

		//显示当前用户在线列表
		fmt.Println("当前用户在线列表如下：")
		for _,v := range loginResMes.UserIds{

			//不显示自己
			if v==userId{
				continue
			}
			fmt.Println("用户id:\t",v)
			user := &common.User{
				UserId:     v,
				UserStatus: common.UserOnline,
			}
			onlineUsers[v] = user



		}
		fmt.Printf("\n\n")
		//开启一个协程
		//服务器有数据推送过来，则接收并显示在客户端
		go serverProcessMes(conn)
		//调用显示登录成功的菜单
		for{
			ShowMenu()
		}
	} else  {
		fmt.Println(loginResMes.Error)
	}

	return
}
