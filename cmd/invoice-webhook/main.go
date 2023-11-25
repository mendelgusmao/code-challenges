package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/mendelgusmao/tony/adapters"
	invoicewebhook "github.com/mendelgusmao/tony/application/invoice-webhook"
	"github.com/mendelgusmao/tony/application/invoice-webhook/config"
	"github.com/mendelgusmao/tony/domain/models"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
)

const appName = "TONY_INVOICE_WEBHOOK"

func main() {
	var config config.Specification
	err := envconfig.Process(appName, &config)

	if err != nil {
		log.Fatal(err)
	}

	privateKeyContent, err := os.ReadFile(config.PrivateKeyPath)

	if err != nil {
		log.Fatal("reading private key:", err)
	}

	log.Println("initializing tony invoice webhook")

	starkbankUser := project.Project{
		Id:          config.ProjectID,
		PrivateKey:  checks.CheckPrivateKey(string(privateKeyContent)),
		Environment: checks.CheckEnvironment(config.Environment),
	}

	transferDestination := models.TransferDestination{
		BankCode:      config.DestinationBankCode,
		Branch:        config.DestinationBranch,
		AccountNumber: config.DestinationAccountNumber,
		Name:          config.DestinationName,
		TaxID:         config.DestinationTaxID,
		AccountType:   config.DestinationAccountType,
	}

	starkBankTransferService := adapters.NewStarkBankTransferService(starkbankUser)
	starkBankEventHandler := adapters.NewStarkBankEventHandler(starkbankUser, transferDestination, starkBankTransferService)
	webhookCallbackHandler := invoicewebhook.NewWebhookCallbackHandler(config, starkBankEventHandler)

	http.HandleFunc(config.Endpoint, webhookCallbackHandler.Handler)
	log.Fatal(http.ListenAndServe(config.ListenAddress, nil))
}
