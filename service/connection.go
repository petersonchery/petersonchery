package service

import (
	"FirstProject/controller"
	"FirstProject/model"
)

func ConnectUser(u *model.Users) error {
	return controller.ConnectUser(u)
}
