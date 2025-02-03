package transaction

import (
	"errors"

	"github.com/fauzan264/crowdfunding/backend/campaign"
	"github.com/google/uuid"
)

type service struct {
	repository Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{
		repository,
		campaignRepository,
	}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaignID, err := uuid.Parse(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	getCampaign, err := s.campaignRepository.FindByID(campaignID)
	if err != nil {
		return []Transaction{}, err
	}

	if getCampaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := s.repository.GetCampaignByID(campaignID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}