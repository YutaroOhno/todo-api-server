package entities

type Todo struct {
	ID int `gorm:"primary_key"`
	Title string `json:"title" gorm:"not null"`
	Text string `json:"text" gorm:"not null"`
}

// FindAllで取ったものをそのまま返しているが、これはoutputportを使った方が良いのだろうか？もしそれなら、このstructにjsonの記述はいらなそう。