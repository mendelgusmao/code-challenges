package adapters

import (
	"encoding/json"
	"log"

	"github.com/mendelgusmao/tony/domain/enums"
	"github.com/mendelgusmao/tony/domain/models"
	"github.com/mendelgusmao/tony/domain/ports"
	starkbankEvent "github.com/starkbank/sdk-go/starkbank/event"
	"github.com/starkinfra/core-go/starkcore/user/project"
)

type StarkBankEventHandler struct {
	project             project.Project
	transferDestination models.TransferDestination
	transferService     ports.TransferService
}

func NewStarkBankEventHandler(project project.Project, transferDestination models.TransferDestination, transferService ports.TransferService) *StarkBankEventHandler {
	return &StarkBankEventHandler{
		project:             project,
		transferDestination: transferDestination,
		transferService:     transferService,
	}
}

func (h *StarkBankEventHandler) Handle(callbackEvent models.CallbackEvent) error {
	eventContent := starkbankEvent.Parse(
		callbackEvent.Data,
		callbackEvent.Signature,
		h.project,
	)

	event := models.BaseEvent{}
	err := json.Unmarshal([]byte(eventContent.(string)), &event)

	if err != nil {
		return err
	}

	if event.Event.Log.Invoice.Status != enums.InvoiceStatusPaid {
		log.Printf("invoice #%v has status %s. ignoring", event.Event.Log.Invoice.Id, event.Event.Log.Invoice.Status)
		return nil
	}

	log.Printf("transferring value from invoice #%v to destination account", event.Event.Log.Invoice.Id)

	transfer := h.toTransferFromInvoiceEvent(event)
	_, err = h.transfer(transfer)

	if err != nil {
		return err
	}

	log.Printf("finished transfer for invoice #%v", event.Event.Log.Invoice.Id)

	return nil
}

func (h *StarkBankEventHandler) transfer(transfer models.Transfer) (*models.Transfer, error) {
	transfers := []models.Transfer{transfer}
	starkTransfers, errors := h.transferService.CreateTransfers(transfers)

	if len(errors) > 0 {
		return nil, errors[0]
	}

	return &starkTransfers[0], nil
}

func (h *StarkBankEventHandler) toTransferFromInvoiceEvent(event models.BaseEvent) models.Transfer {
	invoice := event.Event.Log.Invoice
	amount := invoice.Amount - invoice.Fee

	return models.Transfer{
		Amount:        amount,
		BankCode:      h.transferDestination.BankCode,
		BranchCode:    h.transferDestination.Branch,
		AccountNumber: h.transferDestination.AccountNumber,
		AccountType:   h.transferDestination.AccountType,
		TaxID:         invoice.TaxId,
		Name:          invoice.Name,
	}
}
