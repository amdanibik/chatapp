package rooms

type RoomRegister struct {
	Id       uint   `form:"id"`
	Sender   string `form:"sender"`
	Receiver string `form:"receiver"`
}
