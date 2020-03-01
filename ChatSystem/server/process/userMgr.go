package process2

import "fmt"

var (
	userMgr *UserMgr
)


type UserMgr struct {
	onlineUsers map[int] *UserProcess
}


func init(){
	userMgr = &UserMgr{
		onlineUsers:make(map[int]*UserProcess,1024),
	}
}

//添加用户
func (this *UserMgr) AddOnlineUser(up *UserProcess){
	this.onlineUsers[up.UserId]=up
}
//删除用户
func (this *UserMgr) DelOnlineUser(userId int){
	delete(this.onlineUsers,userId)
}
//返回当前所有在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess{
	return  this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int)(up *UserProcess,err error) {
	up , ok := this.onlineUsers[userId]
	if !ok {//说明查找的用户，当前不在线
		err = fmt.Errorf("用户%d 不存在",userId)
		return
	}
	return
}






