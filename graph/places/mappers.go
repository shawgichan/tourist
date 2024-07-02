package places

import "context"

type PlaceDynamicData struct {
	Places string `json:"places"`
}

func (r *queryResolver) QueryCheckerForPlaces(ctx context.Context, query string) (*PlaceDynamicData, error) {
	return &PlaceDynamicData{
		Places: query,
	}, nil
}
