package transaction

import (
	"time"

	"github.com/google/uuid"
)

type CampaignTransactionFormatter struct {
	ID				uuid.UUID	`json:"id"`
	Name			string		`json:"name"`
	Amount			int			`json:"amount"`
	CreatedAt		time.Time	`json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	
	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	transactionsFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		transactionFormatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID 			uuid.UUID 			`json:"id"`
	Amount 		int 				`json:"amount"`
	Status 		string 				`json:"status"`
	CreatedAt 	time.Time 			`json:"created_at"`
	Campaign	CampaignFormatter	`json:"campaign"`
}

type CampaignFormatter struct {
	Title		string 		`json:"title"`
	ImageURL 	string 		`json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.Title = transaction.Campaign.Title
	CampaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		CampaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}
	formatter.Campaign = CampaignFormatter
	
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	transactionsFormatter := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		transactionFormatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, transactionFormatter)
	}

	return transactionsFormatter
}