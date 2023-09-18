package controller

import (
	"FirstProject/dao"
	"FirstProject/model"
	"fmt"
)

//implementation des fonctions qui sont dans l'interface

func CreateClient(b *model.Client) error {

	client, err := dao.GetByNameClient(b.Nom_client)
	if err != nil {
		return err
	}
	if client.Nom_client != "" {
		return fmt.Errorf("Ce Client est existe deja")
	}
	fmt.Println("client nan controller a", b)
	err = dao.InsertClient(b)
	return nil
}

func GetClientById(c *model.Client) error {

	return dao.SelectClient(c)

}

func UpdateClient(c *model.Client) error {
	client, err := dao.GetByNameClient(c.Nom_client)
	if err != nil {
		fmt.Println("y a erreur dans controller")
		return nil
	}
	if client.Nom_client == c.Nom_client {
		fmt.Println("y a deja un client a cet id")
		return fmt.Errorf("y a deja un client a cet id")
	}
	fmt.Println("on va faire la mise a jour de", c)
	err = dao.UpdateClient(c)
	return err
}

func DeleteClient(c *model.Client) error {
	client, err := dao.GetClientById(c.Id_client)
	if err != nil {
		return err
	}
	if client.Id_client != c.Id_client {
		return fmt.Errorf("y a pas de client qui correspond a cet id")
	}

	fmt.Println("On va faire la suppression de")
	fmt.Println(c)
	err = dao.DeleteClientById(c)

	return err
}

func ListerClient(id_bibliotheque int) (*model.Client, error) {
	return nil, nil
}
