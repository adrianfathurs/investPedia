package transaction

import (
	"errors"
	"investPedia/campaign"
)

type Service interface {
	GetTransactionsByCampaignID(campaignTransaction CampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}
type service struct {
	transactionRepository Repository
	campaignRepository    campaign.Repository
}

func NewService(transactionRepository Repository, campaignRepository campaign.Repository) *service {
	return &service{transactionRepository, campaignRepository}
}

func (s *service) GetTransactionsByCampaignID(campaignTransaction CampaignTransactionInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindCampaignByID(campaignTransaction.ID)
	if err != nil {
		return []Transaction{}, err
	}
	if campaignTransaction.User.ID != campaign.UserId {
		return []Transaction{}, errors.New("you dont have to access this transaction data")
	}

	transaction, err := s.transactionRepository.GetTransactionsByCampaignID(campaignTransaction.ID)

	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transaction, err := s.transactionRepository.GetTransactionsByUserID(userID)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
