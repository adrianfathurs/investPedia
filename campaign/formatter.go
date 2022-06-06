package campaign

import (
	"strings"
)

type CampaignFormatter struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Subtitle      string `json:"sub_title"`
	ImageUrl      string `json:"image_url"`
	TargetInvest  int    `json:"target_invest"`
	CurrentInvest int    `json:"current_invest"`
	Slug          string `json:"slug"`
	UserId        int    `json:"user_id"`
}

type CampaignDetailFormatter struct {
	ID             int                       `json:"id"`
	Title          string                    `json:"title"`
	Subtitle       string                    `json:"sub_title"`
	Description    string                    `json:"description"`
	ImageUrl       string                    `json:"image_url"`
	TargetInvest   int                       `json:"target_invest"`
	CurrentInvest  int                       `json:"current_invest"`
	UserId         int                       `json:"user_id"`
	Slug           string                    `json:"slug"`
	Perks          []string                  `json:"perks"`
	User           CampaignUSerFormatter     `json:"user"`
	CampaignImages []CampaignImagesFormatter `json:"campaign_images"`
}

type CampaignUSerFormatter struct {
	FullName string `json:"full_name"`
	Avatar   string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	FileName  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:            campaign.ID,
		Title:         campaign.Title,
		Subtitle:      campaign.SubTitle,
		ImageUrl:      "",
		TargetInvest:  campaign.TargetInvest,
		CurrentInvest: campaign.CurrentInvest,
		Slug:          campaign.Slug,
		UserId:        campaign.UserId,
	}
	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaignImages(campaign CampaignImages) CampaignImagesFormatter {
	var formatter CampaignImagesFormatter
	formatter.FileName = campaign.FileName
	isPrimary := false
	if campaign.IsPrimary == 1 {
		isPrimary = true
	}
	formatter.IsPrimary = isPrimary

	return formatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	perks := []string{}
	formatter := CampaignDetailFormatter{
		ID:            campaign.ID,
		Title:         campaign.Title,
		Subtitle:      campaign.SubTitle,
		ImageUrl:      "",
		TargetInvest:  campaign.TargetInvest,
		CurrentInvest: campaign.CurrentInvest,
		Slug:          campaign.Slug,
		UserId:        campaign.UserId,
		Description:   campaign.Description,
		Perks:         perks,
	}

	//Perks
	for _, item := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(item))
	}
	formatter.Perks = perks

	//campaignImages
	var campaignImages []CampaignImagesFormatter
	if len(campaign.CampaignImages) > 0 {
		for _, item := range campaign.CampaignImages {
			formatterImage := FormatCampaignImages(item)
			campaignImages = append(campaignImages, formatterImage)
		}
		formatter.CampaignImages = campaignImages
	}

	//campaignUser
	CampaignUSerFormatter := CampaignUSerFormatter{
		FullName: campaign.User.FullName,
		Avatar:   campaign.User.Avatar,
	}
	formatter.User = CampaignUSerFormatter

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{}

	for _, item := range campaigns {
		campaignFormatter := FormatCampaign(item)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
