package ports

import "github.com/mendelgusmao/tony/domain/models"

type CallbackEventHandler interface {
	Handle(models.CallbackEvent) error
}
