package controller

import (
	"Api-Aula1/models"
	"Api-Aula1/persistency"
	"Api-Aula1/repository"
	"Api-Aula1/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var newUser models.Users
	if err = json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	// Chama os métodos de preparação do User
	if err = newUser.Prepare("register"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	
	repo := repository.NewUsersRepo(db)
	newUser.ID, err = repo.Create(newUser)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	responses.JSON(w, http.StatusCreated, newUser)
}

func FetchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := persistency.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepo(db)
	user, err := repo.FetchByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	// Se não achou nada, o ID na struct vai estar zerado
	if user.ID == 0 {
		responses.Error(w, http.StatusNotFound, errUserNotFound())
		return
	}

	// Retorna o usuário exatamente como está salvo no banco
	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {

}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {}

// helpers básicos dos erros
func errMissingID() error {
	return errors.New("informe o id do usuário (?id=)")
}

func errInvalidID() error {
	return errors.New("id inválido")
}

func errUserNotFound() error {
	return errors.New("usuário não encontrado")
}
