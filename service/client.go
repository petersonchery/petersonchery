package service

import (
	"FirstProject/controller"
	"FirstProject/model"
)

func CreateClient(c *model.Client) error {
	return controller.CreateClient(c)
}

func GetClientById(c *model.Client) error {
	return controller.GetClientById(c)
}

func DeleteClientById(c *model.Client) error {
	return controller.DeleteClient(c)

}
func UpdateClient(c *model.Client) error {
	return controller.UpdateClient(c)
}
