package dao

import (
	"FirstProject/config"
	"FirstProject/model"
	"fmt"
	"log"
)

func GetRoleById(id int) (model.Role, error) {
	return model.Role{}, nil
}

func CreateRole(r *model.Role) error {
	//insertion

	insertSQL := `INSERT INTO roles (nom_role) VALUES ($1) RETURNING id_role`
	err := config.DB.QueryRow(insertSQL, r.Nom_role).Scan(&r.Id_role)
	if err != nil {
		log.Print(err)
		fmt.Println("on est la dans l'insertLivre")
		return err
	}
	fmt.Println("Donnees entrees avec succès")
	return nil
}

func DeleteRoleById(r *model.Role) error {
	query := "UPDATE roles SET isdeleted=true WHERE id_role = $1"
	_, err := config.DB.Exec(query, r.Id_role)
	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans la requete de suppression")
	}
	return fmt.Errorf("Suppression reussie")

}

func UpdateRoleById(r *model.Role) error {
	query := "UPDATE roles SET nom_role= $2 WHERE id_role =$1 AND isdeleted = false"
	_, err := config.DB.Exec(query, r.Id_role, r.Nom_role)
	if err != nil {
		return fmt.Errorf("y a erreur dans la requete")
	}
	return fmt.Errorf("mise a jour reussie")

}
