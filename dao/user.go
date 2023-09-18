package dao

import (
	"FirstProject/config"
	"FirstProject/model"
	"database/sql"
	"errors"
	"fmt"
)

func GetUserById(id int) (model.Users, error) {
	fmt.Println("id:", id)
	var u model.Users
	query := "SELECT id_user, nom_user, email_user, password_user FROM users WHERE id_user = $1 AND isdeleted = false"
	err := config.DB.QueryRow(query, id).Scan(&u.Id_user, &u.Nom_user, &u.Email_user, &u.Password_user)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Users{}, fmt.Errorf("y a erreur")
	}
	if err != nil {
		fmt.Println("y a erreur dans la selection")
		return model.Users{}, nil
	}

	return u, nil
}

func GetUserByEmail(email, pass string) (model.Users, error) {
	var b model.Users
	fmt.Println("l'email:", email, " le pass:", pass)
	query := "SELECT id_user, id_role, nom_user, email_user, password_user FROM users WHERE (email_user= $1 AND password_user = $2) AND isdeleted =false"
	err := config.DB.QueryRow(query, email, pass).Scan(&b.Id_user, &b.Id_role, &b.Nom_user, &b.Email_user, &b.Password_user)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("y a erreur dans dao.GetUserByEmail")
		return model.Users{}, nil

	}
	if err != nil {
		// ecriture dans le journal
		return model.Users{}, err
	}
	fmt.Println("dans dao", b)
	return b, nil
}

func InsertUser(u *model.Users) error {
	query := "INSERT INTO users(nom_user, email_user, password_user, id_role, id_tenant, id_session) VALUES($1, $2, $3, $4, $5, $6) RETURNING id_user"
	err := config.DB.QueryRow(query, u.Nom_user, u.Email_user, u.Password_user, u.Id_role, u.Id_tenant, u.Id_session).Scan(&u.Id_user)
	if err != nil {
		fmt.Println("y a erreur dans dao.InsertUser")
		return fmt.Errorf("y a erreur dans dao.InsertUser")
	}
	return nil
}
func InsertRole(r *model.Role) error {
	query := "INSERT INTO roles (nom_role) VALUES ($1) RETURNING id_role"
	err := config.DB.QueryRow(query, r.Nom_role).Scan(&r.Id_role)
	if err != nil {
		fmt.Println("y a erreur dans dao.InsertRole")
		return fmt.Errorf("y a erreur dans dao.InsertRole")
	}
	return nil

}

func SelectUser(u *model.Users) error {

	query := "SELECT id_user, nom_user, email_user, password_user FROM users WHERE id_user = $1 AND isdeleted=false"
	fmt.Println("ou la", u)
	err := config.DB.QueryRow(query, u.Id_user).Scan(&u.Id_user, &u.Nom_user, &u.Email_user, &u.Password_user)
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
	fmt.Println("enfin")
	return nil

}
func DeleteUser(u *model.Users) error {
	query := "UPDATE users SET isdeleted =true WHERE id_user= $1"
	_, err := config.DB.Exec(query, u.Id_user)
	if err != nil {
		fmt.Println("y a erreur dans la mise a jour")
		return fmt.Errorf("y a erreur dans la mise a jour")
	}
	return fmt.Errorf("Suppression reussie")
}
func UpdateUser(u *model.Users) error {
	query := "UPDATE users SET nom_user= $2, email_user=$4, password_user=$3 WHERE id_user =$1 AND isdeleted =false"
	_, err := config.DB.Exec(query, u.Id_user, u.Nom_user, u.Password_user, u.Email_user)
	if err != nil {
		return fmt.Errorf("y a erreur dans dao.UpdateUser")
	}
	return fmt.Errorf("mise a jour reussie")

}

func HasPermission(idUser, idPermission int) (bool, error) {
	var exists bool
	const statement = `SELECT EXISTS(select id_role_permission
		from roles_permissions rp 
		inner join roles r on r.id_role=rp.id_role 
		inner join users u ON u.id_role=r.id_role 
		where u.id_user=$1 AND rp.id_permission=$2)`

	err := config.DB.QueryRow(statement, idUser, idPermission).Scan(&exists)

	return exists, err
}

func GetUserBySession(token string) (model.Users, error) {
	var b model.Users
	fmt.Println("le token:", token)
	query := "SELECT id_user FROM sessions WHERE token2= $1 "
	err := config.DB.QueryRow(query, token).Scan(&b.Id_user)
	if errors.Is(err, sql.ErrNoRows) {
		return model.Users{}, nil
	}
	if err != nil {
		// ecriture dans le journal
		return model.Users{}, err
	}
	fmt.Println("dans dao.GetUserBySession", b)
	return b, nil

}

func GetUserByNameAndPass(username, password string) (model.Users, error) {
	var b model.Users
	fmt.Println("le nom:", username, "le password:", password)
	query := "SELECT id_user, id_role, nom_user, email_user, password_user FROM users WHERE (nom_user= $1 AND password_user = $2) AND isdeleted =false"
	err := config.DB.QueryRow(query, username, password).Scan(&b.Id_user, &b.Id_role, &b.Nom_user, &b.Email_user, &b.Password_user)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("y a erreur dans dao.GetUserByName")
		return model.Users{}, nil

	}
	if err != nil {
		// ecriture dans le journal
		fmt.Println("y a erreur")
		return model.Users{}, err
	}
	fmt.Println("dans dao", b)
	return b, nil
}
