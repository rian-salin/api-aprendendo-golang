package controller

import (
	"Api-Aula1/auth"
	"Api-Aula1/models"
	"Api-Aula1/persistency"
	"Api-Aula1/repository"
	"Api-Aula1/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func CreateBook(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book models.Book
	if err = json.Unmarshal(bodyRequest, &book); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	book.UserID = userID

	if err = book.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewBooksRepo(db)
	book.ID, err = repo.Create(book)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, book)
}

func FetchBooksByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["usuarioId"], 10, 64)
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

	repo := repository.NewBooksRepo(db)
	books, err := repo.FetchByUserID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID, err := strconv.ParseUint(params["bookId"], 10, 64)
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

	repo := repository.NewBooksRepo(db)
	if err = repo.Delete(bookID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseUint(params["bookId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book models.Book
	if err = json.Unmarshal(bodyRequest, &book); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewBooksRepo(db)
	if err = repo.Update(bookID, book); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}