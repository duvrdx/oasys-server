package controllers

import (
	"net/http"

	"github.com/duvrdx/oasys-server/config"
	"github.com/duvrdx/oasys-server/models"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context) {
	var groups []models.Group
	config.DB.Find(&groups)

	publicGroups := make([]models.PublicGroup, len(groups))
	for i, group := range groups {
		publicGroups[i] = group.ToPublicGroup()
	}

	c.JSON(http.StatusOK, gin.H{"data": publicGroups})
}

func CreateGroup(c *gin.Context) {
	var input models.Group
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input.ToPublicGroup()})
}

func GetGroup(c *gin.Context) {
	var group models.Group
	if err := config.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": group.ToPublicGroup()})
}

func UpdateGroup(c *gin.Context) {
	var group models.Group
	if err := config.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Group
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupUpdated, err := group.Update(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groupUpdated.ToPublicGroup()})
}

func DeleteGroup(c *gin.Context) {
	var group models.Group
	if err := config.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&group)
	c.JSON(http.StatusOK, gin.H{"data": "Group deleted!"})
}

func AddUserToGroup(c *gin.Context) {
	var group models.Group
	if err := config.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Model(&group).Association("User").Append(&user)
	c.JSON(http.StatusOK, gin.H{"data": group.ToPublicGroup()})
}

func RemoveUserFromGroup(c *gin.Context) {
	var group models.Group
	if err := config.DB.Where("id = ?", c.Param("id")).First(&group).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var user models.User
	if err := config.DB.Where("id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Model(&group).Association("User").Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": group.ToPublicGroup()})
}
