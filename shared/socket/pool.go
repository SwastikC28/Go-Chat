package socket

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

var logger = zerolog.Ctx(context.Background()).With().Str("module", "pool").Logger()

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			logger.Info().Msg(fmt.Sprintf("Total Connection pool - %d.", len(pool.Clients)))
			for k := range pool.Clients {
				k.Conn.WriteJSON(Message{Type: "1", Body: "New user joined"})
			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			logger.Info().Msg(fmt.Sprintf("Total Connection pool - %d.", len(pool.Clients)))
			for k := range pool.Clients {
				k.Conn.WriteJSON(Message{Type: "2", Body: "User Disconnected"})
			}

		case message := <-pool.Broadcast:
			fmt.Print("Broadcasting a message")
			for k := range pool.Clients {
				if err := k.Conn.WriteJSON(message); err != nil {
					logger.Err(err).Msg("Error writing JSON")
					return
				}

			}
		}

	}
}
