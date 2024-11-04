package controllers

import (
	"encoding/json"
	"golang-api/internal/models"
	"golang-api/internal/models/responses"
	"golang-api/internal/pkg/auth"
	"golang-api/internal/pkg/security"
	"golang-api/internal/repository"
	"golang-api/internal/repository/database"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	userDb, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userDb.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	authResponse, err := auth.CreateToken(userDb.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, authResponse)
}
