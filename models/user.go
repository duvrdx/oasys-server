package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"` // Validate email format
	Password string `json:"password" binding:"required,min=6"`            // Minimum password length of 6
}

type User struct {
	gorm.Model
	Name        string
	Email       string
	Password    string
	Groups      []*Group `gorm:"many2many:user_groups;"`
	OwnedVaults []Vault  `gorm:"foreignKey:OwnerID"`
	Vaults      []*Vault `gorm:"many2many:user_vaults;"`
}

type PublicUser struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Groups      []string `json:"groups"`
	OwnedVaults []uint   `json:"owned_vaults"`
	Vaults      []uint   `json:"vaults"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

func (u *User) ToPublicUser() PublicUser {

	var groups []string
	var owned_vaults []uint
	var vaults []uint

	for _, vault := range u.Vaults {
		vaults = append(vaults, vault.ID)
	}

	for _, vault := range u.OwnedVaults {
		owned_vaults = append(owned_vaults, vault.ID)
	}

	for _, group := range u.Groups {
		groups = append(groups, group.Name)
	}

	return PublicUser{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Groups:      groups,
		OwnedVaults: owned_vaults,
		Vaults:      vaults,
		CreatedAt:   u.CreatedAt.String(),
		UpdatedAt:   u.UpdatedAt.String(),
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

func (u *UserInput) ToUser() *User {
	return &User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
