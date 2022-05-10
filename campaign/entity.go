package campaign

import "time"

type Campaign struct {
	ID             int
	Title          string
	SubTitle       string
	Description    string
	TargetInvest   int
	CurrentInvest  int
	Perks          string
	Slug           string
	BackerCount    int
	UserId         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CampaignImages []CampaignImages
}

type CampaignImages struct {
	ID         int
	FileName   string
	IsPrimary  int
	CampaignID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
