package entity

type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name" gorm:"unique" validate:"min=3,unique"`
	Password string `json:"password" validate:"required,min=5"`
}
