package places

import "context"

func (r *queryResolver) GetDetailsPageByID(ctx context.Context, id int) (*DetailsPage, error) {
	return nil, nil
}
func (r *queryResolver) GetAllDetailsPage(ctx context.Context) ([]DetailsPage, error) {
	return []DetailsPage{}, nil
}
