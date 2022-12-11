package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/gorilla/mux"
)

func CharityRoutes(r *mux.Router) {
	charityRepository := repositories.RepositoryCharity(mysql.DB)
	h := handlers.HandlerCharity(charityRepository)

	r.HandleFunc("/charities", h.FindCharities).Methods("GET")
	r.HandleFunc("/charity/{id}", middleware.Auth(h.GetCharity)).Methods("GET")
	r.HandleFunc("/charity-user/{id}",middleware.Auth(h.GetCharityByUser)).Methods("GET")
	r.HandleFunc("/charity", middleware.Auth(middleware.UploadFile(h.CreateCharity))).Methods("POST")
	r.HandleFunc("/charity/{id}", middleware.Auth(h.UpdateCharity)).Methods("PATCH")
	r.HandleFunc("/charity/{id}", middleware.Auth(h.DeleteCharity)).Methods("DELETE")
}