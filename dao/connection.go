package dao

import (
	"FirstProject/config"
	"FirstProject/model"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func ConnectUser(u *model.Session) error {
	insertSql := "INSERT INTO sessions(id_session, nom_session, duree) VALUES ($1) RETURNING id_session "
	err := config.DB.QueryRow(insertSql, u.Id_session, u.Nom_session, u.Duree).Scan(&u.Id_session)
	if err != nil {
		fmt.Println("y a erreur dans l'insertion")
		return fmt.Errorf("y a erreur dans l'insertion")
	}

	return fmt.Errorf("connection reussie, Session ouverte")

}

func UpdateSession(u *model.Session) error {
	query := "UPDATE sessions SET nom_session = $1, duree = $3 WHERE id_session = $2"
	_, err := config.DB.Exec(query, u.Nom_session, u.Id_session, u.Duree)
	if err != nil {
		fmt.Println("y a erreur")
		return fmt.Errorf("y a erreur dans la requete")
	}
	return fmt.Errorf("connection reussie, Session ouverte")

}

func CreateSession(token uuid.UUID, id_user int) error {

	tx, err := config.DB.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer tx.Rollback()

	id := InsertSession(token, tx)
	if id == 0 {
		if err != nil {
			fmt.Println(err)
		}
		return err
	}

	err = UpdateUserSession(id_user, id, tx)
	if err != nil {
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}

	tx.Commit()
	return nil

}

func InsertSession(token uuid.UUID, tx *sql.Tx) int {
	var s = &model.Session{}
	insertSql := "INSERT INTO sessions (token2) VALUES ($1) RETURNING id_session"
	err := tx.QueryRow(insertSql, token).Scan(&s.Id_session)
	if err != nil {
		fmt.Println("y a erreur dans l'insertion pour session")
		return 0
	}
	return s.Id_session

}

func UpdateUserSession(id_user int, id int, tx *sql.Tx) error {
	sql_statement := "UPDATE users SET id_session= $1 WHERE id_user = $2"
	_, err := tx.Exec(sql_statement, id, id_user)
	if err != nil {
		fmt.Println("y a erreur dans l'insertion pour session")
		return nil
	}
	return nil

}
