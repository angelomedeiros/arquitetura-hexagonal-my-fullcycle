package app_test

import (
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
