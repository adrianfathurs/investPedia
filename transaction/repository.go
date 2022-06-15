package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetTransactionsByCampaignID(campaignID int) ([]Transaction, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransactionsByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
