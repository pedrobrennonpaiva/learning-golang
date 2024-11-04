package routes

import (
	"net/http"
	"webapp/internal/app/controllers"
)

var postsRoutes = []Route{
	{
		Uri:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}/like",
		Method:  http.MethodPost,
		Handler: controllers.LikePost,
		Auth:    true,
	},
}
