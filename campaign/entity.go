package campaign

import (
	"gostartup/user"
	"time"
)

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	updatedAt  time.Time
}

// type CampaignImage struct {
// 	ID         int
// 	UserID     user.User
// 	CampaignID int
// 	FileName   string
// 	IsPrimary  int
// 	CreatedAt  time.Time
// 	updatedAt  time.Time
// }
