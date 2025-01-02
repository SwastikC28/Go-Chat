package websocket

import (
	"net/http"

	"shared/socket"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type ChatWebSocket struct {
}

func NewChatWebSocket() socket.WebSocket {
	return &ChatWebSocket{}
}

func (chatsocket *ChatWebSocket) RegisterSocket(router *mux.Router) {
}

func (chatsocket *ChatWebSocket) Upgrade(w http.ResponseWriter, r *http.Request) {
	logger := zerolog.Ctx(r.Context()).With().Str("action", "chatwebsocket:upgrade").Logger()
	_, err := socket.CreateWebSocket(w, r)
	if err != nil {
		logger.Err(err).Msg("There was some error upgrading connection")
		return
	}

}
