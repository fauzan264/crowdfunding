package campaign

import "github.com/fauzan264/crowdfunding/backend/user"

type GetCampaignDetailInput struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type CreateCampaignInput struct {
	Title				string `json:"title" binding:"required"`
	ShortDescription 	string `json:"short_description" binding:"required"`
	Description 		string `json:"description" binding:"required"`
	GoalAmount 			int `json:"goal_amount" binding:"required"`
	Perks 				string `json:"perks" binding:"required"`
	User				user.User
}