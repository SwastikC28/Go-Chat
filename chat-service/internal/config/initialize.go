package config

import (
	"go-chat/internal/websocket"
	"shared/routing"
	"shared/socket"
	"sync"

	"github.com/gorilla/mux"
)

func Initialize() *mux.Router {
	router := routing.NewDefaultRouter()
	registerHandlers(router)

	return router
}

func registerHandlers(router *mux.Router) {
	pool := socket.NewPool()
	mutex := &sync.Mutex{}

	go pool.Start()

	chatSocketHandler := websocket.NewChatWebSocket(pool, mutex)
	chatSocketHandler.RegisterSocket(router)
}
