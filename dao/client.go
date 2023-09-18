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

func GetByNameClient(name string) (model.Client, error) {
	var b model.Client
	query := "SELECT id_client, nom_client FROM clients WHERE nom_client = $1 AND isdeleted = false"
	err := config.DB.QueryRow(query, name).Scan(&b.Id_client, &b.Nom_client)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Client{}, nil
	}
	if err != nil {
		// ecriture dans le journal
		return model.Client{}, err
	}
	return b, nil
}

func InsertClient(b *model.Client) error {
	//insertion

	insertSQL := `INSERT INTO clients(nom_client) VALUES ($1) RETURNING id_client`
	err := config.DB.QueryRow(insertSQL, b.Nom_client).Scan(&b.Id_client)
	if err != nil {
		log.Print(err)
		fmt.Println("on est la")
		return err
	}
	fmt.Println("Donnees entrees avec succès")
	return nil
}

func SelectClient(b *model.Client) error {

	query := "SELECT id_client, nom_client FROM clients WHERE id_client = $1 AND isdeleted = false"
	fmt.Println("ou la", b)
	err := config.DB.QueryRow(query, b.Id_client).Scan(&b.Id_client, &b.Nom_client)
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

func GetClientById(id int) (model.Client, error) {
	var c model.Client
	query := "SELECT id_client, nom_client FROM clients WHERE id_client = $1 AND isdeleted = false "
	err := config.DB.QueryRow(query, id).Scan(&c.Id_client, &c.Nom_client)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Client{}, nil
	}
	if err != nil {
		return model.Client{}, fmt.Errorf(" y a erreur dans la selection")
	}
	return c, nil
}

func DeleteClientById(c *model.Client) error {
	//var b model.Bibliothèque
	fmt.Println("avan delete")
	query := "UPDATE clients SET isdeleted =true WHERE id_client = $1"
	fmt.Println("apres delete")
	_, err := config.DB.Exec(query, c.Id_client)

	//err := DB.QueryRow(query, b.Id_bibliotheque).Scan(&b.Id_bibliotheque, &b.Nom_bibliotheque)
	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans la requete de suppression")
	}
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("Aucune donnee trouvee")
	}
	return fmt.Errorf("Suppression reussie")

}
func UpdateClient(c *model.Client) error {

	query := "UPDATE clients SET nom_client =$2, isdeleted=false WHERE id_client = $1  "
	_, err := config.DB.Exec(query, c.Id_client, c.Nom_client)

	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans l'update")
	}
	return fmt.Errorf("Mise a jour reussie")
}
