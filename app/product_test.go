package app_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"hexagonal-aluno/app"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Price = 10
	product.Status = app.DISABLED
	product.Name = "Hello"

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := app.Product{}
	product.Price = 0
	product.Status = app.ENABLED
	product.Name = "Hello"

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Price = 10
	product.Status = app.DISABLED
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid status"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than zero", err.Error())

	product.Price = 10
	product.ID = "invalid id"
	_, err = product.IsValid()
	require.NotNil(t, err)
}
