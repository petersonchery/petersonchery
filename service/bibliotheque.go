package service

import (
	"FirstProject/controller"
	"FirstProject/model"
)

func GetBibliothequeById(b *model.Bibliotheque) error {
	return controller.GetBibliothequeById(b)
}

func CreateBibliotheque(b *model.Bibliotheque) error {
	return controller.CreateBibliotheque(b)
}
func DeleteBibliothequeById(b *model.Bibliotheque) error {
	return controller.DeleteBibliotheque(b)
}

func UpdateBibliotheque(b *model.Bibliotheque) error {
	return controller.UpdateBibliotheque(b)
}
func ListerBibliotheque(b *model.Bibliotheque) ([]model.Bibliotheque, error) {
	return controller.ListerBibliotheque(b)
}
