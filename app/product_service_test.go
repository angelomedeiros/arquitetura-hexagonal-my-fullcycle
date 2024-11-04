package app_test

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"hexagonal-aluno/app"
	mockapp "hexagonal-aluno/app/mocks"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mockapp.NewMockProductInterface(ctrl)
	persistence := mockapp.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil)

	service := app.ProductService{
		Persistence: persistence,
	}

	_, err := service.Get("1")
	require.Nil(t, err)
	require.Equal(t, product, product)
}