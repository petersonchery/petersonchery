package service

import (
	"FirstProject/controller"
	"FirstProject/model"
)

func CreateUser(u *model.Users) error {
	return controller.CreateUser(u)
}

func GetUserById(u *model.Users) error {
	return controller.GetUserById(u)
}
func DeleteUser(u *model.Users) error {
	return controller.DeleteUser(u)
}
func UpdateUser(u *model.Users) error {
	return controller.UpdateUser(u)
}

func HasPermission(idUser, createUserPermission int) (bool, error) {

	isExist, err := controller.HasPermission(idUser, createUserPermission)
	return isExist, err
}
