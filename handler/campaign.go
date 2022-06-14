package handler

import (
	"fmt"
	"investPedia/campaign"
	"investPedia/helper"
	"investPedia/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
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
	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Invalid Input", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to create campaign", "statusOK", http.StatusOK, campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaign", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	responseFormatter := campaign.FormatCampaigns(campaigns)
	responses := helper.APIResponse("Ok", "success", http.StatusOK, responseFormatter)
	c.JSON(http.StatusOK, responses)
	return

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.CampaignInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Invalid Input", "Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Invalid Input", "Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatCampaignDetail := campaign.FormatCampaignDetail(campaignDetail)
	response := helper.APIResponse("OK", "success", http.StatusOK, formatCampaignDetail)
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {

	var inputID campaign.CampaignInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update Campaign, ID is not valid", "Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Invalid Input Form", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Invalid in update data", "Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update campaign", "statusOK", http.StatusOK, campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
	return
}

// CampaignImages
func (h *campaignHandler) UploadCampaignImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Invalid Input Campaign Images", "Failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	file, err := c.FormFile("campaign_image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload images", "Failed", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	filePath := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload images", "Failed", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	input.User = currentUser
	_, err = h.service.SaveCampaignImage(input, filePath)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload images or you dont have access to upload image this campaign", "Failed", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Successfully to upload images", "Success", http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
