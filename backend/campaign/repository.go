package campaign

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID uuid.UUID) ([]Campaign, error)
	FindByID(ID uuid.UUID) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Order("campaigns.created_at ASC").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID uuid.UUID) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Order("campaigns.updated_at ASC").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID uuid.UUID) (Campaign, error) {
	var campaign Campaign

	result := r.db.Preload("User").
				Preload("CampaignImages").
				Where("id = ?", ID).
				Order("campaigns.created_at ASC").
				Find(&campaign)
	
	if result.RowsAffected == 0 {
		return campaign, errors.New("Data Not Found")
	}

	if result.Error != nil {
		return campaign, result.Error
	}
	
	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	if err := r.db.Create(&campaign).Error; err != nil {
		return campaign, err
	}

	return campaign, nil
}