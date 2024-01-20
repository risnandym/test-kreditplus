package handlers

import (
	"net/http"
	"test-kreditplus/src/contract"

	"github.com/gin-gonic/gin"
)

// type newsInput struct {
// 	Title    string `json:"title"`
// 	Content  string `json:"content"`
// 	Link_URL string `json:"link_url"`
// }

// // GetAllNews godoc
// // @Summary Get all News.
// // @Description Get a list of News.
// // @Tags News
// // @Produce json
// // @Success 200 {object} []entities.News
// // @Router /news [get]
// func GetAllNews(c *gin.Context) {
// 	// get db from gin context
// 	db := c.MustGet("app").(*gorm.DB)
// 	var news []entities.News
// 	db.Find(&news)

// 	c.JSON(http.StatusOK, gin.H{"data": news})
// }

// CreateTransaction godoc
//
//	@Summary		Create New Transaction. (Admin only)
//	@Description	Creating a new Transaction.
//	@Tags			Transaction
//	@Param			Body			body	contract.TransactionInput	true	"the body to create a new transaction"
//	@Security		kreditplus-token
//	@Produce		json
//	@Success		200	{object}	contract.TransactionOutput
//	@Router			/news [post]
func CreditTransaction(svc TransactionService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.ValidateAndBuildCreditInput(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Credit(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}

// // GetNewsById godoc
// // @Summary Get News by Id.
// // @Description Get a News by id.
// // @Tags News
// // @Produce json
// // @Param id path string true "phone id"
// // @Success 200 {object} entities.News
// // @Router /news/{id} [get]
// func GetNewsById(c *gin.Context) { // Get model if exist
// 	var phone entities.News

// 	db := c.MustGet("app").(*gorm.DB)
// 	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": phone})
// }

// // UpdateNews godoc
// // @Summary Update News. (Admin only)
// // @Description Update phone by id.
// // @Tags News
// // @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// // @Security BearerToken
// // @Produce json
// // @Param id path string true "phone id"
// // @Param Body body newsInput true "the body to update an phone"
// // @Success 200 {object} entities.News
// // @Router /news/{id} [patch]
// func UpdateNews(c *gin.Context) {

// 	db := c.MustGet("app").(*gorm.DB)
// 	// Get model if exist
// 	var phone entities.News
// 	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input newsInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Get user active info
// 	user, _ := c.Get("user")
// 	users := user.(entities.Auth)

// 	var updatedInput entities.News
// 	updatedInput.CreatorName = users.Email
// 	updatedInput.UserID = users.ID
// 	updatedInput.Title = input.Title
// 	updatedInput.Content = input.Content
// 	updatedInput.Link_URL = input.Link_URL
// 	updatedInput.UpdatedAt = time.Now()

// 	db.Model(&phone).Updates(updatedInput)

// 	c.JSON(http.StatusOK, gin.H{"data": phone})
// }

// // DeleteNews godoc
// // @Summary Delete one News. (Admin only)
// // @Description Delete a phone by id.
// // @Tags News
// // @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// // @Security BearerToken
// // @Produce json
// // @Param id path string true "phone id"
// // @Success 200 {object} map[string]boolean
// // @Router /news/{id} [delete]
// func DeleteNews(c *gin.Context) {
// 	// Get model if exist
// 	db := c.MustGet("app").(*gorm.DB)
// 	var phone entities.News
// 	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Get user active info

// 	db.Delete(&phone)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }

// // var detail map[string]interface{}

// // GetSpecCommentByNewsId godoc
// // @Summary Get Comments by NewsId.
// // @Description Get all Specs and Comment by NewsId.
// // @Tags News
// // @Produce json
// // @Param id path string true "News id"
// // @Success 200
// // @Router /news/{id}/comments [get]
// func GetCommentByNewsId(c *gin.Context) { // Get model if exist
// 	var comments []entities.CommentsNews

// 	db := c.MustGet("app").(*gorm.DB)

// 	if err := db.Where("news_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"comments": comments})
// }
