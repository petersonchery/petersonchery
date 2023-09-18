package service

import (
	"FirstProject/controller"
	"FirstProject/model"
)

func GetLivreById(b *model.Livre) error {
	return controller.GetLivreById(b)
}

func CreateLivre(b *model.Livre) error {
	return controller.CreateLivre(b)
}
func DeleteLivreById(l *model.Livre) error {
	return controller.DeleteLivre(l)
}
func UpdateLivreById(l *model.Livre) error {
	return controller.UpdateLivre(l)
}
