package controller

import (
	"FirstProject/dao"
	"FirstProject/model"
	"fmt"
)

func CreateUser(u *model.Users) error {
	user, err := dao.GetUserByEmail(u.Email_user, u.Password_user)
	fmt.Println("user dans controller.GetUserByName", user.Nom_user, u.Nom_user)
	if err != nil {
		fmt.Println("y a erreur dans getUserByName")
		return fmt.Errorf("y a erreur dans getUserByName")
	}
	if user.Email_user != "" {
		fmt.Println("Ce user existe deja")
		return fmt.Errorf("ce user existe deja")
	}
	if user.Password_user != "" {
		fmt.Println("Ce user existe deja")
		return fmt.Errorf("ce user existe deja")

	}

	err = dao.InsertUser(u)
	//err = dao.InsertRole(r)
	return err
}

func GetUserById(u *model.Users) error {
	return dao.SelectUser(u)

}
func DeleteUser(u *model.Users) error {
	user, err := dao.GetUserById(u.Id_user)
	if err != nil {
		fmt.Println("y a erreur dans la selection de l'Id")
		return fmt.Errorf("y a erreur dans la selection de l'Id")
	}
	if user.Id_user != user.Id_user {
		fmt.Println("y a aucun user qui correspond a cet id")
		return fmt.Errorf("y a auncun user qui correspond a cet ide")

	}
	err = dao.DeleteUser(u)
	return err
}
func UpdateUser(u *model.Users) error {
	user, err := dao.GetUserById(u.Id_user)
	fmt.Println("l'id usee", user.Id_user, u.Id_user)
	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans l'update")
	}
	if user.Id_user == 0 {
		fmt.Println("cet user n'existe pas")
		return fmt.Errorf("cet user n'existe pas")
	}
	/*
		if user.Email_user != u.Email_user {
			fmt.Println("y a pas de user a cet email ")
			return fmt.Errorf("y a pas de user a cet email")
		}
		if user.Password_user != u.Password_user {
			fmt.Println("y a pas de user a cet email ")
			return fmt.Errorf("y a pas de user a cet email")
		}*/
	err = dao.UpdateUser(u)
	return err
}

func HasPermission(idUser, createUserPermission int) (bool, error) {
	exist, err := dao.HasPermission(idUser, createUserPermission)
	if err != nil {
		fmt.Println("y a erreur dans la permission au niveau du controller")
		return false, fmt.Errorf("y a erreur dans la permission au niveau du controller")
	}
	return exist, nil
}
