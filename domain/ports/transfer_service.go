package ports

import "github.com/mendelgusmao/tony/domain/models"

type TransferService interface {
	CreateTransfers([]models.Transfer) ([]models.Transfer, []error)
}
