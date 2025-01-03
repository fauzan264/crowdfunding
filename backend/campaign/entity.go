package campaign

import (
	"time"

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
	UpdatedAt			uuid.UUID
}