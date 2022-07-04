package transaction

import (
	"investPedia/campaign"
	"investPedia/user"
	"time"
)

type Transaction struct {
	ID                int
	Amount            int
	UserId            int
	CampaignId        int
	TransactionStatus string
	Code              string
	User              user.User
	Campaign          campaign.Campaign
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
