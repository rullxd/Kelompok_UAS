package models

type Complaint struct {
	ID          uint `gorm:"primaryKey"`
	StudentID   uint
	Title       string
	Description string
	Status      string

	Attachments []Attachment
	Comments    []Comment
}
