package transaction

import "investPedia/user"

type CampaignTransactionInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
