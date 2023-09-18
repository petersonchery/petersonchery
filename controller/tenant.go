package controller

import (
	"FirstProject/dao"
	"FirstProject/model"
	"fmt"
)

//implementation des fonctions qui sont dans l'interface

func CreateTenant(tenant *model.Tenant) error {

	ten, err := dao.GetTenantByName(tenant.Name_tenant)
	if err != nil {
		return err
	}
	if ten.Name_tenant != "" {
		return fmt.Errorf("Ce tenant existe deja")
	}

	fmt.Println("on va inserer ce tenant")
	err = dao.InsertTenant(tenant)
	return nil
}

func GetTenantById(t *model.Tenant) error {

	return dao.SelectTenant(t)
}

func DeleteTenant(t *model.Tenant) error {
	ten, err := dao.GetTenantById(t.Id_Tenant)
	if err != nil {
		return err
	}
	if ten.Id_Tenant != t.Id_Tenant {
		return fmt.Errorf("y a pas de livre qui correspond a cet id")
	}

	fmt.Println("On va faire la suppression de")
	fmt.Println(t)
	err = dao.DeleteTenantById(t)

	return err
}

func UpdateTenant(t *model.Tenant) error {
	ten, err := dao.GetTenantByName(t.Name_tenant)
	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans l'update")
	}
	if ten.Name_tenant == t.Name_tenant {
		fmt.Println("y a deja un tenant a cet id")
		return fmt.Errorf("y a deja un tenant a cet id")
	}
	err = dao.UpdateTenant(t)
	return err
}
func ListerTenant(id_bibliotheque int) (*model.Livre, error) {
	return nil, nil
}
