package config

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	validPortalContent = `
  portal_rules:
    foobar:
      rules:
        - "1 > 0"`
	invalidPortalContent         = "invalid"
	invalidPortalRulesExpression = `
  portal_rules:
    foobar:
      rules:
      - "..."`
)

func TestLoadPortals(t *testing.T) {
	tempfile, err := ioutil.TempFile(os.TempDir(), "portal-rules.*.yaml")
	defer os.Remove(tempfile.Name())

	if err != nil {
		t.Fatal(err)
	}

	tempfile.WriteString(validPortalContent)

	if err := loadPortals(tempfile.Name()); err != nil {
		t.Fatalf("expecting yaml file to be loaded, got error `%v`", err)
	}
}

func TestLoadPortalsNonexistentFile(t *testing.T) {
	tempfile, err := ioutil.TempFile(os.TempDir(), "portal-rules.*.yaml")

	if err != nil {
		t.Fatal(err)
	}

	tempfile.WriteString(validPortalContent)
	os.Remove(tempfile.Name())

	if err := loadPortals(tempfile.Name()); err == nil {
		t.Fatal("expecting yaml file to not to be found")
	}
}

func TestLoadPortalsInvalidFile(t *testing.T) {
	tempfile, err := ioutil.TempFile(os.TempDir(), "portal-rules.*.yaml")
	defer os.Remove(tempfile.Name())

	if err != nil {
		t.Fatal(err)
	}

	tempfile.WriteString(invalidPortalContent)

	if err := loadPortals(tempfile.Name()); err == nil {
		t.Fatal("expecting yaml file to be invalid")
	}
}

func TestLoadPortalsInvalidExpression(t *testing.T) {
	tempfile, err := ioutil.TempFile(os.TempDir(), "portal-rules.*.yaml")
	defer os.Remove(tempfile.Name())

	if err != nil {
		t.Fatal(err)
	}

	tempfile.WriteString(invalidPortalRulesExpression)

	if err := loadPortals(tempfile.Name()); err == nil {
		t.Fatal("expecting expression to be invalid")
	}
}
