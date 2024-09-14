package db

import (
	"context"
	"testing"

	"github.com/shawgichan/tourist/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateRestaurant(t *testing.T) {
	restaurant, err := testingStore.CreateRestaurant(context.Background(), utils.RandomString(6))
	require.NoError(t, err)
	require.NotEmpty(t, restaurant)
	require.NotZero(t, restaurant)
}

func TestGetRestaurant(t *testing.T) {
	restaurant, err := testingStore.GetRestaurant(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, restaurant)
	require.NotZero(t, restaurant)
}

func TestListRestaurants(t *testing.T) {
	arg := ListRestaurantsParams{
		Limit:  5,
		Offset: 0,
	}
	restaurants, err := testingStore.ListRestaurants(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, restaurants)
	require.NotZero(t, restaurants)
}
