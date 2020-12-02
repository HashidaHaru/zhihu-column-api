package service

import (
	"fmt"
	"zhihu-column-api/db"
	"zhihu-column-api/model"
)

// FindUserByUsername 通过用户名找到用户
func FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.DB.Where("username=?", username).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("db:%w", err)
	}
	return &user, nil
}

// UserAdd 新增用户
func UserAdd(m *model.User) error {
	return db.DB.Create(m).Error
}

// UserDetail 用户详情
func UserDetail(id uint) (*model.User, error) {
	user := model.User{}
	user.ID = id
	err := db.DB.First(&user).Error
	return &user, err
}

// UserModify 修改用户详情
func UserModify(id uint, m model.User) error {
	return db.DB.Model(&model.User{}).Where("id = ?", id).Updates(m).Error
}
