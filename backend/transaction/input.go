package transaction

import "github.com/fauzan264/crowdfunding/backend/user"


type GetCampaignTransactionsInput struct {
	ID 			string `uri:"id" binding:"required"`
	User		user.User
}