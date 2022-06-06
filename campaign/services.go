package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input CampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}

	campaign.Title = input.Title
	campaign.SubTitle = input.SubTitle
	campaign.Description = input.Desctiption
	campaign.TargetInvest = input.TargetInvest
	campaign.Perks = input.Perks
	campaign.UserId = input.User.ID
	slugCandidate := fmt.Sprintf("%s %d", input.Title, input.User.ID)
	//pembuatan slug
	campaign.Slug = slug.Make(slugCandidate)

	campaigm, err := s.repository.Save(campaign)
	if err != nil {
		return campaign, err
	}
	return campaigm, nil
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindCampaignByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAllCampaign()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input CampaignInput) (Campaign, error) {
	campaigns, err := s.repository.FindCampaignByID(input.ID)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}
