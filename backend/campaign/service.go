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
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(uuid.MustParse(inputID.ID))
	
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Title = inputData.Title
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount
	campaign.UpdatedBy = inputData.User.ID
	campaign.UpdatedAt = time.Now()

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaignID, err := uuid.Parse(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	campaign, err := s.repository.FindByID(campaignID)
	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}
	
	if input.IsPrimary {
		_, err := s.repository.MarkAllImagesAsNonPrimary(campaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.ID = uuid.New()
	campaignImage.CampaignID = campaignID
	campaignImage.IsPrimary = input.IsPrimary
	campaignImage.FileName = fileLocation
	campaignImage.CreatedBy = input.User.ID
	campaignImage.CreatedAt = time.Now()

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}