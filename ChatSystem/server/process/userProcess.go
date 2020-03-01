package process2

import (
	common "ChatSystem/common/message"
	"ChatSystem/server/model"
	"ChatSystem/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {

	Conn net.Conn
	UserId int
}

//通知所有在线的用户的方法
//userId 通知其它的在线用户
func  (this *UserProcess)NotifyOtherOnlineUser(userId int){
	//遍历onlineUsers,然后一个个的发送
	for id,up := range  userMgr.onlineUsers{
		if id == userId{
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int){
	var mes common.Message
	mes.Type = common.NotifyUserStatusMesType

	var notifyUserStatusMes common.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = common.UserOnline


	data ,err := json.Marshal(notifyUserStatusMes)
	if err !=nil{
		fmt.Println("json.Marshal err =",err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值geimes.Data
	mes.Data =string(data)

	//对mes再次序列化，准备发送
	data ,err = json.Marshal(mes)
	if err !=nil{
		fmt.Println("json.Marshal err =",err)
		return
	}

	//发送，创建Transfer实例，发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("NotifyMeOnline err= ",err)
		return
	}
}

func (this *UserProcess) ServerProcessLogin(mes *common.Message)(err error){
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

	//到redis数据库完成一个验证
	user,err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	fmt.Printf("%v登录成功\n",user)
	if err !=nil{

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code=500
			loginResMes.Error = err.Error()

		}else if err==model.ERROR_USER_PWD{
			loginResMes.Code=403
			loginResMes.Error = err.Error()
		}else {
			loginResMes.Code=505
			fmt.Println("服务器内部错误")
		}

	}else {
		loginResMes.Code=200

		//将登录成功用户userId赋值给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其它在线用户上线
		this.NotifyOtherOnlineUser(loginMes.UserId)
		//将当前在线用户id放入loginResMes.UsersID
		for id, _:=range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds,id)
		}
		fmt.Println(user,"登录成功")

	}


	fmt.Println(loginResMes.Code)

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
	tf:=&utils.Transfer{
		Conn: this.Conn,
	}
	err =tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *common.Message) (err error) {

	var registerMes common.RegisterMes
	err =json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}

	//先声明一个resMes
	var resMes common.Message
	resMes.Type = common.RegisterResMesType

	//再声明一个registerResMes
	var registerResMes common.LoginReMes
	//到redis数据库完成一个验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	}else{
		registerResMes.Code = 200
	}
	data,err := json.Marshal(registerResMes)
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
	tf:=&utils.Transfer{
		Conn: this.Conn,
	}
	err =tf.WritePkg(data)
	return
}