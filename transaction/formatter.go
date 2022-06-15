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
