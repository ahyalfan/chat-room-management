package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

type Message struct {
	Content  string `json:"content"`
	Username string `json:"username"`
	RoomID   string `json:"room_id"`
}

//untuk menulis pesan kepada client yang bersangkutan
func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		// ada pesan maka return dalam bentuk json
		message, ok := <-c.Message
		if !ok {
			return
		}
		err := c.Conn.WriteJSON(message)
		if err != nil {
			return
		}
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// baca dari connection websocketnya
		// jadi semisal ada yang connect dan dia kasih sebuah message maka akan dibaca
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error : ", err)
			}
			break
		}
		// kirim ke hub broadcast message yg bersangkutan
		msg := &Message{
			Content:  string(m),
			Username: c.Username,
			RoomID:   c.RoomID,
		}
		hub.Broadcast <- msg
	}
}
