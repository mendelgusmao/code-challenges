package config

import "time"

type Specification struct {
	ProjectID         string        `envconfig:"project_id"`
	PrivateKeyPath    string        `envconfig:"private_key_path" default:"/certs/privateKey.pem"`
	Environment       string        `envconfig:"environment" default:"sandbox"`
	IssuingInterval   time.Duration `envconfig:"issuing_interval" default:"3h"`
	MinInvoicesAmount int           `envconfig:"min_invoices_amount" default:"8"`
	MaxInvoicesAmount int           `envconfig:"max_invoices_amount" default:"12"`
}
