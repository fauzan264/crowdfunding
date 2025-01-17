package campaign

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID uuid.UUID) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
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

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.ID = uuid.New()
	campaign.Title = input.Title
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	slugCampaign := fmt.Sprintf("%d %s", time.Now().UnixNano() / int64(time.Millisecond), input.Title)
	campaign.Slug = slug.Make(slugCampaign)
	campaign.CreatedBy = input.User.ID
	campaign.CreatedAt = time.Now()

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}