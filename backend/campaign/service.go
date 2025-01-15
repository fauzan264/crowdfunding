package campaign

import (
	"errors"

	"github.com/google/uuid"
)

type Service interface {
	GetCampaigns(userID uuid.UUID) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID uuid.UUID) ([]Campaign, error) {
	if userID != uuid.Nil {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, errors.New("user ID not found or UUID not valid")
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	
	campaign, err := s.repository.FindByID(uuid.MustParse(input.ID))

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}