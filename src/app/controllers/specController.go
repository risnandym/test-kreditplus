package controllers

import (
	"net/http"
	"time"

	"final-project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type specInput struct {
	PhoneID   uint   `json:"phone_id"`
	Processor string `json:"processor"`
	Memory    string `json:"memory"`
	Storage   string `json:"storage"`
	Screen    string `json:"screen"`
	Camera    string `json:"camera"`
	Price     string `json:"price"`
	Review    string `json:"review"`
}

// GetAllSpec godoc
// @Summary Get all Spec.
// @Description Get a list of Spec.
// @Tags Spec
// @Produce json
// @Success 200 {object} []models.Spec
// @Router /specs [get]
func GetAllSpec(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var specs []models.Spec
	db.Find(&specs)

	c.JSON(http.StatusOK, gin.H{"data": specs})
}

// CreateSpec godoc
// @Summary Create New Spec. (Admin only)
// @Description Creating a new Spec.
// @Tags Spec
// @Param Body body specInput true "the body to create a new Spec"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Spec
// @Router /specs [post]
func CreateSpec(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input specInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var rating models.Phone
	if err := db.Where("id = ?", input.PhoneID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhoneID not found!"})
		return
	}

	user, _ := c.Get("user")
	users := user.(models.User)

	// Create Spec
	spec := models.Spec{
		PhoneID:   input.PhoneID,
		Processor: input.Processor,
		Memory:    input.Memory,
		Storage:   input.Storage,
		Screen:    input.Screen,
		Camera:    input.Camera,
		Price:     input.Price,
		Review:    input.Review,
		UserID:    users.ID,
	}

	db.Create(&spec)

	c.JSON(http.StatusOK, gin.H{"data": spec})
}

// GetSpecById godoc
// @Summary Get Spec.
// @Description Get an Spec by id.
// @Tags Spec
// @Produce json
// @Param id path string true "Spec id"
// @Success 200 {object} models.Spec
// @Router /specs/{id} [get]
func GetSpecById(c *gin.Context) { // Get model if exist
	var spec models.Spec

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&spec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": spec})
}

// UpdateSpec godoc
// @Summary Update Spec. (Admin only)
// @Description Update Spec by id.
// @Tags Spec
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Spec id"
// @Param Body body specInput true "the body to update age spec category"
// @Success 200 {object} models.Spec
// @Router /specs/{id} [patch]
func UpdateSpec(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var spec models.Spec
	if err := db.Where("id = ?", c.Param("id")).First(&spec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input specInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var rating models.Phone
	if err := db.Where("id = ?", input.PhoneID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PhoneID not found!"})
		return
	}

	var updatedInput models.Spec
	updatedInput.PhoneID = input.PhoneID
	updatedInput.Processor = input.Processor
	updatedInput.Memory = input.Memory
	updatedInput.Storage = input.Storage
	updatedInput.Screen = input.Screen
	updatedInput.Camera = input.Camera
	updatedInput.Price = input.Price
	updatedInput.Review = input.Review
	updatedInput.UpdatedAt = time.Now()

	db.Model(&spec).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": spec})
}

// DeleteSpec godoc
// @Summary Delete one Spec. (Admin only)
// @Description Delete a Spec by id.
// @Tags Spec
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Spec id"
// @Success 200 {object} map[string]boolean
// @Router /specs/{id} [delete]
func DeleteSpec(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var spec models.Spec
	if err := db.Where("id = ?", c.Param("id")).First(&spec).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&spec)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
