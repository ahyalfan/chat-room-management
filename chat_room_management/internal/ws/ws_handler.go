package ws

import (
	"ahyalfan.my.id/chat_rom_management/dto"
	"ahyalfan.my.id/chat_rom_management/internal/config"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	hub *Hub
	cnf *config.Config
}

func NewHandler(ap *fiber.App, h *Hub, authMid fiber.Handler, cnf *config.Config) {
	handler := Handler{hub: h, cnf: cnf}

	app := ap.Group("/api/v1/", authMid)

	app.Post("/create-room", handler.CreateRoom)
	app.Get("/get-room", handler.GetRooms)
	app.Get("/get-room-client/:roomId", handler.GetRoomClients)
	ap.Use("/ws", handler.AllowUpgrade)
	ap.Get("/ws/join-room/:roomId", websocket.New(handler.JoinRoom))
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(ctx *fiber.Ctx) error {
	var req CreateRoomReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request",
		})
	}

	room := &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	h.hub.Rooms[req.ID] = room
	return ctx.JSON(map[string]interface{}{
		"message": "Room created successfully",
		"room":    room,
	})
}

func (h *Handler) AllowUpgrade(ctx *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(ctx) {
		return ctx.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (h *Handler) JoinRoom(c *websocket.Conn) {
	defer c.Close()

	roomId := c.Params("roomId")
	clientId := c.Query("userId")
	username := c.Query("username")
	authMsg := c.Query("token")
	cl := &Client{
		Conn:     c,
		RoomID:   roomId,
		Username: username,
		ID:       clientId,
		Message:  make(chan *Message),
	}

	_, err := jwt.Parse(authMsg, func(token *jwt.Token) (any, error) {
		return []byte(h.cnf.JWT.Key), nil
	})

	if err != nil {
		cl.Conn.Close()
		return
	}

	room := h.hub.Rooms[roomId]
	if room == nil {
		cl.Conn.Close()
		return
	}
	for _, v := range room.Clients {
		if v.ID == clientId {
			v.Message <- &Message{
				Content:  "User has reconnected",
				Username: username,
				RoomID:   roomId,
			}
			cl.Conn.Close()
			return
		}
	}

	m := &Message{
		Content:  "A new user has joined the room",
		Username: username,
		RoomID:   roomId,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(ctx *fiber.Ctx) error {
	rooms := make([]RoomRes, 0, len(h.hub.Rooms))
	for _, room := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   room.ID,
			Name: room.Name,
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "List of rooms",
		"rooms":   rooms,
	})
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetRoomClients(ctx *fiber.Ctx) error {
	var clients []ClientRes
	roomId := ctx.Params("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		return ctx.JSON(dto.CreateResponseSuccess(fiber.StatusOK, fiber.Map{
			"data": clients,
		}))
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	return ctx.JSON(dto.CreateResponseSuccess(fiber.StatusOK, fiber.Map{
		"data": clients,
	}))
}

func (h *Handler) CloseRoom(c *websocket.Conn) {
	defer c.Close()
	roomID := c.Query("roomId")
	// client := c.Query("clientId")
	cl := &Client{
		Conn:   c,
		RoomID: roomID,
	}

	if room, ok := h.hub.Rooms[roomID]; ok {
		for _, v := range room.Clients {
			if v.ID == cl.ID {
				delete(room.Clients, cl.ID)
				break
			}
		}

		h.hub.Unregister <- cl
		m := &Message{
			Content:  "A user has left the room",
			Username: cl.Username,
			RoomID:   roomID,
		}
		h.hub.Broadcast <- m
	}
}
