package models

type User struct {
	ID       uint   `gorm:"primaryKey;column:id;"`
	Username string `gorm:"size:100;unique;column:username;"`
	Password string `gorm:"column:password;"`
}

func (User) TableName() string {
	return "users"
}
