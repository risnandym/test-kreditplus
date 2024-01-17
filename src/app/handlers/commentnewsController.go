package handlers

import (
	"kredit_plus/src/app/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentNewsInput struct {
	NewsID  uint   `json:"news_id"`
	Comment string `json:"comment"`
}

// GetAllCommentNews godoc
// @Summary Get all Comment.
// @Description Get a list of Comment.
// @Tags Comment News
// @Produce json
// @Success 200 {object} []entities.CommentsNews
// @Router /comments-news [get]
func GetAllCommentNews(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var comments []entities.CommentsNews
	db.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// CreateCommentNews godoc
// @Summary Create New Comment.
// @Description Creating a new Comment.
// @Tags Comment News
// @Param Body body commentNewsInput true "the body to create a new Comment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} entities.CommentsNews
// @Router /comments-news [post]
func CreateCommentNews(c *gin.Context) {
	// Validate input
	var input commentNewsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var rating entities.News
	if err := db.Where("id = ?", input.NewsID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NewsID not found!"})
		return
	}

	user, _ := c.Get("user")
	users := user.(entities.User)

	// Create Comment
	comment := entities.CommentsNews{
		NewsID:  input.NewsID,
		Name:    users.Username,
		Comment: input.Comment,
		UserID:  users.ID,
	}
	db.Create(&comment)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteCommentNews godoc
// @Summary Delete one Comment.
// @Description Delete a Comment by id.
// @Tags Comment News
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} map[string]boolean
// @Router /comments-news/{id} [delete]
func DeleteCommentNews(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var comment entities.CommentsNews
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
