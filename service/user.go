// 機能としてはコントローラから振られたdb操作を引き受け、結果を返す

package service

import (
	"go_crud/model"
)

type UserService struct{}

func (UserService) SetUser(user *model.User) error {
	db := SqlConnect()
	db.Create(user)
	defer db.Close()

	return nil
}

func (UserService) GetUserList() []model.User {
	db := SqlConnect()
	users := make([]model.User, 0)
	db.Order("created_at asc").Find(&users)
	defer db.Close()

	return users
}

// 動いてるけど、よく分からん
func (UserService) UpdateUser(id int, newUser *model.User, data *model.User) error {
	db := SqlConnect()
	db.Where("ID = ?", id).First(newUser).Updates(data)
	defer db.Close()

	return nil
}

func (UserService) DeleteUser(id int) error {
	db := SqlConnect()
	var user model.User
	db.First(&user, id)
	db.Delete(&user)
	defer db.Close()

	return nil
}
