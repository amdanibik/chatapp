package rooms

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	Sender    string         `json:"sender"`
	Receiver  string         `json:"receiver"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
