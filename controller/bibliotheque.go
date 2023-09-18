package controller

import (
	"FirstProject/dao"
	"FirstProject/model"
	"bufio"
	"fmt"
	"os"
)

// implementation des fonctions qui sont dans l'interface
var clavier = bufio.NewScanner(os.Stdin)

func CreateBibliotheque(b *model.Bibliotheque) error {

	biblio, err := dao.GetByName(b.Nom_bibliotheque)
	fmt.Println(biblio)
	if err != nil {
		return err
	}
	if biblio.Nom_bibliotheque != "" {
		return fmt.Errorf("bibliotheque deja existante")
	}

	fmt.Println(b.Nom_bibliotheque)
	fmt.Println("dans controller")
	err = dao.Insert(b)
	return nil
}

func GetBibliothequeById(b *model.Bibliotheque) error {

	// exécution de requête
	fmt.Println("avant", b)
	return dao.Select(b)

}

func DeleteBibliotheque(b *model.Bibliotheque) error {
	biblio, err := dao.GetById(b.Id_bibliotheque)
	if err != nil {
		return err
	}
	if biblio.Id_bibliotheque != b.Id_bibliotheque {
		return fmt.Errorf("y a pas de Bibliotheque qui correspond a cet id")
	}

	fmt.Println("On va faire la suppression de", b, "dans controller")
	fmt.Println(b)
	err = dao.Delete(b)

	return err
}

func ListerBibliotheque(b *model.Bibliotheque) ([]model.Bibliotheque, error) {
	//var bi =[]model.Bibliotheque{}
	bi, err := dao.SelectAll(b)
	if err != nil {
		return nil, fmt.Errorf("y a erreur")
	}
	return bi, nil
}

func UpdateBibliotheque(b *model.Bibliotheque) error {

	biblio, err := dao.GetById(b.Id_bibliotheque)
	if err != nil {
		fmt.Println("y a erreur dans controller")
		return nil
	}
	if biblio.Nom_bibliotheque != "" {
		fmt.Println("y a deja une bibliotheque qui existe a ce nom, ou l'id ne correspond pas")
		return fmt.Errorf("y a deja une bibliotheque qui existe a ce nom, ou l'id ne correspond pas")
	}
	fmt.Println("on va faire la mise a jour de", b)
	err = dao.Update(b)
	return err

}

/*
func UpdateBibliotheque(id int) (*model.Bibliotheque, error) {

	reader := bufio.NewReader(os.Stdin)
	var biblio = model.Bibliotheque{}
	fmt.Print("Entrez l'id du bibliotheque a modifier: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		os.Exit(1)
	}

	val, _ := strconv.Atoi(input)
	fmt.Println("Vous avez entre :", input)
	biblio.Id_bibliotheque = val
	// exécution de requête
	fmt.Println("avant")
	dao.Select(&biblio)

	fmt.Print("Entrez le champs du bibliotheque que vous voulez modifier: ")
	input, err = reader.ReadString('\n')

	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		os.Exit(1)
	}

	return nil, nil
}
*/
