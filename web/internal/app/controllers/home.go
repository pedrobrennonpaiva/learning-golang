package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/internal/config"
	"webapp/internal/models"
	"webapp/internal/pkg"
	"webapp/internal/pkg/cookies"
	"webapp/internal/pkg/requests"
	"webapp/internal/pkg/responses"
)

func Home(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf("%s/posts", config.GetConfig().ApiUrl)

	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	authUser, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(authUser.ID, 10, 64)

	pkg.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userId,
	})
}
