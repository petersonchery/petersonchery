package controller

import (
	"FirstProject/dao"
	"FirstProject/model"
	"fmt"
)

//implementation des fonctions qui sont dans l'interface

func CreateLivre(livre *model.Livre) error {

	liv, err := dao.GetByTitreLivre(livre.Titre_livre)
	if err != nil {
		return err
	}
	if liv.Titre_livre != "" {
		return fmt.Errorf("Ce livre existe deja")
	}

	err = dao.InsertLivre(livre)
	return nil
}

func GetLivreById(c *model.Livre) error {

	return dao.SelectLivre(c)
}

func DeleteLivre(l *model.Livre) error {
	liv, err := dao.GetLivreByName(l.Titre_livre)
	if err != nil {
		return err
	}
	if liv.Titre_livre == l.Titre_livre {
		return fmt.Errorf("y a pas de livre qui correspond a cet id")
	}

	fmt.Println("On va faire la suppression de")
	fmt.Println(l)
	err = dao.DeleteLivreById(l)

	return err
}

func UpdateLivre(l *model.Livre) error {

	liv, err := dao.GetLivreByName(l.Titre_livre)
	if err != nil {
		fmt.Println("y a erreur")
		return nil
	}
	if liv.Titre_livre == l.Titre_livre {
		fmt.Println("y a deja un livre a cet id")
		return fmt.Errorf("y a deja un livre a cet id")
	}
	err = dao.UpdateLivre(l)
	return err
}

func ListerLivre(id_bibliotheque int) (*model.Livre, error) {
	return nil, nil
}
