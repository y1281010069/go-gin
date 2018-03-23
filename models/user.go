package models

type User struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
}

func Login(maps interface {}) (user []User) {
	db.Where(maps).Find(&user)
	return
}