package controller

import (
	"FirstProject/dao"
	"FirstProject/model"
	"fmt"
)

func GetRoleById(r *model.Role) error {
	rol, err := dao.GetRoleById(r.Id_role)
	if err != nil {
		fmt.Println("y a erreur dans GetRoleById")
		return nil
	}
	if rol.Nom_role == r.Nom_role {
		return nil
	}
	return nil
}

func CreateRole(role *model.Role) error {

	rol, err := dao.GetRoleById(role.Id_role)
	if err != nil {
		return err
	}
	if rol.Id_role != 0 {
		return fmt.Errorf("Ce livre existe deja")
	}

	err = dao.InsertRole(role)
	return nil
}

func DeleteRoleById(role *model.Role) error {
	rol, err := dao.GetRoleById(role.Id_role)
	if err != nil {
		return err
	}
	if rol.Id_role == role.Id_role {
		return fmt.Errorf("y a pas de role qui correspond a cet id")
	}

	fmt.Println("On va faire la suppression de")
	fmt.Println(role)
	err = dao.DeleteRoleById(role)

	return err
}

func UpdateRoleById(role *model.Role) error {

	rol, err := dao.GetRoleById(role.Id_role)
	if err != nil {
		fmt.Println("y a erreur")
		return nil
	}
	if rol.Id_role == role.Id_role {
		fmt.Println("y a deja un role a cet id")
		return fmt.Errorf("y a deja un role a cet id")
	}
	err = dao.UpdateRoleById(role)
	return err
}

func ListerRole(id_role int) (*model.Role, error) {
	return nil, nil
}
