package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	Name string
	Occupation string
	Email string
	Password string
	AvatarFileName string
	Role string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedBy uuid.UUID
	UpdatedAt time.Time
}