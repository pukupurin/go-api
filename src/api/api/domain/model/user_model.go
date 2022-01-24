package model

import (
	"time"
)

type User struct {
	Id int64             `gorm:"id"`
	Name string          `gorm:"name"`

	CreatedAt *time.Time `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}
type Users []User

// NewUser User のコンストラクタ
func NewUser(name string) (*User, error) {
	
	user := &User{
		Name : name,
	}

	return user, nil
}