package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"name" binding:"required"`
	Email    string   `json:"email" gorm:"unique" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Groups   []*Group `json:"groups" gorm:"many2many:user_groups;"`
}

type PublicUser struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Groups    []string `json:"groups"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func (u *User) ToPublicUser() PublicUser {

	var groups []string

	for _, group := range u.Groups {
		groups = append(groups, group.Name)
	}

	return PublicUser{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Groups:    groups,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}

func (u *User) EncryptPassword() (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	_, err = u.EncryptPassword()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(user *User) (*User, error) {
	if user.Password != "" {
		u.Password = user.Password
		_, err := u.EncryptPassword()
		if err != nil {
			return nil, err
		}
	}

	if user.Name != "" {
		u.Name = user.Name
	}

	if user.Email != "" {
		u.Email = user.Email
	}

	return u, nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
