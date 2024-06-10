package biz

import "github.com/omalloc/contrib/kratos/orm"

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"column:username;type:varchar(64);comment:用户名"`
	Password string `json:"password" gorm:"column:password;type:varchar(64);comment:密码"`
	Salt     string `json:"salt" gorm:"column:salt;type:varchar(64);comment:盐"`
	Email    string `json:"email" gorm:"column:email;type:varchar(64);comment:邮箱"`
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(64);comment:昵称"`
	Bio      string `json:"bio" gorm:"column:bio;type:varchar(255);comment:个人简介"`
	AvatarID int64  `json:"avatar_id" gorm:"column:avatar_id;comment:头像"`

	orm.DBModel
}

func (User) TableName() string {
	return "users"
}
