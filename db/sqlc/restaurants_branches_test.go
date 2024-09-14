package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	util "github.com/shawgichan/tourist/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateRestaurantBranch(t *testing.T) {
	arg := CreateRestaurantBranchParams{
		Name:        util.RandomString(6),
		ResturantID: pgtype.Int8{Int64: 1, Valid: true},
	}
	restaurantBranch, err := testingStore.CreateRestaurantBranch(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, restaurantBranch)
	require.NotZero(t, restaurantBranch)
}

func TestListRestaurantBranches(t *testing.T) {
	arg := ListRestaurantBranchesParams{
		Limit:       5,
		Offset:      0,
		ResturantID: pgtype.Int8{Int64: 1, Valid: true},
	}
	restaurantBranches, err := testingStore.ListRestaurantBranches(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, restaurantBranches)
	require.NotZero(t, restaurantBranches)
}
