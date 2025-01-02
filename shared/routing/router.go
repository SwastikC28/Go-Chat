package routing

import "github.com/gorilla/mux"

func NewDefaultRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}
