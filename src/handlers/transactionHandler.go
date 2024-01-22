package handlers

import (
	"net/http"
	"test-kreditplus/src/contract"

	"github.com/gin-gonic/gin"
)

// CreditTransaction godoc
//
//	@Summary		Create New Credit Transaction. (Admin only)
//	@Description	Creating a new Credit Transaction.
//	@Tags			Transaction
//	@Param			Body	body	contract.CreditInput	true	"the body to create a new credit transaction"
//	@Security		kreditplus-token
//	@Produce		json
//	@Success		200	{object}	contract.CreditOutput
//	@Router			/kredit-plus/transaction/credit [post]
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

// DebitTransaction godoc
//
//	@Summary		Create New Debit Transaction. (Admin only)
//	@Description	Creating a new Debit Transaction.
//	@Tags			Transaction
//	@Param			Body	body	contract.CreditInput	true	"the body to create a new transaction"
//	@Security		kreditplus-token
//	@Produce		json
//	@Success		200	{object}	contract.CreditOutput
//	@Router			/news [post]
// func DebitTransaction(svc TransactionService) gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		request, err := contract.ValidateAndBuildDebitInput(c)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		response, err := svc.Debit(request)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"data": response})
// 	}
// }
