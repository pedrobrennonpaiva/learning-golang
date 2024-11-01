package routes

import "golang-api/internal/api/controllers"

var routesLogin = []Route{
	{
		Uri:     "/login",
		Method:  "POST",
		Handler: controllers.Login,
		Auth:    false,
	},
}
