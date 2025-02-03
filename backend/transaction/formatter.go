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