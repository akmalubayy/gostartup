package transaction

import (
	"errors"
	"gostartup/campaign"
)

type Service interface {
	GetTranscationByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	GetAllTransactions() ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{
		repository,
		campaignRepository,
	}
}

func (s *service) GetTranscationByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of campaign")
	}

	transaction, err := s.repository.GetByCampaignID(campaign.ID)

	if err != nil {
		return transaction, err
	}

	return transaction, nil

}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetAllTransactions() ([]Transaction, error) {
	transactions, err := s.repository.FindAll()

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
