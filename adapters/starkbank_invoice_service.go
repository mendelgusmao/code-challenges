package adapters

import (
	"fmt"

	"github.com/mendelgusmao/tony/domain/models"
	starkbankInvoice "github.com/starkbank/sdk-go/starkbank/invoice"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

type StarkBankInvoiceService struct {
	project project.Project
}

func NewStarkBankInvoiceService(project project.Project) *StarkBankInvoiceService {
	return &StarkBankInvoiceService{
		project: project,
	}
}

func (s StarkBankInvoiceService) CreateInvoices(invoices []models.Invoice) ([]models.Invoice, []error) {
	starkInvoices := make([]starkbankInvoice.Invoice, len(invoices))

	for index, invoice := range invoices {
		starkInvoices[index] = starkbankInvoice.Invoice{
			Id:             invoice.Id,
			Amount:         invoice.Amount,
			Name:           invoice.Name,
			TaxId:          invoice.TaxId,
			Due:            invoice.Due,
			Expiration:     invoice.Expiration,
			Fine:           invoice.Fine,
			Interest:       invoice.Interest,
			Discounts:      invoice.Discounts,
			Tags:           invoice.Tags,
			Rules:          invoice.Rules,
			Descriptions:   invoice.Descriptions,
			Pdf:            invoice.Pdf,
			Link:           invoice.Link,
			NominalAmount:  invoice.NominalAmount,
			FineAmount:     invoice.FineAmount,
			InterestAmount: invoice.InterestAmount,
			DiscountAmount: invoice.DiscountAmount,
			Brcode:         invoice.Brcode,
			Status:         invoice.Status,
			Fee:            invoice.Fee,
			TransactionIds: invoice.TransactionIds,
			Created:        invoice.Created,
			Updated:        invoice.Updated,
		}
	}

	createdInvoices, err := starkbankInvoice.Create(starkInvoices, s.project)

	if err.Errors != nil {
		errors := make([]error, len(err.Errors))

		for index, e := range err.Errors {
			errors[index] = fmt.Errorf("invoice error. code: %s, message: %s", e.Code, e.Message)
		}

		return nil, errors
	}

	domainInvoices := make([]models.Invoice, len(createdInvoices))

	for index, starkInvoice := range createdInvoices {
		domainInvoices[index] = models.Invoice{
			Id:             starkInvoice.Id,
			Amount:         starkInvoice.Amount,
			Name:           starkInvoice.Name,
			TaxId:          starkInvoice.TaxId,
			Due:            starkInvoice.Due,
			Expiration:     starkInvoice.Expiration,
			Fine:           starkInvoice.Fine,
			Interest:       starkInvoice.Interest,
			Discounts:      starkInvoice.Discounts,
			Tags:           starkInvoice.Tags,
			Rules:          starkInvoice.Rules,
			Descriptions:   starkInvoice.Descriptions,
			Pdf:            starkInvoice.Pdf,
			Link:           starkInvoice.Link,
			NominalAmount:  starkInvoice.NominalAmount,
			FineAmount:     starkInvoice.FineAmount,
			InterestAmount: starkInvoice.InterestAmount,
			DiscountAmount: starkInvoice.DiscountAmount,
			Brcode:         starkInvoice.Brcode,
			Status:         starkInvoice.Status,
			Fee:            starkInvoice.Fee,
			TransactionIds: starkInvoice.TransactionIds,
			Created:        starkInvoice.Created,
			Updated:        starkInvoice.Updated,
		}
	}

	return domainInvoices, nil
}
