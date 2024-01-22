package handlers

import (
	"net/http"
	"test-kreditplus/src/contract"

	"github.com/gin-gonic/gin"
)

// CreateCommentNews godoc
//
//	@Summary		Create Profile.
//	@Description	Save Customer Profile.
//	@Tags			Customer
//	@Param			Body	body	contract.ProfileInput	true	"the body to create a new Profile"
//	@Security		kreditplus-token
//	@Produce		json
//	@Success		200	{object}	contract.ProfileOutput
//	@Router			/kredit-plus/customer/profile [post]
func CreateProfile(svc ProfileService) gin.HandlerFunc {
	return func(c *gin.Context) {

		request, err := contract.ValidateAndBuildProfileInput(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := svc.Create(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": response})
	}
}
