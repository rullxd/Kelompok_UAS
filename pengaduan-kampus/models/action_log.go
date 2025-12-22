package models

import "time"

type ActionLog struct {
	ID          uint `gorm:"primaryKey"`
	ComplaintID uint
	AdminID     uint
	Action      string
	CreatedAt   time.Time
}
