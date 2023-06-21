package handler

import (
	"gostartup/campaign"
	"gostartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Tangkap parameter di handler
// Handler ke service
// service menentukan repository mana yang di-hit
// repsitory : getAll dan GetByUserID
// akses db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{
		service,
	}
}

// routing ke api/v1/campaigns
func (handler *campaignHandler) GetCampaigns(c *gin.Context) {
	// convert string to number(int)
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := handler.service.GetCampaigns(userID)

	if err != nil {
		response := helper.APIResponse(
			"Error to get campaigns",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// formatter := campaign.FormatCampaign(campaign.Campaign{})

	response := helper.APIResponse(
		"List of campaigns",
		http.StatusOK,
		"success",
		campaign.FormatCampaigns(campaigns),
	)

	c.JSON(http.StatusOK, response)
}

func (handler *campaignHandler) GetCampaign(c *gin.Context) {

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse(
			"Error to get details campaign",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := handler.service.GetCampaignByID(input)

	if err != nil {
		response := helper.APIResponse(
			"Error to get details campaign",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Succes get campaign Details",
		http.StatusOK,
		"success",
		campaign.FormatCampaignDetail(campaignDetail),
	)

	c.JSON(http.StatusOK, response)
}
