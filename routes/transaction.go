package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transactions-charity/{id}", h.GetTransactionByCharity).Methods("GET")
	r.HandleFunc("/transactionsPending/{id}", h.GetTransactionByCharityPending).Methods("GET")
	r.HandleFunc("/transactions-user/{id}", middleware.Auth(h.GetTransactionByUser)).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
