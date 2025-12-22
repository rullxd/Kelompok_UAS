package models

type Comment struct {
	ID          uint `gorm:"primaryKey"`
	ComplaintID uint
	UserID      uint
	Message     string
}
