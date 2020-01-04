package config

import (
	"errors"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	expectedConfig := Specification{
		Source: "source.to/json",
	}

	os.Setenv("ZAPCHALLENGE_SOURCE", expectedConfig.Source)
	os.Setenv("ZAPBACKEND_PORTALS", "")

	if err := Load(); err != nil {
		t.Fatal(err)
	}

	if Backend.Source != expectedConfig.Source {
		t.Fatal("expected config doesnt match")
	}
}

func TestHook(t *testing.T) {
	ok := false

	Hook(func(backend Specification) error {
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

func TestHookError(t *testing.T) {
	Hook(func(backend Specification) error {
		return errors.New("dummy")
	})

	if err := Load(); err == nil {
		t.Fatal("expected an error")
	}
}
