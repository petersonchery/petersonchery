package service

import (
	"FirstProject/controller"
	"FirstProject/model"
)

func CreateEmprunt(b *model.Emprunt) error {
	return controller.CreateEmprunt(b)
}

func GetEmpruntById(b *model.Emprunt) error {
	return controller.GetEmpruntById(b)
}
func DeleteEmpruntById(e *model.Emprunt) error {
	return controller.DeleteEmprunt(e)
}
func UpdateEmpruntById(e *model.Emprunt) error {
	return controller.UpdateEmpruntById(e)
}
