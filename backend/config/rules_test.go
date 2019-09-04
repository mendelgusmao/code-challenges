package config

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	validRulesContent = `{
    "1 >= 1": "price + 1"
  }
  `
	invalidRulesContent = "invalid"
)

func TestLoadRules(t *testing.T) {
	tempfile, err := ioutil.TempFile(os.TempDir(), "taxrules.*.json")
	defer os.Remove(tempfile.Name())

	if err != nil {
		t.Fatal(err)
	}

	tempfile.WriteString(validRulesContent)

	if err := loadRules(tempfile.Name()); err != nil {
		t.Fatal("expecting json file to be loaded")
	}
}

func TestLoadRulesNonexistentFile(t *testing.T) {
	tempfile, err := ioutil.TempFile(os.TempDir(), "taxrules.*.json")

	if err != nil {
		t.Fatal(err)
	}

	tempfile.WriteString(validRulesContent)
	os.Remove(tempfile.Name())

	if err := loadRules(tempfile.Name()); err == nil {
		t.Fatal("expecting json file to not to be found")
	}
}

func TestLoadRulesInvalidFile(t *testing.T) {
	tempfile, err := ioutil.TempFile(os.TempDir(), "rules.*.json")
	defer os.Remove(tempfile.Name())

	if err != nil {
		t.Fatal(err)
	}

	tempfile.WriteString(invalidRulesContent)

	if err := loadRules(tempfile.Name()); err == nil {
		t.Fatal("expecting json file to not to be found")
	}
}
