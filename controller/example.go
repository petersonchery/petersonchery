package controller

import (
	"FirstProject/model"
	"FirstProject/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//"fmt"
//"FirstProject/model"

var Nom = "la"

/*
	func Ajouter_livre(id_livre int, titre_livre string, desc_livre string) model.Livre {
		return model.Livre{
			Id_livre:    id_livre,
			Titre_livre: titre_livre,
			Desc_livre:  desc_livre,
		}

}
*/

func Ajouter_livre() model.Livre {
	var (
		id_livre    int
		titre_livre string
		desc_livre  string
	)
	var clavier = bufio.NewScanner(os.Stdin)
	fmt.Println("veuillez entrer les info conccernant le livre a enregistrer")
	fmt.Println("Entrer l'id du livre")
	clavier.Scan()
	id_livre, err := strconv.Atoi(clavier.Text())
	utils.VerifError(err)

	fmt.Println("Entrer le titre du livre", titre_livre)
	titre_livre, err = bufio.NewReader(os.Stdin).ReadString('\n')
	utils.VerifError(err)

	fmt.Println("Entrer la des du livre", desc_livre)
	desc_livre, err = bufio.NewReader(os.Stdin).ReadString('\n')
	utils.VerifError(err)
	fmt.Println("Votre livre a ete ajoute avec succes")
	tre_livre := ",le titre :"
	de_livre := ", La description :"
	titre_livre = tre_livre + titre_livre
	desc_livre = de_livre + desc_livre

	return model.Livre{
		Id_livre:    id_livre,
		Titre_livre: titre_livre,
		Desc_livre:  desc_livre,
	}
}

func Update_livre() {

}
func Enlever_livre() {

}
func Livre_dispo() {

}
func Afficher_tout() {

}
