package dao

/*

package dao

import (
	"time"
	"personnel/config"
	connect "personnel/connexion"
	"personnel/model"
)

func Login(user model.Users) (model.Users, error) {
	DB := connect.DatabaseInit()
	defer DB.Close()
	var c model.NullUsers

	rechercher_client := `SELECT id,nom,prenom,adresse,email,username,password,
	role,ville,pays,etat,codepostal,tel1,tel2,note,datecreation,datemodification,
	isactivated,usercreate,userupdate,tenantid FROM users WHERE username=$1`
	err := DB.QueryRow(rechercher_client, user.Username).Scan(&c.Id, &c.Nom, &c.Prenom,
		&c.Adresse, &c.Email, &c.Username, &c.Password, &c.Role, &c.Ville, &c.Pays,
		&c.Etat, &c.Codepostal, &c.Tel1, &c.Tel2, &c.Note, &c.Datecreation, &c.Datemodification,
		&c.Isactivated, &c.UserCreate, &c.UserUpdate, &c.TenantId)
	user = ConvertNullUsersToUsers(c)
	config.WriteLog(user.Username)
	return user, err
}
func CreateSession(tenantid int, iduser int, roleid int, token string) string {

	var errorStr string = ""

	conn := connect.DatabaseInit()
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	stmt, err := tx.Prepare("INSERT INTO session (iduser, idrole,token, " +
		"isvalide, lastused,createdate,tenantid) VALUES ($1,$2,$3,true,CURRENT_TIMESTAMP," +
		"CURRENT_TIMESTAMP,$4) ;")
	if err != nil {

		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}
	defer stmt.Close()

	var idsession int
	err = stmt.QueryRow(iduser, roleid, token, tenantid).Scan(&idsession)
	//defer rows.Close()
	if err != nil {

		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	if err = tx.Commit(); err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	return errorStr

}
func FindSession(token string) (model.Session, string) {
	var session model.Session
	var errorStr string = ""

	conn := connect.DatabaseInit()
	defer conn.Close()

	query := `SELECT id, iduser, idrole, token, isvalide, lastused, createdate, tenantid
	FROM session WHERE token=$1`
	rows, err := conn.Query(query, token)
	if err != nil {
		config.WriteLog(err.Error())
		return session, errorStr
	}
	defer rows.Close()
	if err != nil {

		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	if rows.Next() {
		var createat, lastused time.Time
		err := rows.Scan(&session.Id, &session.IdUser, &session.IdRole, &session.Token,
			&session.IsValid, &lastused, &createat, &session.TenantId)

		if err != nil {

			config.WriteLog(err.Error())
			return session, errorStr
		}
		session.LastUsed = config.Convert_date_time_to_string(lastused)
		session.CreateDate = config.Convert_date_time_to_string(createat)

	}

	return session, errorStr
}
func IfHasPermission(roleid, permissionid int) (bool, string) {

	var errorStr string = ""

	conn := connect.DatabaseInit()
	defer conn.Close()

	query := "select* from assrolepermission where(role=$1 and permission=$2)"

	rows, err := conn.Query(query, roleid, permissionid)
	if err != nil {
		config.WriteLog(err.Error())

		return false, errorStr
	}
	defer rows.Close()

	if err != nil {
		config.WriteLog(err.Error())

		errorStr = err.Error()
	}

	if rows.Next() {
		return true, errorStr
	}

	return false, errorStr
}
func DeactivateSession(token string) string {
	var errorStr string = ""

	conn := connect.DatabaseInit()
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		config.WriteLog(err.Error())

		errorStr = "Désolé, une erreur est survenue !"
	}
	query := "UPDATE session SET isvalide ='false'  WHERE (token=$1)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}
	defer stmt.Close()
	_, err = stmt.Exec(token)
	//defer rows.Close()
	if err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	if err = tx.Commit(); err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	return errorStr
}
func GetDate() (time.Time, string) {
	errStr := ""
	var dateNow time.Time
	query := "SELECT CURRENT_TIMESTAMP as datenow;"
	conn := connect.DatabaseInit()
	defer conn.Close()
	rows, err := conn.Query(query)
	if err != nil {
		config.WriteLog(err.Error())
		errStr = err.Error()
	}
	defer rows.Next()
	if rows.Next() {
		err := rows.Scan(&dateNow)

		if err != nil {
			config.WriteLog(err.Error())
			errStr = err.Error()
		}

	}

	return dateNow, errStr
}
func UpdateLastUsedDateToken(token string) string {
	var errorStr string = ""

	conn := connect.DatabaseInit()
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	stmt, err := tx.Prepare("UPDATE session SET lastused = CURRENT_TIMESTAMP" +
		"  WHERE (TOKEN = $1);")
	if err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}
	defer stmt.Close()
	_, err = stmt.Exec(token)
	//defer rows.Close()
	if err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	if err = tx.Commit(); err != nil {
		config.WriteLog(err.Error())
		errorStr = "Désolé, une erreur est survenue !"
	}

	return errorStr
}
*/
