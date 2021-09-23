package chats

type ChatResponses struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Chat `json:"data"`
}

type ChatResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Chat   `json:"data"`
}
