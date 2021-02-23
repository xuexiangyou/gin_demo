package models

import orm "github.com/xuexiangyou/gin_demo/database"

type User struct {
	Id int64 `json:"id" form:"id" gorm:"primary_key;auto_increment;"`
	GroupId int64 `json:"groupId" form:"groupId" gorm:"type:bigint(20);"`
	Name string `json:"name" form:"name" gorm:"type:varchar(36);"`
	Email string `json:"email" form:"email" gorm:"type:varchar(36);"`
	Password string `json:"password" form:"password" gorm:"type:varchar(36);"`
	Token string `json:"token" form:"token" gorm:"type:varchar(36);"`
}

func NewUser() *User {
	return new(User)
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetOne() (User, error) {
	var user User
	table := orm.Eloquent.Table(u.TableName())
	err := table.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}