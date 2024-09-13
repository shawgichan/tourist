package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	util "github.com/shawgichan/tourist/utils"
	"github.com/stretchr/testify/require"
)

func testCreatePlace(t *testing.T) {
	arg := CreatePlaceParams{
		Name:              util.RandomString(6),
		Description:       pgtype.Text{},
		OpeningHours:      time.Time{},
		ClosingHours:      time.Time{},
		Rating:            pgtype.Numeric{},
		TicketCategory:    0,
		TicketPrice:       "",
		LocationID:        0,
		PlaceTypeID:       0,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		CoverImageUrl:     "",
		ProfileImageUrl:   "",
		ResturantBranchID: 0,
		PreferenceMatch:   []int64{},
	}
	place, err := testingStore.CreatePlace(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, place)
	require.NotZero(t, place)
}
