package controllers

import (
	"app/configs"
	"app/models/chats"
	"app/models/rooms"
	"net/http"

	"github.com/labstack/echo"
)

func ListRooms(c echo.Context) error {
	var roomCheck []rooms.Room
	userId := c.Param("userid")
	err := configs.DB.Order("updated_at DESC").Find(&roomCheck, "sender = ? OR receiver = ?", userId, userId).Error
	// err := configs.DB.Model(&roomCheck).Where("receiver = ? AND status = 0 ", userId).Association("chats")
	//err := configs.DB.Joins("JOIN chats on chats.room_id = rooms.id").Order("rooms.updated_at DESC").Where("rooms.sender = ? OR rooms.receiver = ?", userId, userId).Find(&roomCheck)
	// countUnread := configs.DB.Joins("JOIN chats on chats.room_id = rooms.id").Where("chats.status = 0").Find(&roomCheck).RowsAffected
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rooms.RoomResponses{
			false, "Failed Get List Room", nil,
		})
	}

	// var unread = map[string]int{"unread": int(countUnread)}
	// roomCheck = append(roomCheck, unread)
	return c.JSON(http.StatusOK, rooms.RoomResponses{
		true, "Success Get List Room", roomCheck,
	})
}

func ListChats(c echo.Context) error {
	var listChats []chats.Chat
	roomId := c.Param("roomid")
	err := configs.DB.Order("created_at DESC").Find(&listChats, "room_id = ?", roomId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, chats.ChatResponses{
			false, "Failed get chats in this room", nil,
		})
	}
	if updateStatus(roomId) == 0 {
		return c.JSON(http.StatusInternalServerError, chats.ChatResponses{
			false, "Failed update status", nil,
		})
	}
	return c.JSON(http.StatusOK, chats.ChatResponses{
		true, "Success get chats in room", listChats,
	})
}

func updateStatus(roomid string) int {
	var chatDB chats.Chat
	chatDB.Status = 1
	err := configs.DB.Where("status != 1 AND room_id = ?", roomid).Updates(&chatDB).Error
	if err != nil {
		return 0
	}
	return 1
}
