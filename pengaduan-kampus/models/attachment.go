package models

type Attachment struct {
	ID          uint `gorm:"primaryKey"`
	ComplaintID uint
	FilePath    string
}
