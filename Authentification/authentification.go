package authentification

import (
	"FirstProject/dao"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var sessions = make(map[uuid.UUID]int) // token to user ID mapping

func GetUser(token string) int {
	user, err := dao.GetUserBySession(token)
	if err != nil {
		fmt.Println(" y a erreur dans getUserBySession")
		return 0
	}
	sessions[uuid.UUID([]byte(token))] = user.Id_user

	return sessions[uuid.UUID([]byte(token))]

}

func Login(username, password string) error {
	// Fetch user from database based on username
	user, err := dao.GetUserByNameAndPass(username, password)
	if err != nil {
		return err
	}

	// Verify password
	if user.Password_user != password { //bcrypt
		return errors.New("invalid data")
	}
	if user.Nom_user != username {
		return errors.New("donnee incorrect")
	}
	fmt.Println("on va connecter", user)
	// Create session token
	token := uuid.New()
	sessions[token] = user.Id_user
	//var s = &model.Session{} // assigne a la machine les sessions n'est pas une bonne idee--a gerer
	fmt.Println("id user a", user.Id_user)
	err = dao.CreateSession(token, user.Id_user) //mais en realite c'est un update
	if err != nil {
		fmt.Println("y a erreur dans la creation de session")
		//return fmt.Errorf(" y a erreur dans la creation de session")
	}
	return nil
}

func Logout(token uuid.UUID) {
	delete(sessions, token)
}

func Authenticate(token uuid.UUID) (int, bool) {
	userID, exists := sessions[token]
	return userID, exists
}
