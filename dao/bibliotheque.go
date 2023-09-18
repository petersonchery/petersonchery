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

func Insert(b *model.Bibliotheque) error {
	//insertion
	fmt.Println(b)
	insertSQL := `INSERT INTO bibliotheques(nom_bibliotheque) VALUES ($1) RETURNING id_bibliotheque`
	err := config.DB.QueryRow(insertSQL, b.Nom_bibliotheque).Scan(&b.Id_bibliotheque)
	if err != nil {
		log.Print(err)
		fmt.Println("on est la")
		return err
	}
	fmt.Println("Donnees entrees avec succès")
	return nil
}

func GetByName(name string) (model.Bibliotheque, error) {
	var b model.Bibliotheque
	fmt.Println("le nom:", name)
	query := "SELECT id_bibliotheque, nom_bibliotheque FROM bibliotheques WHERE nom_bibliotheque= $1 AND isdeleted=false"
	err := config.DB.QueryRow(query, name).Scan(&b.Id_bibliotheque, &b.Nom_bibliotheque)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Bibliotheque{}, nil
	}
	if err != nil {
		// ecriture dans le journal
		return model.Bibliotheque{}, err
	}
	fmt.Println("dans dao", b)
	return b, nil
}

func Select(b *model.Bibliotheque) error {

	query := "SELECT id_bibliotheque, nom_bibliotheque FROM bibliotheques WHERE id_bibliotheque = $1 AND isdeleted=false"
	fmt.Println("ou la", b)
	err := config.DB.QueryRow(query, b.Id_bibliotheque).Scan(&b.Id_bibliotheque, &b.Nom_bibliotheque)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("aucune donnee trouvee")
	}
	fmt.Println("on est la")
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête :", err)
		return err
	}
	//defer rows.Close() // Fermer les lignes de résultat à la fin de la fonction
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

func Delete(b *model.Bibliotheque) error {
	//var b model.Bibliotheque
	fmt.Println("avan delete")
	query := "UPDATE bibliotheques SET isdeleted=true WHERE id_bibliotheque = $1"
	fmt.Println("apres delete")
	_, err := config.DB.Exec(query, b.Id_bibliotheque)

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

func Update(b *model.Bibliotheque) error {

	query := "UPDATE bibliotheques SET isdeleted=false, nom_bibliotheque =$2 WHERE id_bibliotheque = $1"
	_, err := config.DB.Exec(query, b.Id_bibliotheque, b.Nom_bibliotheque)

	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans l'update")
	}
	return fmt.Errorf("Mise a jour reussie")
}

func GetById(id int) (model.Bibliotheque, error) {
	var b model.Bibliotheque
	query := "SELECT id_bibliotheque, nom_bibliotheque FROM bibliotheques WHERE id_bibliotheque = $1 AND isdeleted=FALSE "
	err := config.DB.QueryRow(query, id).Scan(&b.Id_bibliotheque, &b.Nom_bibliotheque)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Bibliotheque{}, nil
	}
	if err != nil {
		return model.Bibliotheque{}, fmt.Errorf(" y a erreur dans la selection")
	}
	return b, nil

}

func SelectAll(b *model.Bibliotheque) ([]model.Bibliotheque, error) {
	var biblios []model.Bibliotheque

	query := "SELECT * FROM bibliotheques WHERE isdeleted=false"
	rows, err := config.DB.Query(query)
	if err != nil {
		fmt.Println("don't ok")
		return nil, nil
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&b.Id_bibliotheque, &b.Nom_bibliotheque, &b.Id_tenant, &b.Is_deleted); err != nil {

			log.Fatal("error")

		}

	}

	biblios = append(biblios, *b)
	return biblios, err
}
