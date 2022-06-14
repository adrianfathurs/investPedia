package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input CampaignInput) (Campaign, error)
	UpdateCampaign(inputID CampaignInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImages, error)
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
	campaign.Description = input.Description
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

func (s *service) UpdateCampaign(inputID CampaignInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindCampaignByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.ID {
		return campaign, errors.New("you dont have access to update this campaign")
	}

	campaign.Title = inputData.Title
	campaign.SubTitle = inputData.SubTitle
	campaign.Description = inputData.Description
	campaign.TargetInvest = inputData.TargetInvest
	campaign.Perks = inputData.Perks
	campaign.User = inputData.User

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImages, error) {
	campaign, err := s.repository.FindCampaignByID(input.CampaignID)
	if err != nil {
		return CampaignImages{}, err
	}

	if campaign.UserId != input.User.ID {
		return CampaignImages{}, errors.New("you dont have access to upload image on this campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImages{}, err
		}
	}
	campaignImage := CampaignImages{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.FileName = fileLocation
	campaignImage.IsPrimary = isPrimary
	newCampaignImage, err := s.repository.SaveCampaignImage(campaignImage)

	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
