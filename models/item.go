package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type ItemInput struct {
	Data    json.RawMessage `json:"data"`
	VaultID uint            `json:"vault_id" binding:"required"`
}

type Item struct {
	gorm.Model
	Data    json.RawMessage `gorm:"type:json"`
	Vault   Vault           `gorm:"foreignKey:VaultID"` // Explicitly define the foreign key
	VaultID uint            // Foreign key field
}

type PublicItem struct {
	ID        uint            `json:"id"`
	Data      json.RawMessage `json:"data"`
	VaultID   uint            `json:"vault_id"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
}
