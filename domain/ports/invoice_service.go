package ports

import "github.com/mendelgusmao/tony/domain/models"

type InvoiceService interface {
	CreateInvoices([]models.Invoice) ([]models.Invoice, []error)
}
