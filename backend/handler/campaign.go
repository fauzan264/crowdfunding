package handler

import (
	"net/http"

	"github.com/fauzan264/crowdfunding/backend/campaign"
	"github.com/fauzan264/crowdfunding/backend/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	getUserID := c.Query("user_id")
	
	if getUserID == "" {
		getUserID = uuid.Nil.String()
	}

	userID, err := uuid.Parse(getUserID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	formatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}