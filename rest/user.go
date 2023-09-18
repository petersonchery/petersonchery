package rest

import (
	authentification "FirstProject/Authentification"
	"FirstProject/model"
	"FirstProject/service"
	"FirstProject/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	createUserPermission = iota + 1
	getUserPermission
	deleteUserPermission
	updateUserPermission
)

func AuthAndAuto(idUser int) error {

	fmt.Println("voici l'idUser", idUser)
	if idUser == 0 {
		fmt.Println("Cet identifiant n'existe pas(token)")
		return nil
	}

	isExist, err := service.HasPermission(idUser, createUserPermission)
	if err != nil {
		utils.LogWriter()
		fmt.Println("ya erreur dans la fonction hasPermission")
		return nil
	}
	if isExist == false {
		fmt.Println("vous avez pas la permission de creer user")
		return nil
	}
	return nil
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user
// @Tags User
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param User body model.Users true "project"
// @Success 200 {object}  model.Users
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u model.Users
	//a besoin de la permission 1
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	AuthAndAuto(idUser)
	err := readRequestBody(r.Body, &u)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, err, nil)
		return
	}
	fmt.Println("Je vois que vous avez cette permission donc vous allex creer cet user")

	err = service.CreateUser(&u)
	serverResponse(w, err, u)
}

// GetUserById godoc
// @Summary Get a user by his ID
// @Description Get a user by his ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_User path int true "ID User"
// @Success 200 {object}  model.Users
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/users{id} [get]
func GetUserById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	var u = &model.Users{Id_user: id}
	if err != nil {
		utils.LogWriter()
		fmt.Println("y a erreur")
		serverResponse(w, fmt.Errorf("y a erreur"), nil)
	}

	err = service.GetUserById(u)
	serverResponse(w, err, u)

}

// UserById godoc
// @Summary Delete a user by his ID
// @Description Delete a user by his ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_User path int true "ID User"
// @Success 200 {object}  model.Users
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/users/sup/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	fmt.Println(token)
	idUser := authentification.GetUser(token)

	AuthAndAuto(idUser)
	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, fmt.Errorf("Id non valide"), nil)
		return
	}
	var u = &model.Users{Id_user: id}
	err = service.DeleteUser(u)
	serverResponse(w, err, u)
}

// UpdateUserById godoc
// @Summary Update a user by his ID
// @Description Update a user by his ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param ID_User path int true "ID User"
// @Param User body model.Users true "project"
// @Success 200 {object}  model.Users
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/users/up/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idReq := vars["id"]
	id, err := strconv.Atoi(idReq)
	if err != nil {
		utils.LogWriter()
		fmt.Println("y a erreur")
		serverResponse(w, fmt.Errorf("y a erreur"), nil)
		return
	}
	var u = &model.Users{}
	err = readRequestBody(r.Body, u)
	if err != nil {
		utils.LogWriter()
		fmt.Println("y a erreur dans rest.UpdateUser")
	}
	u.Id_user = id
	fmt.Println("id user", u.Id_user)
	err = service.UpdateUser(u)
	serverResponse(w, err, u)

}

// ListerUser godoc
// @Summary Lister tous les Utilisateurs
// @Description Lister tous les utilisateurs
// @Tags User
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Success 200 {object}  model.Users
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/users/lis [get]
func ListerUser(w http.ResponseWriter, r *http.Request) {

	// var biblio = &model.Bibliotheque{}
	// biblios, err := service.ListerBibliotheque(biblio)

	// for _, bib := range biblios {
	// 	fmt.Printf("ID: %d, Username: %d, Email: %s\n", bib.Id_bibliotheque, bib.Id_tenant, bib.Nom_bibliotheque)
	// }
	// serverResponse(w, err, biblios)

}
