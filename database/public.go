package database

type Public struct {
	ID int `gorm:"AUTO_INCREMENT" gorm:"primary_key"`
	UserID string
	UserName string
	Title string
	Content string
	Like int
	UUID string
}