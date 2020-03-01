package main

import (
	"ChatSystem/server/model"
	"fmt"
	"net"
	"time"
)




//处理与客户端的通信
func process(conn net.Conn)  {
	//这里需要延时关闭conn
	defer conn.Close()
	//创建一个总控
	processor := &Processsor{
		Conn:conn,
	}
	err := processor.startProcess()
	if err != nil{
		fmt.Println("客户端和服务器通信协程错误err=",err)
	}
}
func main()  {

	//服务器启动是，初始化连接池
	initPool("localhost:6379",16,0,300*time.Second)
	initUserDao()

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

//编写一个函数，完成对UserDao的初始化
func initUserDao(){

	//pool本身是一个全局的变量
	model.MyUserDao =model.NewUserDao(pool)
}