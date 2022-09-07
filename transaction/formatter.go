package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type CampaignTitleandImages struct {
	Title    string `json:"name"`
	FileName string `json:"image_url"`
}
type UserTransactionsFormatter struct {
	ID                int                    `json:"id"`
	Amount            int                    `json:"amount"`
	TransactionStatus string                 `json:"status"`
	CreatedAt         time.Time              `json:"created_at"`
	Campaign          CampaignTitleandImages `json:"campaign"`
}

type TransactionsFormatter struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	Amount            int    `json:"amount"`
	Code              string `json:"code"`
	TransactionStatus string `json:"status"`
	PaymentURL        string `json:"payment_url"`
	CampaignID        int    `json:"campaign_id"`
}

func FormatTransaction(transaction Transaction) TransactionsFormatter {
	formatter := TransactionsFormatter{
		ID:                transaction.ID,
		UserID:            transaction.UserId,
		CampaignID:        transaction.CampaignId,
		TransactionStatus: transaction.TransactionStatus,
		Code:              transaction.Code,
		Amount:            transaction.Amount,
		PaymentURL:        transaction.PaymentURL,
	}
	return formatter
}
func FormatCampaignTransaction(campaignTransaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{
		ID:        campaignTransaction.ID,
		FullName:  campaignTransaction.User.FullName,
		Amount:    campaignTransaction.Amount,
		CreatedAt: campaignTransaction.CreatedAt,
	}
	return formatter
}

func FormatListCampaignTransaction(campaignTransaction []Transaction) []CampaignTransactionFormatter {
	if len(campaignTransaction) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var formatter []CampaignTransactionFormatter

	for _, item := range campaignTransaction {
		formatter = append(formatter, FormatCampaignTransaction(item))
	}
	return formatter
}

func FormatListUserTransaction(transactions []Transaction) []UserTransactionsFormatter {
	if len(transactions) == 0 {
		return []UserTransactionsFormatter{}
	}

	var formatter []UserTransactionsFormatter

	for _, item := range transactions {
		formatter = append(formatter, FormatUserTransaction(item))
	}
	return formatter
}

func FormatUserTransaction(transaction Transaction) UserTransactionsFormatter {
	campaign := CampaignTitleandImages{}
	campaign.Title = transaction.Campaign.Title
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaign.FileName = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter := UserTransactionsFormatter{
		ID:                transaction.ID,
		Amount:            transaction.Amount,
		TransactionStatus: transaction.TransactionStatus,
		CreatedAt:         transaction.CreatedAt,
		Campaign:          campaign,
	}
	return formatter
}
