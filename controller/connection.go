package controller

import (
	authentification "FirstProject/Authentification"
	"FirstProject/model"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ConnectUser(u *model.Users) error {
	err := authentification.Login(u.Nom_user, u.Password_user)
	if err != nil {
		fmt.Println("y a erreur dans la connexion")
		return err
	}
	fmt.Println("ca passe", u)
	return fmt.Errorf("Connexion etablie")
}

//dao.CreateSession(u.Id_tenant, u.Id_user, u.Role.Id_role, token)
/*	if err != nil {
		fmt.Println("y a erreur dans la creation de session")
		return fmt.Errorf("y a erreur dans la creation de session")
	}
	fmt.Println("Connexion et creation de session reussie de", u)
	return fmt.Errorf("Connexion et creation de session reussie pour")
	/*
		user, err := dao.GetUserByName(u.Nom_user)
		if err != nil {
			fmt.Println("y a erreur dans getUserByName")
			return fmt.Errorf("y a erreur dans GetUserByName")
		}
		if user.Nom_user != u.Nom_user {
			fmt.Println("Vos coordonnees ne correspondent a aucun user")
			return fmt.Errorf("Vos coordonnees ne correspondent a aucun user(nom)")
		}
		if user.Email_user != u.Email_user {
			fmt.Println("Vos coordonnees ne correspondent a aucun user")
			return fmt.Errorf("Vos coordonees ne correspond a aucun user(email)")
		}
		fmt.Println("On va connecter :", u)

		// Clé de signature
		key := []byte("clé_secrète")

		// Création d'un token
		token := jwt.New(jwt.SigningMethodHS256)

		// Ajouter des revendications (claims)
		claims := token.Claims.(jwt.MapClaims)
		claims["user_id"] = u.Id_user
		claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Expiration dans 1 heure

		// Signer le token
		tokenString, err := token.SignedString(key)
		if err != nil {
			fmt.Println("Erreur lors de la création du token:", err)
			return nil
		}

		//	jwt.Parse()

		fmt.Println("Token:", tokenString)
		var s = &model.Session{}
		err = dao.UpdateSession(s)


}
*/
// verification du token
func verifyToken(tokenString string, key []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Vérifier le type de signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Méthode de signature inattendue: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

/*
func main() {
    // ...
    tokenString := "votre_token_jwt"

    token, err := verifyToken(tokenString, key)
    if err != nil {
        fmt.Println("Erreur lors de la vérification du token:", err)
        return
    }

    if token.Valid {
        fmt.Println("Token valide")
        claims := token.Claims.(jwt.MapClaims)
        fmt.Println("User ID:", claims["user_id"])
    } else {
        fmt.Println("Token invalide")
    }
}
*/
