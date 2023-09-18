package dao

import (
	"FirstProject/config"
	"FirstProject/model"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InsertTenant(t *model.Tenant) error {
	//insertion

	insertSQL := `INSERT INTO tenants (nom_tenant) VALUES ($1) RETURNING id_tenant`
	err := config.DB.QueryRow(insertSQL, t.Name_tenant).Scan(&t.Id_Tenant)
	if err != nil {
		log.Print(err)
		fmt.Println("on est la dans l'insertTenant")
		return err
	}
	fmt.Println("Donnees entrees avec succès")
	return nil
}

func SelectTenant(t *model.Tenant) error {

	query := "SELECT id_tenant, nom_tenant FROM tenants WHERE id_tenant = $1 AND isdeleted=false"
	fmt.Println("ou la", t)
	err := config.DB.QueryRow(query, t.Id_Tenant).Scan(&t.Id_Tenant, &t.Name_tenant)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("aucune donnee trouvee")
	}
	if err != nil {
		// ecriture dans le journal
		return err
	}

	fmt.Println("on est la")
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête :", err)
		return err
	}

	fmt.Println("ca passe")
	// for rows.Next() {
	// 	// var column1Value int
	// 	// var column2Value string
	// 	err := rows.Scan()
	// 	if err != nil {
	// 		fmt.Println("Erreur lors de la lecture des résultats :", err)
	// 		return err
	// 	}
	// 	fmt.Printf("Colonne 1 : %d, Colonne 2 : %s\n", b.Id_bibliotheque, b.Nom_bibliotheque)

	// }
	fmt.Println("enfin")
	return nil

}

func GetTenantById(id int) (model.Tenant, error) {
	var t model.Tenant
	query := "SELECT id_tenant, nom_tenant FROM tenants WHERE id_tenant = $1 AND isdeleted=false "
	err := config.DB.QueryRow(query, id).Scan(&t.Id_Tenant, &t.Name_tenant)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Tenant{}, nil
	}
	if err != nil {
		return model.Tenant{}, fmt.Errorf(" y a erreur dans la selection")
	}
	return t, nil
}

func GetTenantByName(name string) (model.Tenant, error) {
	var b model.Tenant
	fmt.Println("le nom:", name)
	query := "SELECT id_tenant, nom_tenant FROM tenants WHERE nom_tenant= $1 AND isdeleted =false"
	err := config.DB.QueryRow(query, name).Scan(&b.Id_Tenant, &b.Name_tenant)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Tenant{}, nil
	}
	if err != nil {
		// ecriture dans le journal
		return model.Tenant{}, err
	}
	fmt.Println("dans dao", b)
	return b, nil
}

func DeleteTenantById(t *model.Tenant) error {
	query := "UPDATE tenants SET isdeleted=true WHERE id_tenant = $1"
	_, err := config.DB.Exec(query, t.Id_Tenant)

	//err := DB.QueryRow(query, b.Id_bibliotheque).Scan(&b.Id_bibliotheque, &b.Nom_bibliotheque)
	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans la requete de suppression")
	}
	return fmt.Errorf("Suppression reussie")

}

func UpdateTenant(t *model.Tenant) error {
	query := "UPDATE tenants SET nom_tenant= $2 WHERE id_tenant =$1 AND isdeleted = false"
	_, err := config.DB.Exec(query, t.Id_Tenant, t.Name_tenant)
	if err != nil {
		return fmt.Errorf("y a erreur dans la requete")
	}
	return fmt.Errorf("mise a jour reussie")

}
