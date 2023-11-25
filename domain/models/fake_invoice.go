package models

import (
	"log"
	"math/rand"
	"reflect"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/mvrilo/go-cpf"
)

type FakeInvoice struct {
	Id         string    `faker:"-"`
	Amount     int       `faker:"oneof: 999, 4990, 11990, 99990"`
	Name       string    `faker:"name"`
	TaxId      string    `faker:"tax-id"`
	Due        time.Time `faker:"due-date"`
	Expiration int       `faker:"expiration"`
	Fine       float64   `faker:"oneof: 2, 3, 5, 8"`
	Interest   float64   `faker:"oneof: 1, 3, 5, 7"`
	Tags       []string  `faker:"tags"`
}

func NewFakeInvoice() FakeInvoice {
	invoice := FakeInvoice{}
	err := faker.FakeData(&invoice)

	if err != nil {
		log.Fatal("programming error creating fake invoice:", err)
	}

	return invoice
}

func (f FakeInvoice) ToInvoice() Invoice {
	return Invoice{
		Amount:     f.Amount,
		Name:       f.Name,
		TaxId:      f.TaxId,
		Due:        &f.Due,
		Expiration: f.Expiration,
		Fine:       f.Fine,
		Interest:   f.Interest,
		Tags:       f.Tags,
	}
}

func init() {
	_ = faker.AddProvider("tax-id", func(v reflect.Value) (interface{}, error) {
		return cpf.GeneratePretty(), nil
	})

	_ = faker.AddProvider("tags", func(v reflect.Value) (interface{}, error) {
		type FakeValue struct {
			Value string `faker:"uuid_digit"`
		}
		fakeValue := FakeValue{}
		faker.FakeData(&fakeValue)

		return []string{"Test invoice", fakeValue.Value}, nil
	})

	_ = faker.AddProvider("due-date", func(v reflect.Value) (interface{}, error) {
		due := time.Now().Add(24 * time.Hour)
		return due, nil
	})

	_ = faker.AddProvider("expiration", func(v reflect.Value) (interface{}, error) {
		return 1e9 + rand.Intn(1e10-1e9), nil
	})
}
