package ports

import "github.com/mendelgusmao/tony/domain/models"

type InvoiceTransferService interface {
	Transfer(models.Transfer) error
}
