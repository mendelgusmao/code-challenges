package adapters

import (
	"fmt"

	"github.com/mendelgusmao/tony/domain/models"
	starkbankTransfer "github.com/starkbank/sdk-go/starkbank/transfer"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

type StarkBankTransferService struct {
	project project.Project
}

func NewStarkBankTransferService(project project.Project) *StarkBankTransferService {
	return &StarkBankTransferService{
		project: project,
	}
}

func (s StarkBankTransferService) CreateTransfers(transfers []models.Transfer) ([]models.Transfer, []error) {
	starkTransfers := make([]starkbankTransfer.Transfer, len(transfers))

	for index, transfer := range transfers {
		starkTransfers[index] = starkbankTransfer.Transfer{
			Amount:        transfer.Amount,
			BankCode:      transfer.BankCode,
			BranchCode:    transfer.BranchCode,
			AccountNumber: transfer.AccountNumber,
			AccountType:   transfer.AccountType,
			TaxId:         transfer.TaxID,
			Name:          transfer.Name,
		}
	}

	createdTransfers, err := starkbankTransfer.Create(starkTransfers, s.project)

	if err.Errors != nil {
		errors := make([]error, len(err.Errors))

		for index, e := range err.Errors {
			errors[index] = fmt.Errorf("transfer error. code: %s, message: %s", e.Code, e.Message)
		}
	}

	domainTransfers := make([]models.Transfer, len(createdTransfers))

	for index, starkTransfer := range createdTransfers {
		domainTransfers[index] = models.Transfer{
			Amount:        starkTransfer.Amount,
			BankCode:      starkTransfer.BankCode,
			BranchCode:    starkTransfer.BranchCode,
			AccountNumber: starkTransfer.AccountNumber,
			AccountType:   starkTransfer.AccountType,
			TaxID:         starkTransfer.TaxId,
			Name:          starkTransfer.Name,
		}
	}

	return domainTransfers, nil
}
