package routes

import (
	"golang-api/internal/api/controllers"
)

var routesUser = []Route{
	{
		Uri:     "/users",
		Method:  "GET",
		Handler: controllers.GetUsers,
		Auth:    true,
	},
	{
		Uri:     "/users/{id}",
		Method:  "GET",
		Handler: controllers.GetUserById,
		Auth:    true,
	},
	{
		Uri:     "/users",
		Method:  "POST",
		Handler: controllers.CreateUser,
		Auth:    false,
	},
	{
		Uri:     "/users/{id}",
		Method:  "PUT",
		Handler: controllers.UpdateUser,
		Auth:    true,
	},
	{
		Uri:     "/users/{id}",
		Method:  "DELETE",
		Handler: controllers.DeleteUser,
		Auth:    true,
	},
	{
		Uri:     "/users/{id}/follow",
		Method:  "POST",
		Handler: controllers.FollowUser,
		Auth:    true,
	},
	{
		Uri:     "/users/{id}/unfollow",
		Method:  "POST",
		Handler: controllers.UnfollowUser,
		Auth:    true,
	},
	{
		Uri:     "/users/{id}/followers",
		Method:  "GET",
		Handler: controllers.GetFollowers,
		Auth:    true,
	},
	{
		Uri:     "/users/{id}/following",
		Method:  "GET",
		Handler: controllers.GetFollowing,
		Auth:    true,
	},
	{
		Uri:     "/users/{id}/update-password",
		Method:  "POST",
		Handler: controllers.UpdatePassword,
		Auth:    true,
	},
}
