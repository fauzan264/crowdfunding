package transaction

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetCampaignByID(campaignID uuid.UUID) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignByID(campaignID uuid.UUID) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("created_at DESC").Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}