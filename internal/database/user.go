package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid         uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();" json:"uuid"`
	FirstName    string         `gorm:"size:50" json:"first_name"`
	LastName     string         `gorm:"size:50" json:"last_name"`
	DOB          string         `gorm:"size:50" json:"dob"`
	HashPassword string         `gorm:"not null" json:"hash_password"`
	Email        string         `gorm:"size:50;unique" json:"email"`
	Phone        string         `gorm:"size:15" json:"phone"`
	About        string         `gorm:"size:10000" json:"about"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}

func (s *service) FindUsers() []User {
	var users []User
	s.db.Find(&users)
	return users
}

func (s *service) FindUser(uuid string) User {
	var user User
	s.db.Where("uuid = ?", uuid).First(&user)
	return user
}
