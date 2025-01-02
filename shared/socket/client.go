package socket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	ID    uuid.UUID
	Conn  *websocket.Conn
	Pool  *Pool
	mutex *sync.Mutex
}

func NewClient(conn *websocket.Conn, pool *Pool, mutex *sync.Mutex) *Client {
	return &Client{
		ID:    uuid.NewV4(),
		Conn:  conn,
		Pool:  pool,
		mutex: mutex,
	}
}

type Message struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		m := Message{Type: string(msgType), Body: string(msg)}
		c.Pool.Broadcast <- m
	}
}
