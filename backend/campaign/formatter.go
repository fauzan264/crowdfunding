package campaign

import "github.com/google/uuid"

type CampaignFormatter struct {
	ID 					uuid.UUID `json:"id"`
	Title				string `json:"title"`
	ShortDescription	string `json:"short_description"`
	ImageURL			string `json:"image_url"`
	GoalAmount			int `json:"goal_amount"`
	CurrentAmount		int `json:"current_amount"`
	Slug				string `json:"slug"`
	UserID				uuid.UUID
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID: campaign.ID,
		Title: campaign.Title,
		ShortDescription: campaign.ShortDescription,
		ImageURL: "",
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		Slug: campaign.Slug,
		UserID: campaign.UserID,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}