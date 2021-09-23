package chats

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	RoomId    int            `json:"roomid"`
	Sender    string         `json:"sender"`
	Receiver  string         `json:"receiver"`
	Message   string         `json:"message"`
	Status    int            `json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
