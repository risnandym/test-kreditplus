package handlers

// type phoneInput struct {
// 	Type    string `json:"type"`
// 	Year    int    `json:"year"`
// 	BrandID uint   `json:"brand_id"`
// }

// // GetAllPhones godoc
// // @Summary Get all phones.
// // @Description Get a list of Phones.
// // @Tags Phone
// // @Produce json
// // @Success 200 {object} []entities.Phone
// // @Router /phones [get]
// func GetAllPhone(c *gin.Context) {
// 	// get db from gin context
// 	db := c.MustGet("app").(*gorm.DB)
// 	var phones []entities.Phone
// 	db.Find(&phones)

// 	c.JSON(http.StatusOK, gin.H{"data": phones})
// }

// // CreatePhone godoc
// // @Summary Create New Phone. (Admin only)
// // @Description Creating a new Phone.
// // @Tags Phone
// // @Param Body body phoneInput true "the body to create a new phone"
// // @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// // @Security BearerToken
// // @Produce json
// // @Success 200 {object} entities.Phone
// // @Router /phones [post]
// func CreatePhone(c *gin.Context) {
// 	db := c.MustGet("app").(*gorm.DB)

// 	// Validate input
// 	var input phoneInput
// 	var rating entities.Brand
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.Where("id = ?", input.BrandID).First(&rating).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "BrandID not found!"})
// 		return
// 	}

// 	// Get user active info
// 	user, _ := c.Get("user")
// 	users := user.(entities.Auth)

// 	// Create Phone
// 	phone := entities.Phone{UserID: users.ID, EditorName: users.Email, Type: input.Type, Year: input.Year, BrandID: input.BrandID}
// 	db.Create(&phone)

// 	c.JSON(http.StatusOK, gin.H{"data": phone})
// }

// // GetPhoneById godoc
// // @Summary Get Phone.
// // @Description Get a Phone by id.
// // @Tags Phone
// // @Produce json
// // @Param id path string true "phone id"
// // @Success 200 {object} entities.Phone
// // @Router /phones/{id} [get]
// func GetPhoneById(c *gin.Context) { // Get model if exist
// 	var phone entities.Phone

// 	db := c.MustGet("app").(*gorm.DB)
// 	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": phone})
// }

// // UpdatePhone godoc
// // @Summary Update Phone. (Admin only)
// // @Description Update phone by id.
// // @Tags Phone
// // @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// // @Security BearerToken
// // @Produce json
// // @Param id path string true "phone id"
// // @Param Body body phoneInput true "the body to update an phone"
// // @Success 200 {object} entities.Phone
// // @Router /phones/{id} [patch]
// func UpdatePhone(c *gin.Context) {

// 	db := c.MustGet("app").(*gorm.DB)
// 	// Get model if exist
// 	var phone entities.Phone
// 	var rating entities.Brand
// 	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	// Validate input
// 	var input phoneInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.Where("id = ?", input.BrandID).First(&rating).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "BrandID not found!"})
// 		return
// 	}

// 	// Get user active info
// 	user, _ := c.Get("user")
// 	users := user.(entities.Auth)

// 	var updatedInput entities.Phone
// 	updatedInput.Type = input.Type
// 	updatedInput.Year = input.Year
// 	updatedInput.UserID = users.ID
// 	updatedInput.BrandID = input.BrandID
// 	updatedInput.UpdatedAt = time.Now()

// 	db.Model(&phone).Updates(updatedInput)

// 	c.JSON(http.StatusOK, gin.H{"data": phone})
// }

// // DeletePhone godoc
// // @Summary Delete one phone. (Admin only)
// // @Description Delete a phone by id.
// // @Tags Phone
// // @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// // @Security BearerToken
// // @Produce json
// // @Param id path string true "phone id"
// // @Success 200 {object} map[string]boolean
// // @Router /phones/{id} [delete]
// func DeletePhone(c *gin.Context) {
// 	// Get model if exist
// 	db := c.MustGet("app").(*gorm.DB)
// 	var phone entities.Phone
// 	if err := db.Where("id = ?", c.Param("id")).First(&phone).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	db.Delete(&phone)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }

// // var detail map[string]interface{}

// // GetSpecCommentByPhoneId godoc
// // @Summary Get Specs & CommentsPhone.
// // @Description Get all Specs and Comment by PhoneId.
// // @Tags Phone
// // @Produce json
// // @Param id path string true "Phone id"
// // @Success 200
// // @Router /phones/{id}/specs-comments [get]
// func GetSpecCommentByPhoneId(c *gin.Context) { // Get model if exist

// 	var spec []entities.Spec

// 	db := c.MustGet("app").(*gorm.DB)

// 	if err := db.Where("phone_id = ?", c.Param("id")).Find(&spec).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	var comments []entities.CommentsPhone

// 	dbb := c.MustGet("app").(*gorm.DB)

// 	if err := dbb.Where("phone_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"spec": spec, "comments": comments})
// }
