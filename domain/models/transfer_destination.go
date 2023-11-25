package models

type TransferDestination struct {
	BankCode      string
	Branch        string
	AccountNumber string
	Name          string
	TaxID         string
	AccountType   string
}
