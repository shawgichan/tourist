package places

import (
	"context"

	"github.com/shawgichan/tourist/graph"
)

func (r *queryResolver) GetDetailsPageByID(ctx context.Context, id int) (*DetailsPage, error) {
	detailsPage, err := mapDetailsPage(ctx, id, r)
	if err != nil {
		return nil, err
	}
	return &detailsPage, nil
}
func (r *queryResolver) GetAllDetailsPage(ctx context.Context) ([]DetailsPage, error) {
	details, err := mapAllDetailsPage(ctx, r)
	if err != nil {
		return []DetailsPage{}, err
	}
	return details, nil
}

func mapDetailsPage(ctx context.Context, id int, r *queryResolver) (DetailsPage, error) {
	var outputs DetailsPage

	output, err := r.Store.GetPlace(ctx, int64(id))
	if err != nil {
		return DetailsPage{}, err
	}
	outputs.ID = graph.ConvertToInt(int(output.ID))
	outputs.Name = &output.Name
	outputs.Description = &output.Description.String
	//outputs.Image = &output.Image
	//outputs.Location = &output.Location

	return outputs, nil
}

func mapAllDetailsPage(ctx context.Context, r *queryResolver) ([]DetailsPage, error) {
	var outputs []DetailsPage

	output, err := r.Store.GetPlaces(ctx)
	if err != nil {
		return []DetailsPage{}, err
	}
	for _, v := range output {
		outputs = append(outputs, DetailsPage{
			ID:          graph.ConvertToInt(int(v.ID)),
			Name:        &v.Name,
			Description: &v.Description.String,
			//Image:       &v.Image,
			//Location:    &v.Location,
		})
	}

	return outputs, nil
}
