package campaign

import (
	"time"

	"github.com/fauzan264/crowdfunding/backend/user"
	"github.com/google/uuid"
)

type Campaign struct {
	ID 					uuid.UUID
	UserID				uuid.UUID
	Title				string
	ShortDescription 	string
	Description			string
	GoalAmount			int
	CurrentAmount		int
	Perks				string
	BeckerCount			int
	Slug				string
	CreatedBy			uuid.UUID
	CreatedAt			time.Time
	UpdatedBy			uuid.UUID
	UpdatedAt			time.Time
	CampaignImages		[]CampaignImage
	User				user.User
}

type CampaignImage struct {
	ID 				uuid.UUID
	CampaignID 		uuid.UUID
	FileName 		string
	IsPrimary		bool
	CreatedBy 		uuid.UUID
	CreatedAt 		time.Time
	UpdatedBy 		uuid.UUID
	UpdatedAt 		time.Time
}