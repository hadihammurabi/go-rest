package table

import (
	"go-rest/entity"
)

const NameUsers string = "auth.users"

// User model
type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"-"`
}

// BeforeCreate func
// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// id, err := uuid.NewRandom()
// u.ID = id
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
// 	u.Password = string(hashedPassword)
// 	return
// }

// BeforeSave func
// func (u *User) BeforeSave(tx *gorm.DB) (err error) {
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
// 	u.Password = string(hashedPassword)
// 	return
// }

// ToEntity func
func (t User) ToEntity() *entity.User {
	return &entity.User{
		ID:       t.ID,
		Email:    t.Email,
		Password: t.Password,
	}
}

// UserFromEntity func
func UserFromEntity(e *entity.User) *User {
	return &User{
		Base: Base{
			ID: e.ID,
		},
		Email:    e.Email,
		Password: e.Password,
	}
}
