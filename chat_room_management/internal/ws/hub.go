package ws

import "fmt"

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5), // hanya menerima 5 penampungan message
	}
}

func (h *Hub) Run() {
	// lakukan perulangan
	for {
		// pilih mana yg mau dipilih
		select {
		// jika user mau join room
		case cl := <-h.Register:
			fmt.Println("test registration")
			// kita cek dia mau di room mana
			if _, ok := h.Rooms[cl.RoomID]; ok {

				r := h.Rooms[cl.RoomID]

				// kita pastikan user tersebut tidak ada di room yg mau dia daftar
				if _, ok := r.Clients[cl.ID]; !ok {
					// baru kita masukan ke roomnya jika tidak ada
					r.Clients[cl.ID] = cl
				}
			}
		// jika user mau keluar room
		case cl := <-h.Unregister:
			// kita cek ada tidak roomnya
			if _, ok := h.Rooms[cl.RoomID]; ok {
				// jika ada, cek lagi dia ada gak di room tersebut
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					// jika ada, kita hapus dia dari roomnya
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{ // kirim pesan bahwa akan keluar
							Content:  "user left the chat room",
							Username: cl.Username,
							RoomID:   cl.RoomID,
						}
					}
					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}
		// jika ada yg mau kirim pesan
		case msg := <-h.Broadcast:
			fmt.Println("test broadcast")
			//  cek ada roomnya tidak
			if _, ok := h.Rooms[msg.RoomID]; ok {

				// lakukan pengulangan sebanyak jumplah client roomnya
				for _, cl := range h.Rooms[msg.RoomID].Clients {
					// kita kirimkan pesan di setiap client
					cl.Message <- msg
				}
			}
		}
	}
}
