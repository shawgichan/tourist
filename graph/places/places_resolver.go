package places

import (
	"context"
	"errors"
)

func (r *queryResolver) GetDetailsPageByID(ctx context.Context, id int) (*DetailsPage, error) {
	detailsPage, err := mapDetailsPage(id)
	if err != nil {
		return nil, err
	}
	return detailsPage, nil
}
func (r *queryResolver) GetAllDetailsPage(ctx context.Context) ([]DetailsPage, error) {
	return []DetailsPage{}, nil
}

func mapDetailsPage(id int) (*DetailsPage, error) {
	return &DetailsPage{
		ID: &id,
	}, errors.New("not implemented")
}
