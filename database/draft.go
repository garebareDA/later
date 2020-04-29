package database

type Draft struct {
	ID int `gorm:"AUTO_INCREMENT" gorm:"primary_key"`
	UserID string
	Content string
}