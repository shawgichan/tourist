package places

import (
	"context"
	"time"

	"github.com/shawgichan/tourist/graph"
)

func (r *queryResolver) GetHomePage(ctx context.Context) (*HomePage, error) {
	//query := graphql.GetOperationContext(ctx).RawQuery
	//place, err := r.Store.GetPlace(ctx, int64(id))
	//if err != nil {
	//	return nil, err
	//}
	//data, errData := r.QueryCheckerForPlaces(ctx, query)
	homePage, err := mapHomePage(ctx, r)
	if err != nil {
		return nil, err
	}
	return &homePage, nil
}

func mapHomePage(ctx context.Context, r *queryResolver) (HomePage, error) {
	var outputs HomePage
	_, err := r.Store.GetPlaces(ctx)
	if err != nil {
		return HomePage{}, err
	}
	// for _, v := range output {

	// }
	return outputs, nil
}

func (r *queryResolver) GetDetailsPageByID(ctx context.Context, id int) (*DetailsPage, error) {
	//query := graphql.GetOperationContext(ctx).RawQuery
	// place, err := r.Store.GetPlace(ctx, int64(id))
	// if err != nil {
	// 	return nil, err
	// }
	// data, errData := r.QueryCheckerForPlaces(ctx, query)
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
	outputs.Image = &output.CoverImageUrl
	outputs.Time = mapDateObj(&output.OpeningHours, &output.ClosingHours)
	outputs.Date = mapDateObj(&output.CreatedAt, &output.UpdatedAt)
	outputs.Location = mapLocationObj(output.LocationID)
	outputs.Tags = mapAllTagObj([]int{1, 2, 3, 4, 5, 6})
	outputs.Features = mapAllFeatureObj([]int{1, 2, 3, 4, 5, 6})
	outputs.SimilarEvents = mapAllSimEventsObj([]int{1, 2, 3, 4, 5, 6})

	//outputs.Image = &output.Image
	//outputs.Location = &output.Location

	return outputs, nil
}

func mapDateObj(from *time.Time, to *time.Time) *DateObj {
	return &DateObj{
		From: from,
		To:   to,
	}
}
func mapAllTagObj(id []int) []*TagObj {
	//put static data for now
	var tagObjs []*TagObj
	for i := range id {
		tagObjs = append(tagObjs, mapTagObj(i))
	}
	return tagObjs
}
func mapTagObj(id int) *TagObj {
	//put static data for now
	var name, color string
	switch id {
	case 1:
		name = "Food"
		color = "#f7b731"
	case 2:
		name = "Nightlife"
		color = "#f7b731"
	case 3:
		name = "Nightlife"
		color = "#f7b731"
	case 4:
		name = "Nightlife"
		color = "#f7b731"
	case 5:
		name = "Nightlife"
		color = "#f7b731"
	case 6:
		name = "Nightlife"
		color = "#f7b731"
	default:
		name = "Unknown"
		color = "#f7b731"
	}

	return &TagObj{
		ID:    &id,
		Name:  &name,
		Color: &color,
	}
}
func mapLocationObj(locationID int64) *LocationObj {
	var city string
	var lat, lng float64
	switch locationID {
	case 1:
		city = "Tokyo"
		lat = 35.689487
		lng = 139.691711
	case 2:
		city = "Kyoto"
		lat = 35.015
		lng = 135.752
	case 3:
		city = "Osaka"
		lat = 34.693728
		lng = 135.502166
	case 4:
		city = "Nagoya"
		lat = 35.181167
		lng = 136.906389
	case 5:
		city = "Yokohama"
		lat = 35.444167
		lng = 139.638889
	case 6:
		city = "Hiroshima"
		lat = 34.39783
		lng = 132.45075
	default:
		city = "Unknown"
		lat = 0
		lng = 0
	}
	return &LocationObj{
		City: &city,
		Lat:  &lat,
		Lng:  &lng,
	}
}
func mapAllFeatureObj(id []int) []*FeatureObj {
	//put static data for now
	var featureObjs []*FeatureObj
	for i := range id {
		featureObjs = append(featureObjs, mapFeatureObj(i))
	}
	return featureObjs
}
func mapFeatureObj(id int) *FeatureObj {
	//put static data for now
	var name, value string
	switch id {
	case 1:
		name = "Wifi"
		value = "Yes"
	case 2:
		name = "Parking"
		value = "Yes"
	case 3:
		name = "Parking"
		value = "Yes"
	case 4:
		name = "Parking"
		value = "Yes"
	case 5:
		name = "Parking"
		value = "Yes"
	case 6:
		name = "Parking"
		value = "Yes"
	default:
		name = "Unknown"
		value = "Unknown"
	}
	return &FeatureObj{
		ID:    &id,
		Name:  &name,
		Value: &value,
	}
}
func mapAllSimEventsObj(id []int) []*SimEventsObj {
	//put static data for now
	var simEventsObjs []*SimEventsObj
	for i := range id {
		simEventsObjs = append(simEventsObjs, mapSimEventsObj(i))
	}
	return simEventsObjs
}
func mapSimEventsObj(id int) *SimEventsObj {
	//put static data for now
	var name, image string
	var location *LocationObj
	var date *DateObj
	var rating float64
	switch id {
	case 1:
		name = "Hokkaido Festival"
		image = "https://www.hokkaido-festival.com/wp-content/uploads/2019/04/Hokkaido-Festival-2019-04-17-1024x576.jpg"
		location = mapLocationObj(1)
		date = mapDateObj(&time.Time{}, &time.Time{})
		rating = 4.5
	case 2:
		name = "Hokkaido Festival"
		image = "https://www.hokkaido-festival.com/wp-content/uploads/2019/04/Hokkaido-Festival-2019-04-17-1024x576.jpg"
		location = mapLocationObj(2)
		date = mapDateObj(&time.Time{}, &time.Time{})
		rating = 4.5
	case 3:
		name = "Hokkaido Festival"
		image = "https://www.hokkaido-festival.com/wp-content/uploads/2019/04/Hokkaido-Festival-2019-04-17-1024x576.jpg"
		location = mapLocationObj(3)
		date = mapDateObj(&time.Time{}, &time.Time{})
		rating = 4.5
	case 4:
		name = "Hokkaido Festival"
		image = "https://www.hokkaido-festival.com/wp-content/uploads/2019/04/Hokkaido-Festival-2019-04-17-1024x576.jpg"
		location = mapLocationObj(4)
		date = mapDateObj(&time.Time{}, &time.Time{})
		rating = 4.5
	case 5:
		name = "Hokkaido Festival"
		image = "https://www.hokkaido-festival.com/wp-content/uploads/2019/04/Hokkaido-Festival-2019-04-17-1024x576.jpg"
		location = mapLocationObj(5)
		date = mapDateObj(&time.Time{}, &time.Time{})
		rating = 4.5
	case 6:
		name = "Hokkaido Festival"
		image = "https://www.hokkaido-festival.com/wp-content/uploads/2019/04/Hokkaido-Festival-2019-04-17-1024x576.jpg"
		location = mapLocationObj(6)
		date = mapDateObj(&time.Time{}, &time.Time{})
		rating = 4.5
	default:
		name = "Unknown"
		image = "Unknown"
		location = mapLocationObj(0)
		date = mapDateObj(&time.Time{}, &time.Time{})
		rating = 0
	}
	return &SimEventsObj{
		ID:       &id,
		Name:     &name,
		Image:    &image,
		Location: location,
		Rating:   &rating,
		Date:     date,
	}
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
