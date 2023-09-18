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

func GetByTitreLivre(titre string) (model.Livre, error) {
	var b model.Livre

	query := "SELECT id_livre, titre_livre, auteur_livre, desc_livre FROM livres WHERE titre_livre = $1 AND isdeleted=false"
	err := config.DB.QueryRow(query, titre).Scan(&b.Id_livre, &b.Titre_livre, &b.Auteur, &b.Desc_livre)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Livre{}, nil
	}
	if err != nil {
		// ecriture dans le journal
		return model.Livre{}, err
	}
	return b, nil
}

func InsertLivre(b *model.Livre) error {
	//insertion

	insertSQL := `INSERT INTO livres (titre_livre, auteur_livre, desc_livre) VALUES ($1, $2, $3) RETURNING id_livre`
	err := config.DB.QueryRow(insertSQL, b.Titre_livre, b.Auteur, b.Desc_livre).Scan(&b.Id_livre)
	if err != nil {
		log.Print(err)
		fmt.Println("on est la dans l'insertLivre")
		return err
	}
	fmt.Println("Donnees entrees avec succès")
	return nil
}

func SelectLivre(b *model.Livre) error {

	query := "SELECT id_livre, titre_livre, auteur_livre, desc_livre FROM livres WHERE id_livre = $1 AND isdeleted=false"
	fmt.Println("ou la", b)
	err := config.DB.QueryRow(query, b.Id_livre).Scan(&b.Id_livre, &b.Titre_livre, &b.Auteur, &b.Desc_livre)
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

func GetLivreById(id int) (model.Livre, error) {
	var l model.Livre
	query := "SELECT id_livre, titre_livre, auteur_livre, desc_livre FROM livres WHERE id_livre = $1 AND isdeleted=false"
	err := config.DB.QueryRow(query, id).Scan(&l.Id_livre, &l.Titre_livre, &l.Auteur, &l.Desc_livre)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Livre{}, nil
	}
	if err != nil {
		return model.Livre{}, fmt.Errorf(" y a erreur dans la selection")
	}
	return l, nil
}

func GetLivreByName(name string) (model.Livre, error) {
	var b model.Livre
	fmt.Println("le nom:", name)
	query := "SELECT id_livre, titre_livre, auteur_livre, desc_livre FROM livres WHERE titre_livre= $1 AND isdeleted =false"
	err := config.DB.QueryRow(query, name).Scan(&b.Id_livre, &b.Titre_livre, &b.Auteur, &b.Desc_livre)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Livre{}, nil
	}
	if err != nil {
		// ecriture dans le journal
		return model.Livre{}, err
	}
	fmt.Println("dans dao", b)
	return b, nil
}

func DeleteLivreById(l *model.Livre) error {
	query := "UPDATE livres SET isdeleted=true WHERE id_livre = $1"
	_, err := config.DB.Exec(query, l.Id_livre)

	//err := DB.QueryRow(query, b.Id_bibliotheque).Scan(&b.Id_bibliotheque, &b.Nom_bibliotheque)
	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans la requete de suppression")
	}
	return fmt.Errorf("Suppression reussie")

}

func UpdateLivre(l *model.Livre) error {

	query := "UPDATE livres SET titre_livre =$2 WHERE id_livre =$1 AND isdeleted= false"
	_, err := config.DB.Exec(query, l.Id_livre, l.Titre_livre)
	if err != nil {
		fmt.Println("y a erreur dans la requete")
		return fmt.Errorf("y a erreur dans la requete")
	}
	return fmt.Errorf("mise a jour reussie")

}
