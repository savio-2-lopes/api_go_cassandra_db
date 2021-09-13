package Users

import (
	"github.com/go-chi/chi"
)

func UsersRoutes() *chi.Mux {
	router := chi.NewRouter() 
	router.Get("/getUser", getHandler)
	router.Post("/register", postHandler)
	return router
}