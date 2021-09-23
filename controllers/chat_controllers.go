package controllers

import (
	"app/configs"
	"app/models/chats"
	"app/models/rooms"
	"net/http"

	"github.com/labstack/echo"
)

func SendChat(c echo.Context) error {
	var chatRecord chats.ChatRecord
	c.Bind(&chatRecord)
	roomId := checkRoom(chatRecord.Sender, chatRecord.Receiver)
	var chatDB chats.Chat
	chatDB.RoomId = roomId
	chatDB.Sender = chatRecord.Sender
	chatDB.Receiver = chatRecord.Receiver
	chatDB.Message = chatRecord.Message
	inputChat := configs.DB.Create(&chatDB).Error
	if inputChat != nil {
		return c.JSON(http.StatusInternalServerError, chats.ChatResponses{
			false, "Failed Send chat", nil,
		})
	}

	return c.JSON(http.StatusOK, chats.ChatResponse{
		true, "Success Send Chat", chatDB,
	})
}

func checkRoom(sender, receiver string) int {
	var roomRecord rooms.Room
	var roomId int
	roomCheckSR := configs.DB.First(&roomRecord, "sender = ? AND receiver = ?", sender, receiver).Error
	if roomCheckSR != nil {
		roomCheckRS := configs.DB.First(&roomRecord, "sender = ? AND receiver = ? ", receiver, sender).Error
		if roomCheckRS != nil {
			roomRecord.Sender = sender
			roomRecord.Receiver = receiver
			cRoom := configs.DB.Create(&roomRecord).Error
			if cRoom != nil {
				return 0
			}
			roomId = int(roomRecord.Id)
			return roomId
		}
		roomId = int(roomRecord.Id)
		return roomId
	}
	roomId = int(roomRecord.Id)
	return roomId
}
