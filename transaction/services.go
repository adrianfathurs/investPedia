package transaction

import (
	"errors"
	"investPedia/campaign"
	"investPedia/payment"
)

type Service interface {
	GetTransactionsByCampaignID(campaignTransaction CampaignTransactionInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransactionInput(input CreateTransactionInput) (Transaction, error)
}
type service struct {
	transactionRepository Repository
	campaignRepository    campaign.Repository
	paymentService        payment.Services
}

func NewService(transactionRepository Repository, campaignRepository campaign.Repository, paymentService payment.Services) *service {
	return &service{transactionRepository, campaignRepository, paymentService}
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

func (s *service) CreateTransactionInput(input CreateTransactionInput) (Transaction, error) {
	trans := Transaction{}
	trans.Amount = input.Amount
	trans.CampaignId = input.CampaignID
	trans.UserId = input.User.ID
	trans.TransactionStatus = "pending"
	trans.Code = ""
	// create data transaction and column payment_url still null
	newTransaction, err := s.transactionRepository.Save(trans)
	if err != nil {
		return newTransaction, err
	}
	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}
	// we can get payment url from midtrans generate
	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}
	// we data on transaction table and update column payment_url
	newTransaction.PaymentURL = paymentUrl
	newTransaction, err = s.transactionRepository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}
	return newTransaction, nil
}
