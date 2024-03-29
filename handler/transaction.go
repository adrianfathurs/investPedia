package handler

import (
	"investPedia/helper"
	"investPedia/transaction"
	"investPedia/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// pikirkan bahwa id pada uri adalah id campaign, sehingga menggunakan binding uri
// buat formatter untuk response data

type TransactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *TransactionHandler {
	return &TransactionHandler{service}
}

func (h *TransactionHandler) GetTransactionsByCampaignID(c *gin.Context) {
	var input transaction.CampaignTransactionInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Invalid Params", "failed", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get transacations", "Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	transactionsFormatter := transaction.FormatListCampaignTransaction(transactions)
	response := helper.APIResponse("Successfully to get transactions", "Success", http.StatusOK, transactionsFormatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *TransactionHandler) GetTransactionsByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	transactions, err := h.service.GetTransactionsByUserID(currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get transacations", "Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	transactionFormatter := transaction.FormatListUserTransaction(transactions)
	response := helper.APIResponse("Successfully to get transactions", "Success", http.StatusOK, transactionFormatter)
	c.JSON(http.StatusOK, response)
	return
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Invalid Input", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	newTransaction, err := h.service.CreateTransactionInput(input)
	if err != nil {
		response := helper.APIResponse("Failed to post transacations", "Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// newTransaction.Code = "200"
	response := helper.APIResponse("Successfully to post transactions", "Success", http.StatusOK, transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
	return

}
