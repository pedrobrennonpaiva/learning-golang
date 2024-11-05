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
	{
		Uri:     "/posts/{id}/unlike",
		Method:  http.MethodPost,
		Handler: controllers.UnlikePost,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}/update",
		Method:  http.MethodGet,
		Handler: controllers.LoadPostUpdate,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdatePost,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeletePost,
		Auth:    true,
	},
}
