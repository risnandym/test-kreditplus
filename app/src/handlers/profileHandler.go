package handlers

import (
	"kredit_plus/app/src/contract"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllCommentNews godoc
// @Summary Get all Comment.
// @Description Get a list of Comment.
// @Tags Comment News
// @Produce json
// @Success 200 {object} []entities.CommentsNews
// @Router /comments-news [get]
// func GetAllCommentNews(c *gin.Context) {
// 	// get db from gin context
// 	db := c.MustGet("app").(*gorm.DB)
// 	var comments []entities.CommentsNews
// 	db.Find(&comments)

// 	c.JSON(http.StatusOK, gin.H{"data": comments})
// }

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
func CreateProfile(svc ProfileService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("sudah masuk handler")
		request, err := contract.ValidateAndBuildProfileInput(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("sudah lewat validate")

		response, err := svc.Create(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Println("sudah lewat service")

		c.JSON(http.StatusOK, gin.H{"data": response})
	}
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
// func DeleteCommentNews(c *gin.Context) {
// 	// Get model if exist
// 	db := c.MustGet("app").(*gorm.DB)
// 	var comment entities.CommentsNews
// 	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	db.Delete(&comment)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }
