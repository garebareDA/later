package database

type Like struct {
	ID int `gorm:"AUTO_INCREMENT" gorm:"primary_key"`
	UserID string
	UUID string
}