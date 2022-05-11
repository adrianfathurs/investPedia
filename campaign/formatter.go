package campaign

type CampaignFormatter struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Subtitle      string `json:"sub_title"`
	ImageUrl      string `json:"image_url"`
	TargetInvest  int    `json:"target_invest"`
	CurrentInvest int    `json:"current_invest"`
	UserId        int    `json:"user_id"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:            campaign.ID,
		Title:         campaign.Title,
		Subtitle:      campaign.SubTitle,
		ImageUrl:      "",
		TargetInvest:  campaign.TargetInvest,
		CurrentInvest: campaign.CurrentInvest,
		UserId:        campaign.UserId,
	}
	if len(campaign.CampaignImages) > 0 {
		formatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

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
