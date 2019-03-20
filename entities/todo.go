package entities

type Todo struct {
	ID int `gorm:"primary_key"`
	Title string `gorm:"not null"`
	Text string `gorm:"not null"`
}