package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID 				uuid.UUID
	CampaignID 		uuid.UUID
	UserID			uuid.UUID
	Amount			int
	Status			string
	Code			string
	CreatedBy		uuid.UUID
	CreatedAt		time.Time
	UpdatedBy		uuid.UUID
	UpdatedAt		time.Time
}