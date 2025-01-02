package websocket

import (
	"net/http"
	"sync"

	"shared/socket"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type ChatWebSocket struct {
	pool  *socket.Pool
	mutex *sync.Mutex
}

func NewChatWebSocket(pool *socket.Pool, mutex *sync.Mutex) socket.WebSocket {
	return &ChatWebSocket{
		pool:  pool,
		mutex: mutex,
	}
}

func (chatsocket *ChatWebSocket) RegisterSocket(router *mux.Router) {
	router.HandleFunc("/ws", chatsocket.Upgrade)
}

func (chatsocket *ChatWebSocket) Upgrade(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context()).With().Str("action", "chatwebsocket:upgrade").Logger()
	conn, err := socket.CreateWebSocket(w, r)
	if err != nil {
		logger.Err(err).Msg("There was some error upgrading connection")
		return
	}

	client := socket.NewClient(conn, chatsocket.pool, chatsocket.mutex)
	chatsocket.pool.Register <- client

	client.Read()
}
