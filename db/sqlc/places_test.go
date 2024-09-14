package db

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	util "github.com/shawgichan/tourist/utils"
	"github.com/stretchr/testify/require"
)

func TestCreatePlace(t *testing.T) {
	// lat, lng := RandomLatLng()
	// location, _ := testingStore.CreateLocation(context.Background(), CreateLocationParams{
	// 	Lat:       lat,
	// 	Lng:       lng,
	// 	CreatedAt: time.Time{},
	// 	UpdatedAt: time.Time{},
	// })

	arg := CreatePlaceParams{
		Name:              util.RandomString(6),
		Description:       pgtype.Text{},
		OpeningHours:      time.Time{},
		ClosingHours:      time.Time{},
		Rating:            pgtype.Numeric{},
		TicketCategory:    0,
		TicketPrice:       "",
		LocationID:        1,
		PlaceTypeID:       0,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		CoverImageUrl:     "",
		ProfileImageUrl:   "",
		ResturantBranchID: 1,
		PreferenceMatch:   []int64{},
	}
	place, err := testingStore.CreatePlace(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, place)
	require.NotZero(t, place)
}

func TestGetPlace(t *testing.T) {
	place, err := testingStore.GetPlace(context.Background(), 4)
	require.NoError(t, err)
	require.NotEmpty(t, place)
	require.NotZero(t, place)
}

func TestGetPlaces(t *testing.T) {
	places, err := testingStore.GetPlaces(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, places)
	require.NotZero(t, places)
}

// RandomLatLng returns a random latitude and longitude
func RandomLatLng() (string, string) {
	return strconv.FormatInt(int64(util.RandomInteger(10)), 10), strconv.FormatInt(int64(util.RandomInteger(10)), 10)
}
