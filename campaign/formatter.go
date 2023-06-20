package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	// Cara ke-1
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter

	// Cara ke-2
	// formatter := CampaignFormatter{
	// 	ID:               campaign.ID,
	// 	UserID:           campaign.UserID,
	// 	Name:             campaign.Name,
	// 	ShortDescription: campaign.ShortDescription,
	// 	GoalAmount:       campaign.GoalAmount,
	// 	CurrentAmount:    campaign.CurrentAmount,
	// 	Slug:             campaign.Slug,
	// 	ImageURL:         "",
	// }

	// if len(campaign.CampaignImages) > 0 {
	// 	ImageURL := campaign.CampaignImages[0].FileName

	// 	formatter = append(formatter, ImageURL)
	// }

	// return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	// if len(campaigns) == 0 {
	// 	return []CampaignFormatter{}
	// }

	CampaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		CampaignFormatter := FormatCampaign(campaign)
		CampaignsFormatter = append(CampaignsFormatter, CampaignFormatter)
	}

	return CampaignsFormatter
}
