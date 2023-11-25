package invoicewebhook

import (
	"log"
	"net/http"

	"github.com/mendelgusmao/tony/application/invoice-webhook/config"
	"github.com/mendelgusmao/tony/domain/models"
	"github.com/mendelgusmao/tony/domain/ports"
)

type WebhookCallbackHandler struct {
	config       config.Specification
	eventHandler ports.CallbackEventHandler
}

func NewWebhookCallbackHandler(config config.Specification, eventHandler ports.CallbackEventHandler) *WebhookCallbackHandler {
	return &WebhookCallbackHandler{
		config:       config,
		eventHandler: eventHandler,
	}
}

func (h WebhookCallbackHandler) Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("received callback from", r.RemoteAddr)

	if r.Method != "POST" {
		log.Println("unexpected http method:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	event, err := models.NewCallbackEvent(r.Body, r.Header)

	if err != nil {
		log.Println("error extracting event:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errors := h.eventHandler.Handle(*event)

	if errors != nil {
		log.Println("errors handling event:", errors)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("finished processing callback")
	w.WriteHeader(http.StatusOK)
}
