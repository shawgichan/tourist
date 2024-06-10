// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Event struct {
	ID               int64       `json:"id"`
	Name             string      `json:"name"`
	Description      pgtype.Text `json:"description"`
	StartDate        pgtype.Date `json:"start_date"`
	EndDate          pgtype.Date `json:"end_date"`
	PlacesID         pgtype.Int8 `json:"places_id"`
	EventType        pgtype.Int8 `json:"event_type"`
	OrganizerName    pgtype.Text `json:"organizer_name"`
	OrganizerWebsite pgtype.Text `json:"organizer_website"`
	TicketUrl        pgtype.Text `json:"ticket_url"`
	ImageID          pgtype.Int8 `json:"image_id"`
	IsFree           pgtype.Bool `json:"is_free"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

type EventTrip struct {
	ID        int64       `json:"id"`
	EventID   pgtype.Int8 `json:"event_id"`
	TripID    pgtype.Int8 `json:"trip_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type EventsReview struct {
	ID        int64          `json:"id"`
	PlaceID   pgtype.Int8    `json:"place_id"`
	UserID    pgtype.Int8    `json:"user_id"`
	Rating    pgtype.Numeric `json:"rating"`
	Review    pgtype.Text    `json:"review"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Location struct {
	ID        int64     `json:"id"`
	Lat       string    `json:"lat"`
	Lng       string    `json:"lng"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Place struct {
	ID                int64          `json:"id"`
	Name              string         `json:"name"`
	Description       pgtype.Text    `json:"description"`
	OpeningHours      pgtype.Time    `json:"opening_hours"`
	ClosingHours      pgtype.Time    `json:"closing_hours"`
	Rating            pgtype.Numeric `json:"rating"`
	TicketCategory    pgtype.Int2    `json:"ticket_category"`
	TicketPrice       pgtype.Text    `json:"ticket_price"`
	LocationID        pgtype.Int8    `json:"location_id"`
	PlaceTypeID       pgtype.Int8    `json:"place_type_id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	CoverImageUrl     pgtype.Text    `json:"cover_image_url"`
	ProfileImageUrl   pgtype.Text    `json:"profile_image_url"`
	ResturantBranchID pgtype.Int8    `json:"resturant_branch_id"`
	PreferenceMatch   []int64        `json:"preference_match"`
}

type PlacesEvent struct {
	ID        int64              `json:"id"`
	Name      pgtype.Text        `json:"name"`
	PlacesID  pgtype.Int8        `json:"places_id"`
	Date      pgtype.Timestamptz `json:"date"`
	EventType pgtype.Int8        `json:"event_type"`
	EventID   pgtype.Int8        `json:"event_id"`
}

type PlacesReview struct {
	ID        int64          `json:"id"`
	PlaceID   pgtype.Int8    `json:"place_id"`
	UserID    pgtype.Int8    `json:"user_id"`
	Rating    pgtype.Numeric `json:"rating"`
	Review    pgtype.Text    `json:"review"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type Profile struct {
	ID              int64       `json:"id"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	AddressesID     int64       `json:"addresses_id"`
	ProfileImageUrl string      `json:"profile_image_url"`
	PhoneNumber     string      `json:"phone_number"`
	CompanyNumber   string      `json:"company_number"`
	WhatsappNumber  string      `json:"whatsapp_number"`
	Gender          int64       `json:"gender"`
	AllLanguagesID  []int64     `json:"all_languages_id"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	RefNo           string      `json:"ref_no"`
	CoverImageUrl   pgtype.Text `json:"cover_image_url"`
}

type Restaurant struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RestaurantBranch struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	ResturantID pgtype.Int8 `json:"resturant_id"`
}

type StopOffPoint struct {
	ID             int64              `json:"id"`
	TripID         pgtype.Int8        `json:"trip_id"`
	PlacesID       pgtype.Int8        `json:"places_id"`
	SequenceNumber pgtype.Int2        `json:"sequence_number"`
	ArrivalTime    pgtype.Timestamptz `json:"arrival_time"`
	DepartureTime  pgtype.Timestamptz `json:"departure_time"`
	Notes          pgtype.Text        `json:"notes"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

type Trip struct {
	ID                    int64          `json:"id"`
	UserID                pgtype.Int8    `json:"user_id"`
	Name                  string         `json:"name"`
	Description           pgtype.Text    `json:"description"`
	OriginLocationID      pgtype.Int8    `json:"origin_location_id"`
	DestinationLocationID pgtype.Int8    `json:"destination_location_id"`
	PrimaryPlaceID        pgtype.Int8    `json:"primary_place_id"`
	StartDate             pgtype.Date    `json:"start_date"`
	EndDate               pgtype.Date    `json:"end_date"`
	Status                pgtype.Int8    `json:"status"`
	TravelMode            pgtype.Int8    `json:"travel_mode"`
	EstimatedCost         pgtype.Numeric `json:"estimated_cost"`
	Visibility            pgtype.Int8    `json:"visibility"`
	SharedWith            []string       `json:"shared_with"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
}

type TripFeedback struct {
	ID        int64          `json:"id"`
	TripID    pgtype.Int8    `json:"trip_id"`
	UserID    pgtype.Int8    `json:"user_id"`
	Rating    pgtype.Numeric `json:"rating"`
	Feedback  pgtype.Text    `json:"feedback"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type User struct {
	ID             int64       `json:"id"`
	Email          string      `json:"email"`
	Username       string      `json:"username"`
	HashedPassword pgtype.Text `json:"hashed_password"`
	Status         int64       `json:"status"`
	RolesID        pgtype.Int8 `json:"roles_id"`
	ProfilesID     int64       `json:"profiles_id"`
	UserTypesID    int64       `json:"user_types_id"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

type UserFollowing struct {
	ID      int64       `json:"id"`
	UserID  pgtype.Int8 `json:"user_id"`
	PlaceID pgtype.Int8 `json:"place_id"`
}

type UserPreference struct {
	ID                      int64       `json:"id"`
	UsersID                 pgtype.Int8 `json:"users_id"`
	TravelCompanionID       pgtype.Int2 `json:"travel_companion_id"`
	PlaceCategoryPreference pgtype.Text `json:"place_category_preference"`
	PriceRange              pgtype.Text `json:"price_range"`
	EventTypePreference     pgtype.Text `json:"event_type_preference"`
	Interests               pgtype.Text `json:"interests"`
	Budget                  pgtype.Text `json:"budget"`
	CreatedAt               time.Time   `json:"created_at"`
}
