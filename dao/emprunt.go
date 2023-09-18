package dao

import (
	"FirstProject/config"
	"FirstProject/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func GetEmpruntByDate(date time.Time) (model.Emprunt, error) {

	var b model.Emprunt
	query := "SELECT id_emprunt, date_emprunt FROM emprunts WHERE date_emprunt = $1 AND isdeleted=false"
	err := config.DB.QueryRow(query, date).Scan(&b.Id_emprunt, &b.Date_emprunt)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Emprunt{}, nil
	}
	if err != nil {
		// ecriture dans le journal
		return model.Emprunt{}, err
	}
	return b, nil
}
func InsertEmprunt(b *model.Emprunt) error {
	insertSQL := "INSERT INTO emprunts(date_emprunt) VALUES($1) RETURNING id_emprunt"
	err := config.DB.QueryRow(insertSQL, b.Date_emprunt).Scan(&b.Id_emprunt)
	if err != nil {
		log.Print(err)
		fmt.Println("on est dans l'insertEmprunt")
		err = fmt.Errorf("Il y a erreur")
		return err
	}
	fmt.Println("donnees entre avec succes")
	return nil
}

func SelectEmprunt(b *model.Emprunt) error {

	query := "SELECT id_emprunt, date_emprunt FROM emprunts WHERE id_emprunt = $1 AND isdeleted = false"
	fmt.Println("ou la", b)
	err := config.DB.QueryRow(query, b.Id_emprunt).Scan(&b.Id_emprunt, &b.Date_emprunt)
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
	//defer rows.Close() // Fermer les lignes de résultat à la fin de la fonction
	fmt.Println("ca passe")
	return nil
}

func GetEmpruntById(id int) (model.Emprunt, error) {
	var e model.Emprunt
	query := "SELECT id_emprunt, date_emprunt FROM emprunts WHERE id_emprunt = $1 AND isdeleted=false"
	err := config.DB.QueryRow(query, id).Scan(&e.Id_emprunt, &e.Date_emprunt)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Emprunt{}, nil
	}
	if err != nil {
		return model.Emprunt{}, fmt.Errorf(" y a erreur dans la selection")
	}
	return e, nil
}

func DeleteEmpruntById(e *model.Emprunt) error {
	fmt.Println("avan delete")
	query := "UPDATE emprunts SET isdeleted=true WHERE id_emprunt = $1"
	fmt.Println("apres delete")
	_, err := config.DB.Exec(query, e.Id_emprunt)

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

func UpdateEmpruntById(e *model.Emprunt) error {
	query := "UPDATE emprunts SET date_emprunt =$2 WHERE id_emprunt = $1 AND isdeleted=false"
	_, err := config.DB.Exec(query, e.Id_emprunt, e.Date_emprunt)
	if err != nil {
		fmt.Println("y a erreur dans la requete")
		return fmt.Errorf("y a erreur dans la requette")
	}
	return fmt.Errorf("mise a jour reussie")
}
