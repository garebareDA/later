package database

type Draft struct {
	ID int `gorm:"AUTO_INCREMENT" gorm:"primary_key"`
	UserID string
	Title string
	Content string
}