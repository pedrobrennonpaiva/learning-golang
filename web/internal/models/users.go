package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/internal/config"
	"webapp/internal/pkg/requests"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Followers []User    `json:"followers,omitempty"`
	Following []User    `json:"following,omitempty"`
	Posts     []Post    `json:"posts,omitempty"`
}

func GetFullUser(userId uint64, r *http.Request) (User, error) {
	channelUser := make(chan User)
	channelFollowers := make(chan []User)
	channelFollowing := make(chan []User)
	channelPosts := make(chan []Post)

	go GetUser(channelUser, userId, r)
	go GetFollowers(channelFollowers, userId, r)
	go GetFollowing(channelFollowing, userId, r)
	go GetPosts(channelPosts, userId, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-channelUser:
			if userLoaded.ID == 0 {
				return User{}, errors.New("error to get user")
			}
			user = userLoaded
		case followersLoaded := <-channelFollowers:
			if followersLoaded == nil {
				fmt.Println(followersLoaded)
				return User{}, errors.New("error to get followers")
			}
			followers = followersLoaded
		case followingLoaded := <-channelFollowing:
			if followingLoaded == nil {
				return User{}, errors.New("error to get following")
			}
			following = followingLoaded
		case postsLoaded := <-channelPosts:
			if postsLoaded == nil {
				return User{}, errors.New("error to get posts")
			}
			posts = postsLoaded
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func GetUser(channel chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.GetConfig().ApiUrl, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func GetFollowers(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.GetConfig().ApiUrl, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- []User{}
		return
	}

	channel <- followers
}

func GetFollowing(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.GetConfig().ApiUrl, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- []User{}
		return
	}

	channel <- following
}

func GetPosts(channel chan<- []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.GetConfig().ApiUrl, userId)
	response, err := requests.DoRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- []Post{}
		return
	}

	channel <- posts
}
