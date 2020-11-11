package users

type Login struct {
	Base
	Pass string `form:"pass" json:"pass" binding:"required,min=6,max=20"` //  密码为 必填，长度>=6
}
