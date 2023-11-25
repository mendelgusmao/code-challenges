package models

type Transfer struct {
	Amount        int    `json:"amount"`
	BankCode      string `json:"bankCode"`
	BranchCode    string `json:"branchCode"`
	AccountNumber string `json:"accountNumber"`
	AccountType   string `json:"accountType"`
	TaxID         string `json:"taxId"`
	Name          string `json:"name"`
}
