package routes

import (
	"Api-Aula1/controller"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
		Auth:   false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodGet,
		Func:   controller.FetchUser,
		Auth:   true,
	},
	{
		URI:    "/users/{userID}",
	
		Method: http.MethodPut, 
		Func:   controller.UpdateUser,
		Auth:   true,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
		Auth:   true,
	},
}