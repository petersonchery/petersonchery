package service

import (
	"FirstProject/controller"
	"FirstProject/model"
)

func GetRoleById(b *model.Role) error {
	return controller.GetRoleById(b)
}

func CreateRole(b *model.Role) error {
	return controller.CreateRole(b)
}
func DeleteRoleById(r *model.Role) error {
	return controller.DeleteRoleById(r)
}
func UpdateRoleById(r *model.Role) error {
	return controller.UpdateRoleById(r)
}
