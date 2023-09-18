package rest

import (
	"FirstProject/model"
	"FirstProject/service"
	"FirstProject/utils"
	"net/http"
	//"github.com/dgrijalva/jwt-go" // indirect
)

// ConnectUser godoc
// @Summary Connexion pour un utilisateur
// @Description Connexion pour un utilisateur
// @Tags Connexion
// @Accept  json
// @Produce  json
// @Param Authorization header string true "api token"
// @Param User body model.Users true "project"
// @Success 200 {object}  model.Users
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /api/v1/users/con [post]
func ConnectUser(w http.ResponseWriter, r *http.Request) {
	var u model.Users
	err := readRequestBody(r.Body, &u)
	if err != nil {
		utils.LogWriter()
		serverResponse(w, err, nil)
		return
	}
	err = service.ConnectUser(&u)
	serverResponse(w, err, u)

}
