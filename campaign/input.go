package campaign

import "investPedia/user"

type CampaignInput struct {
	ID int `uri:"id" binding:"required`
}

type CreateCampaignInput struct {
	Title        string `json:"title" binding:"required"`
	SubTitle     string `json:"sub_title" binding:"required"`
	Description  string `json:"description" binding:"required"`
	TargetInvest int    `json:"target_invest" binding:"required"`
	Perks        string `json:"perks" binding:"required"`
	User         user.User
}

type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}
