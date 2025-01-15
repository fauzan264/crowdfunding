package campaign

import (
	"strings"

	"github.com/google/uuid"
)

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

type CampaignDetailFormatter struct {
	ID 					uuid.UUID `json:"id"`
	Title 				string `json:"title"`
	ShortDescription 	string `json:"short_description"`
	ImageURL 			string `json:"image_url"`
	GoalAmount 			int `json:"goal_amount"`
	CurrentAmount 		int `json:"current_amount"`
	UserID 				uuid.UUID `json:"user_id"`
	Slug				string `json:"slug"`
	Description 		string `json:"description"`
	Perks				[]string `json:"perks"`
	User				CampaignUserFormatter `json:"user"`
	Images				[]CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name		string `json:"name"`
	ImageURL 	string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL 	string `json:"image_url"`
	IsPrimary	bool `json:"is_primary"`
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	formatter := CampaignDetailFormatter{
		ID: campaign.ID,
		Title: campaign.Title,
		ShortDescription: campaign.ShortDescription,
		ImageURL: "",
		GoalAmount: campaign.GoalAmount,
		CurrentAmount: campaign.CurrentAmount,
		UserID: campaign.UserID,
		Slug: campaign.Slug,
		Description: campaign.Description,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	formatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName
	formatter.User = campaignUserFormatter

	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName
		campaignImageFormatter.IsPrimary = image.IsPrimary

		images = append(images, campaignImageFormatter)
	}
	formatter.Images = images

	return formatter
}