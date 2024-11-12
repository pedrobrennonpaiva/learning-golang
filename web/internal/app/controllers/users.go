package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/internal/config"
	"webapp/internal/models"
	"webapp/internal/pkg"
	"webapp/internal/pkg/cookies"
	"webapp/internal/pkg/requests"
	"webapp/internal/pkg/responses"

	"github.com/gorilla/mux"
)

func LoadRegisterUser(w http.ResponseWriter, r *http.Request) {
	pkg.ExecuteTemplate(w, "register.html", nil)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nickname": r.FormValue("nickname"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.GetConfig().ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	nameOrNick = strings.Replace(nameOrNick, " ", "%20", -1)

	url := fmt.Sprintf("%s/users?user=%s", config.GetConfig().ApiUrl, nameOrNick)

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

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Err: err.Error()})
		return
	}

	pkg.ExecuteTemplate(w, "users.html", users)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userLoggedID, _ := strconv.ParseUint(cookie.ID, 10, 64)

	if userId == userLoggedID {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, err := models.GetFullUser(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}

	pkg.ExecuteTemplate(w, "user.html", struct {
		User         models.User
		UserLoggedID uint64
	}{
		User:         user,
		UserLoggedID: userLoggedID,
	})
}

func LoadProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie.ID, 10, 64)

	user, err := models.GetFullUser(userID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}

	pkg.ExecuteTemplate(w, "profile.html", user)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.GetConfig().ApiUrl, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.GetConfig().ApiUrl, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func LoadEditUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoadEditUser")
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie.ID, 10, 64)

	channel := make(chan models.User)
	go models.GetUser(channel, userID, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: "Error to get user"})
		return
	}

	pkg.ExecuteTemplate(w, "edit-user.html", user)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("EditUser")

	user, err := json.Marshal(models.User{
		Name:     r.FormValue("name"),
		Nickname: r.FormValue("nickname"),
		Email:    r.FormValue("email"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookies, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookies.ID, 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.GetConfig().ApiUrl, userID)
	response, err := requests.DoRequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func LoadChangePassword(w http.ResponseWriter, r *http.Request) {
	pkg.ExecuteTemplate(w, "change-password.html", nil)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	changePassword, err := json.Marshal(models.Password{
		Current: r.FormValue("currentPassword"),
		New:     r.FormValue("newPassword"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Err: err.Error()})
		return
	}

	cookies, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookies.ID, 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-password", config.GetConfig().ApiUrl, userID)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(changePassword))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookies, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookies.ID, 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.GetConfig().ApiUrl, userID)
	response, err := requests.DoRequestWithAuth(r, http.MethodDelete, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
