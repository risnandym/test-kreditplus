package controllers

import (
	"net/http"

	"final-project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentphoneInput struct {
	PhoneID uint   `json:"phone_id"`
	Comment string `json:"comment"`
}

// GetAllCommentPhone godoc
// @Summary Get all Comment.
// @Description Get a list of Comment.
// @Tags Comment Phone
// @Produce json
// @Success 200 {object} []models.CommentsPhone
// @Router /comments-phone [get]
func GetAllCommentPhone(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var comments []models.CommentsPhone
	db.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// CreateCommentPhone godoc
// @Summary Create New Comment.
// @Description Creating a new Comment.
// @Tags Comment Phone
// @Param Body body commentphoneInput true "the body to create a new Comment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.CommentsPhone
// @Router /comments-phone [post]
func CreateCommentPhone(c *gin.Context) {
	// Validate input
	var input commentphoneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var rating models.Phone
	if err := db.Where("id = ?", input.PhoneID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhoneID not found!"})
		return
	}

	user, _ := c.Get("user")
	users := user.(models.User)
	// Create Comment
	comment := models.CommentsPhone{
		PhoneID: input.PhoneID,
		Name:    users.Username,
		Comment: input.Comment,
		UserID:  users.ID,
	}

	db.Create(&comment)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteCommentPhone godoc
// @Summary Delete one Comment.
// @Description Delete a Comment by id.
// @Tags Comment Phone
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} map[string]boolean
// @Router /comments-phone/{id} [delete]
func DeleteCommentPhone(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var comment models.CommentsPhone
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
