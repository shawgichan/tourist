package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateLocation(t *testing.T) {
	arg := CreateLocationParams{
		Lat:       "12345678901234567890123456789012",
		Lng:       "12345678901234567890123456789012",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	location, err := testingStore.CreateLocation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, location)
	require.Equal(t, arg.Lat, location.Lat)
	require.Equal(t, arg.Lng, location.Lng)
	require.WithinDuration(t, arg.CreatedAt, location.CreatedAt, time.Second)
	require.WithinDuration(t, arg.UpdatedAt, location.UpdatedAt, time.Second)
}
