package transaction

import (
	"time"

	"github.com/fauzan264/crowdfunding/backend/user"
	"github.com/google/uuid"
)

type Transaction struct {
	ID 				uuid.UUID
	CampaignID 		uuid.UUID
	UserID			uuid.UUID
	Amount			int
	Status			string
	Code			string
	User			user.User
	CreatedBy		uuid.UUID
	CreatedAt		time.Time
	UpdatedBy		uuid.UUID
	UpdatedAt		time.Time
}