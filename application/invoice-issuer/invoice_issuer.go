package invoiceissuer

import (
	"math/rand"

	"github.com/mendelgusmao/tony/application/invoice-issuer/config"
	"github.com/mendelgusmao/tony/domain/models"
	"github.com/mendelgusmao/tony/domain/ports"
)

type InvoiceIssuer struct {
	config         config.Specification
	invoiceService ports.InvoiceService
}

func NewInvoiceIssuer(config config.Specification, invoiceService ports.InvoiceService) *InvoiceIssuer {
	return &InvoiceIssuer{
		config:         config,
		invoiceService: invoiceService,
	}
}

func (i InvoiceIssuer) IssueInvoices() ([]models.Invoice, []error) {
	invoices := i.generateInvoices()
	return i.invoiceService.CreateInvoices(invoices)
}

func (i InvoiceIssuer) generateInvoices() []models.Invoice {
	amount := i.config.MinInvoicesAmount + rand.Intn(i.config.MaxInvoicesAmount-i.config.MinInvoicesAmount)
	invoices := make([]models.Invoice, amount)
	index := 0

	for index < amount {
		invoices[index] = models.NewFakeInvoice().ToInvoice()
		index++
	}

	return invoices
}
