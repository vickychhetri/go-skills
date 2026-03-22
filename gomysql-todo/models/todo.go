package models

type Todo struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:200;unique;column:title;"`
	Description string `gorm:"column:description;"`
	UserID      uint
	Done        bool `gorm:"column:done;default:false;"`
}

func (Todo) TableName() string {
	return "todos"
}
