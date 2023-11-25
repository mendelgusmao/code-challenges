package invoiceissuer

import (
	"log"
	"time"

	"github.com/mendelgusmao/tony/application/invoice-issuer/config"
)

type InvoiceIssuerWorker struct {
	config        config.Specification
	invoiceIssuer *InvoiceIssuer
	done          chan bool
}

func NewInvoiceIssuerWorker(config config.Specification, invoiceIssuer *InvoiceIssuer) InvoiceIssuerWorker {
	return InvoiceIssuerWorker{
		config:        config,
		invoiceIssuer: invoiceIssuer,
		done:          make(chan bool),
	}
}

func (w InvoiceIssuerWorker) Work() InvoiceIssuerWorker {
	ticker := time.NewTicker(w.config.IssuingInterval)

	go func() {
		for range ticker.C {
			w.issueInvoices()
		}

		w.done <- true
		ticker.Stop()
	}()

	w.issueInvoices()

	return w
}

func (w InvoiceIssuerWorker) issueInvoices() {
	log.Println("worker: issuing invoices")

	invoices, errors := w.invoiceIssuer.IssueInvoices()

	for _, err := range errors {
		log.Println("error issuing invoice:", err)
	}

	log.Printf("worker: issued %d invoices\n", len(invoices))
}

func (w InvoiceIssuerWorker) Wait() {
	for range w.done {
		break
	}
}
