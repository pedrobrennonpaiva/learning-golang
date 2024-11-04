package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/internal/config"
	"webapp/internal/models"
	"webapp/internal/pkg/requests"
	"webapp/internal/pkg/responses"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(models.Post{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	url := fmt.Sprintf("%s/posts", config.GetConfig().ApiUrl)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(post))
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, "Post created successfully")
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.GetConfig().ApiUrl, postId)
	response, err := requests.DoRequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TreatError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
