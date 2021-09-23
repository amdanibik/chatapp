package chats

type ChatRecord struct {
	Id       uint   `form:"id"`
	Sender   string `form:"sender"`
	Receiver string `form:"receiver"`
	Message  string `form:"message"`
}
