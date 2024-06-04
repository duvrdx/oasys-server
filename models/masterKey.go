package models

import (
	oc_models "github.com/duvrdx/oasys-crypto/pkg/models"
	"gorm.io/gorm"
)

type MasterKeyInput struct {
	Key       string `json:"key" binding:"required"`
	KeyLength int    `json:"key_length"`
}

type MasterKey struct {
	gorm.Model
	Key       string `json:"key"`
	Salt      string `json:"salt"`
	KeyLength int    `json:"key_length"`
}

type PublicMasterKey struct {
	ID  uint   `json:"id"`
	Key string `json:"key"`
}

func (m *MasterKey) ToPublic() PublicMasterKey {
	return PublicMasterKey{
		ID:  m.ID,
		Key: m.Key,
	}
}

func (m *MasterKeyInput) ToMasterKey() *MasterKey {
	masterKey, err := oc_models.NewMasterKey(m.Key, m.KeyLength)
	if err != nil {
		return nil
	}

	return &MasterKey{
		Key:       string(masterKey.GetKey()),
		Salt:      string(masterKey.GetSalt()),
		KeyLength: masterKey.GetKeySize(),
	}
}

func (m *MasterKey) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}
