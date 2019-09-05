package config

import (
	"errors"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	expectedConfig := Specification{
		Address: "localhost:4321",
	}

	os.Setenv("TAXCHALLENGE_ADDRESS", expectedConfig.Address)
	os.Setenv("TAXCHALLENGE_TAXRULES", "")

	if err := Load(); err != nil {
		t.Fatal(err)
	}

	if Backend.Address != expectedConfig.Address {
		t.Fatal("expected config doesnt match")
	}
}

func TestAfterLoad(t *testing.T) {
	ok := false

	AfterLoad(func(backend Specification) error {
		ok = true

		return nil
	})

	if err := Load(); err != nil {
		t.Fatal(err)
	}

	if ok != true {
		t.Fatal("expected side effect didnt happen")
	}
}

func TestAfterLoadError(t *testing.T) {
	AfterLoad(func(backend Specification) error {
		return errors.New("dummy")
	})

	if err := Load(); err == nil {
		t.Fatal("expected an error")
	}
}
