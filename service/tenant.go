package service

import (
	"FirstProject/controller"
	"FirstProject/model"
)

func GetTenantById(t *model.Tenant) error {
	return controller.GetTenantById(t)
}

func CreateTenant(t *model.Tenant) error {
	return controller.CreateTenant(t)
}
func DeleteTenantById(t *model.Tenant) error {
	return controller.DeleteTenant(t)
}
func UpdateTenantById(t *model.Tenant) error {
	return controller.UpdateTenant(t)
}
