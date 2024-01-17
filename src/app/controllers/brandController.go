package controllers

import (
	"net/http"
	"time"

	"final-project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type brandInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetAllBrand godoc
// @Summary Get all Brand.
// @Description Get a list of Brand.
// @Tags Brand
// @Produce json
// @Success 200 {object} []models.Brand
// @Router /brands [get]
func GetAllBrand(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var merks []models.Brand
	db.Find(&merks)

	c.JSON(http.StatusOK, gin.H{"data": merks})
}

// CreateBrand godoc
// @Summary Create New Brand. (Admin only)
// @Description Creating a new Brand.
// @Tags Brand
// @Param Body body brandInput true "the body to create a new Brand"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Brand
// @Router /brands [post]
func CreateBrand(c *gin.Context) {
	// Validate input
	var input brandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user active info
	user, _ := c.Get("user")
	users := user.(models.User)

	// Create Brand
	merk := models.Brand{UserID: users.ID, Name: input.Name, Description: input.Description}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&merk)

	c.JSON(http.StatusOK, gin.H{"data": merk})
}

// GetBrandById godoc
// @Summary Get Brand by Id.
// @Description Get an Brand by Id.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} models.Brand
// @Router /brands/{id} [get]
func GetBrandById(c *gin.Context) { // Get model if exist
	var merk models.Brand

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&merk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merk})
}

// GetPhoneByBrandId godoc
// @Summary Get Phones by Brand Id.
// @Description Get all Phones by Brand Id.
// @Tags Brand
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} []models.Phone
// @Router /brands/{id}/phones [get]
func GetPhonesByBrandId(c *gin.Context) { // Get model if exist
	var phones []models.Phone

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("brand_id = ?", c.Param("id")).Find(&phones).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": phones})
}

// UpdateBrand godoc
// @Summary Update Brand. (Admin only)
// @Description Update Brand by id.
// @Tags Brand
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Brand id"
// @Param Body body brandInput true "the body to update age merk category"
// @Success 200 {object} models.Brand
// @Router /brands/{id} [patch]
func UpdateBrand(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var merk models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&merk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input brandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user active
	user, _ := c.Get("user")
	users := user.(models.User)

	var updatedInput models.Brand
	updatedInput.Name = input.Name
	updatedInput.UserID = users.ID
	updatedInput.Description = input.Description
	updatedInput.UpdatedAt = time.Now()

	db.Model(&merk).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": merk})
}

// DeleteBrand godoc
// @Summary Delete one Brand. (Admin only)
// @Description Delete a Brand by id.
// @Tags Brand
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Brand id"
// @Success 200 {object} map[string]boolean
// @Router /brands/{id} [delete]
func DeleteBrand(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var merk models.Brand
	if err := db.Where("id = ?", c.Param("id")).First(&merk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&merk)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
