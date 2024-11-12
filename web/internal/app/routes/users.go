package routes

import (
	"net/http"
	"webapp/internal/app/controllers"
)

var usersRoutes = []Route{
	{
		Uri:     "/register",
		Method:  http.MethodGet,
		Handler: controllers.LoadRegisterUser,
		Auth:    false,
	},
	{
		Uri:     "/register",
		Method:  http.MethodPost,
		Handler: controllers.RegisterUser,
		Auth:    false,
	},
	{
		Uri:     "/search-users",
		Method:  http.MethodGet,
		Handler: controllers.SearchUsers,
		Auth:    true,
	},
	{
		Uri:     "/users/{userId}",
		Method:  http.MethodGet,
		Handler: controllers.SearchUser,
		Auth:    true,
	},
	{
		Uri:     "/profile",
		Method:  http.MethodGet,
		Handler: controllers.LoadProfile,
		Auth:    true,
	},
	{
		Uri:     "/users/{userId}/follow",
		Method:  http.MethodPost,
		Handler: controllers.FollowUser,
		Auth:    true,
	},
	{
		Uri:     "/users/{userId}/unfollow",
		Method:  http.MethodPost,
		Handler: controllers.UnfollowUser,
		Auth:    true,
	},
	{
		Uri:     "/edit-user",
		Method:  http.MethodGet,
		Handler: controllers.LoadEditUser,
		Auth:    true,
	},
	{
		Uri:     "/edit-user",
		Method:  http.MethodPut,
		Handler: controllers.EditUser,
		Auth:    true,
	},
	{
		Uri:     "/change-password",
		Method:  http.MethodGet,
		Handler: controllers.LoadChangePassword,
		Auth:    true,
	},
	{
		Uri:     "/change-password",
		Method:  http.MethodPost,
		Handler: controllers.ChangePassword,
		Auth:    true,
	},
	{
		Uri:     "/delete-user",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
		Auth:    true,
	},
}
