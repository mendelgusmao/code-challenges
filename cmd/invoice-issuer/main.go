package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/mendelgusmao/tony/adapters"
	invoiceissuer "github.com/mendelgusmao/tony/application/invoice-issuer"
	"github.com/mendelgusmao/tony/application/invoice-issuer/config"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/utils/checks"
)

const appName = "TONY_INVOICE_ISSUER"

func main() {
	var err error
	var config config.Specification
	err = envconfig.Process(appName, &config)

	if err != nil {
		log.Fatal(err)
	}

	privateKeyContent, err := os.ReadFile(config.PrivateKeyPath)

	if err != nil {
		log.Fatal("reading private key:", err)
	}

	log.Println("initializing tony invoice issuer")

	starkbankUser := project.Project{
		Id:          config.ProjectID,
		PrivateKey:  checks.CheckPrivateKey(string(privateKeyContent)),
		Environment: checks.CheckEnvironment(config.Environment),
	}

	starkBankInvoiceService := adapters.NewStarkBankInvoiceService(starkbankUser)
	invoiceIssuer := invoiceissuer.NewInvoiceIssuer(config, starkBankInvoiceService)
	invoiceIssuerWorker := invoiceissuer.NewInvoiceIssuerWorker(config, invoiceIssuer)

	invoiceIssuerWorker.Work().Wait()
}
