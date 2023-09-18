package controller

import (
	"FirstProject/dao"
	"FirstProject/model"
	"fmt"
)

//implementation des fonctions qui sont dans l'interface

func CreateEmprunt(b *model.Emprunt) error {
	//var emp model.Emprunt
	emp, err := dao.GetEmpruntByDate(b.Date_emprunt)
	if err != nil {
		fmt.Println(" y a erreur")
	}

	if !emp.Date_emprunt.IsZero() {
		return fmt.Errorf("Cet emprunt existe deja")

	}
	return dao.InsertEmprunt(b)

}

func GetEmpruntById(b *model.Emprunt) error {

	return dao.SelectEmprunt(b)
}

func DeleteEmprunt(e *model.Emprunt) error {
	emp, err := dao.GetEmpruntById(e.Id_emprunt)
	if err != nil {
		return err
	}
	if emp.Id_emprunt != e.Id_emprunt {
		return fmt.Errorf("y a pas d'emprunt qui correspond a cet id")
	}

	fmt.Println("On va faire la suppression de")
	fmt.Println(e)
	err = dao.DeleteEmpruntById(e)

	return err
}

func UpdateEmpruntById(e *model.Emprunt) error {

	emp, err := dao.GetEmpruntByDate(e.Date_emprunt)
	if err != nil {
		fmt.Println("y a erreur ")
		return fmt.Errorf("y a erreur")
	}
	if emp.Date_emprunt == e.Date_emprunt {
		fmt.Println("y a deja une date a cet Id")
		return fmt.Errorf("y a pas deja une date a cet id")
	}
	err = dao.UpdateEmpruntById(e)
	return err
}

func ListerEmprunt(id_bibliotheque int) (*model.Emprunt, error) {
	return nil, nil
}
