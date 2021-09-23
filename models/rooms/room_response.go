package rooms

type RoomResponses struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type RoomResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}

type RoomData struct {
	Id       int    `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Unreade  string `json:"unread"`
}

type RoomResponsesFull struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []RoomData `json:"data"`
}
