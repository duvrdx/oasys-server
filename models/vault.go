package models

import "gorm.io/gorm"

// VaultInput struct for receiving vault creation data
type VaultInput struct {
	Name      string `json:"name" binding:"required"`
	OwnerId   uint   `json:"user_id" binding:"required"` // Use "user_id" for consistency
	MasterKey string `json:"master_key" binding:"required"`
}

type Vault struct {
	gorm.Model
	Name        string
	OwnerID     uint
	Owner       User     `gorm:"foreignKey:OwnerID"`
	Users       []*User  `gorm:"many2many:user_vaults;"`
	Groups      []*Group `gorm:"many2many:group_vaults;"`
	MasterKeyID uint
	MasterKey   MasterKey `gorm:"foreignKey:MasterKeyID"`
	Items       []Item    `gorm:"foreignKey:VaultID"` // Explicitly define the foreign key
}

type PublicVault struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Users     []uint       `json:"users"` // Removed unnecessary nesting (consider User struct instead)
	OwnerID   uint         `json:"owner"` // Use OwnerID for consistency
	MasterKey uint         `json:"master_key"`
	Items     []PublicItem `json:"items"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}
