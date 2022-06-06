package handler

import (
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
		response := helper.APIResponse("Invalid Input", "failed", http.StatusBadRequest, nil)
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
