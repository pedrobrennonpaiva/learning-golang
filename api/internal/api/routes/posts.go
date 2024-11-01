package routes

import "golang-api/internal/api/controllers"

var routesPost = []Route{
	{
		Uri:     "/posts",
		Method:  "POST",
		Handler: controllers.CreatePost,
		Auth:    true,
	},
	{
		Uri:     "/posts",
		Method:  "GET",
		Handler: controllers.GetPosts,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}",
		Method:  "GET",
		Handler: controllers.GetPostById,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}",
		Method:  "PUT",
		Handler: controllers.UpdatePost,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}",
		Method:  "DELETE",
		Handler: controllers.DeletePost,
		Auth:    true,
	},
	{
		Uri:     "/users/{id}/posts",
		Method:  "GET",
		Handler: controllers.GetPostsByUser,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}/like",
		Method:  "POST",
		Handler: controllers.LikePost,
		Auth:    true,
	},
	{
		Uri:     "/posts/{id}/unlike",
		Method:  "POST",
		Handler: controllers.UnlikePost,
		Auth:    true,
	},
}
