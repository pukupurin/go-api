package model

import (
	"time"
)

type User struct {
	ID   string `gorm:"id;default:BIN_TO_UUID(UUID_TO_BIN(UUID(), 1))"`
	Name string `gorm:"name"`

	CreatedAt *time.Time `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}
type Users []User

// NewUser User のコンストラクタ
func NewUser(name string) (*User, error) {

	user := &User{
		Name: name,
	}

	return user, nil
}
