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
	Unreade string `json:"unread"`
}

type RoomResponsesFull struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    RoomData `json:"data"`
}
