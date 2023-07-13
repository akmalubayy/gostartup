package transaction

import (
	"time"
)

type TransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatCampaignTransaction(transaction Transaction) TransactionFormatter {
	campaignTransactionFormatter := TransactionFormatter{}

	campaignTransactionFormatter.ID = transaction.ID
	campaignTransactionFormatter.Name = transaction.User.Name
	campaignTransactionFormatter.Amount = transaction.Amount
	campaignTransactionFormatter.CreatedAt = transaction.CreatedAt

	return campaignTransactionFormatter

}

func FormatCampaignTransactions(transactions []Transaction) []TransactionFormatter {
	if len(transactions) == 0 {
		return []TransactionFormatter{}
	}

	var transactionFormatter []TransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	userTransactionFormatter := UserTransactionFormatter{}

	userTransactionFormatter.ID = transaction.ID
	userTransactionFormatter.Amount = transaction.Amount
	userTransactionFormatter.Status = transaction.Status
	userTransactionFormatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}

	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	userTransactionFormatter.Campaign = campaignFormatter

	return userTransactionFormatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var userTransactionFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		userTransactionFormatter = append(userTransactionFormatter, formatter)
	}

	return userTransactionFormatter
}
