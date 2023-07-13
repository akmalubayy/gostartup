package handler

import (
	"gostartup/helper"
	"gostartup/transaction"
	"gostartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transcationHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transcationHandler {
	return &transcationHandler{
		service,
	}
}

func (handler *transcationHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse(
			"Error to get campaign transaction",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := handler.service.GetTranscationByCampaignID(input)

	if err != nil {
		response := helper.APIResponse(
			"Error to get campaign transcation",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Campaign Transaction Details",
		http.StatusOK,
		"success",
		transaction.FormatCampaignTransactions(transactions),
	)

	c.JSON(http.StatusOK, response)
}

func (handler *transcationHandler) GetUserTrancastions(c *gin.Context) {
	// Mengambil data siapa user yang melakukan request
	currentUser := c.MustGet("currentUser").(user.User)

	userID := currentUser.ID

	transactions, err := handler.service.GetTransactionByUserID(userID)

	if err != nil {
		response := helper.APIResponse(
			"Error to get user transactions",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"List of Users Transactions",
		http.StatusOK,
		"success",
		transaction.FormatUserTransactions(transactions),
	)

	c.JSON(http.StatusOK, response)
}

// GetUserTransaction
// handler
// Ambil nilai user dari JWT/Middleware
// service
// Repo -> ambil data transaction (preload data campaign dan campaignImages)
