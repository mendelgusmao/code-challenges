package config

type Specification struct {
	ProjectID                string `envconfig:"project_id"`
	PrivateKeyPath           string `envconfig:"private_key_path" default:"/certs/privateKey.pem"`
	Environment              string `envconfig:"environment" default:"sandbox"`
	ListenAddress            string `envconfig:"listen_address" default:":8443"`
	Endpoint                 string `envconfig:"endpoint" default:"/webhooks/stark-bank/invoice"`
	DestinationBankCode      string `envconfig:"destination_bank_code"`
	DestinationBranch        string `envconfig:"destination_branch"`
	DestinationAccountNumber string `envconfig:"destination_account_number"`
	DestinationName          string `envconfig:"destination_name"`
	DestinationTaxID         string `envconfig:"destination_tax_id"`
	DestinationAccountType   string `envconfig:"destination_account_type"`
}
