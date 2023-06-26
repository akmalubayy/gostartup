package handler

import (
	"fmt"
	"gostartup/campaign"
	"gostartup/helper"
	"gostartup/user"
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

func (handler *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Opss, Something Wrong, Failed Create Campaign",
			http.StatusBadRequest,
			"error",
			errorMessage,
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := handler.service.CreateCampaign(input)

	if err != nil {

		response := helper.APIResponse(
			"Opss, Something Wrong, Failed Create Campaign",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Create Campaign Succesfuly",
		http.StatusOK,
		"success",
		campaign.FormatCampaign(newCampaign),
	)

	c.JSON(http.StatusOK, response)

}

func (handler *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Opss, Something Wrong, Failed Update Campaign",
			http.StatusBadRequest,
			"error",
			errorMessage,
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse(
			"Opss, Something Wrong, Failed Update Campaign Image",
			http.StatusBadRequest,
			"error",
			errorMessage,
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateCampaign, err := handler.service.UpdateCampaign(inputID, input)

	if err != nil {

		response := helper.APIResponse(
			"Opss, Something Wrong, Failed Update Campaign",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Update Campaign Succesfuly",
		http.StatusOK,
		"success",
		campaign.FormatCampaign(updateCampaign),
	)

	c.JSON(http.StatusOK, response)

}

func (handler *campaignHandler) UploadImageCampaign(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest,
			"error",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign image",
			http.StatusBadRequest,
			"error",
			data,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	input.User = currentUser

	path := fmt.Sprintf("images/campaign-images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload avatar image",
			http.StatusBadRequest,
			"error",
			data,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = handler.service.SaveCampaignImage(input, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse(
			"Failed to upload campaign images",
			http.StatusBadRequest,
			"error",
			data,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse(
		"Success upload image campaign",
		http.StatusOK,
		"success",
		data,
	)

	c.JSON(http.StatusOK, response)

}
