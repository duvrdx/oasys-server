package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name string  `json:"name" binding:"required"`
	User []*User `json:"users" gorm:"many2many:user_groups;"`
}

type PublicGroup struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Users     []PublicUser `json:"users"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}

func (g *Group) ToPublicGroup() PublicGroup {
	var users []PublicUser

	for _, user := range g.User {
		users = append(users, user.ToPublicUser())
	}

	return PublicGroup{
		ID:        g.ID,
		Name:      g.Name,
		Users:     users,
		CreatedAt: g.CreatedAt.String(),
		UpdatedAt: g.UpdatedAt.String(),
	}
}

func (g *Group) Update(group *Group) (*Group, error) {
	g.Name = group.Name
	return g, nil
}

func (g *Group) BeforeSave(tx *gorm.DB) (err error) {
	return nil
}

func (g *Group) AfterSave(tx *gorm.DB) (err error) {
	return nil
}
