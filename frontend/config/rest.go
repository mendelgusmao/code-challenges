package config

import resty "gopkg.in/resty.v1"

func init() {
	AfterLoad(setRestyHostURL)
}

func setRestyHostURL(c *Specification) error {
	resty.SetHostURL(c.BackendAddress)

	return nil
}
