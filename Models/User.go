package Models

import (
	"github.com/joshbedo/serverless-go/Config"
)

type User struct {
	UserID   uint   `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (b *User) TableName() string {
	return "users"
}

func GetAllUsers(b *[]User) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}
