package routes

import (
	"Api-Aula1/controller"
	"net/http"
)

var booksRoutes = []Routes{

	{
		URI:    "/search",
		Method: http.MethodGet,
		Func:   controller.HandleSearch,
		Auth:   true,
	},

	{
		URI:    "/books",
		Method: http.MethodPost,
		Func:   controller.CreateBook,
		Auth:   true,
	},

	{
		URI:    "/users/{usuarioId}/books",
		Method: http.MethodGet,
		Func:   controller.FetchBooksByUser,
		Auth:   true,
	},

	{
		URI:    "/books/{bookId}",
		Method: http.MethodDelete,
		Func:   controller.DeleteBook,
		Auth:   true,
	},

	{
		URI:    "/books/{bookId}",
		Method: http.MethodPut,
		Func:   controller.UpdateBook,
		Auth:   true,
	},
}