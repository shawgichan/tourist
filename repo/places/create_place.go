package places

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/shawgichan/tourist/db/sqlc"
	"github.com/shawgichan/tourist/utils"
)

type Server struct {
	SQLStore db.Store
}

func NewPlaceServer(store db.Store) *Server {
	return &Server{SQLStore: store}
}

type createPlaceRequest struct {
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	OpeningHours      string         `json:"opening_hours"`
	ClosingHours      string         `json:"closing_hours"`
	Rating            pgtype.Numeric `json:"rating"`
	TicketCategory    int16          `json:"ticket_category"`
	TicketPrice       string         `json:"ticket_price"`
	LocationID        int64          `json:"location_id"`
	PlaceTypeID       int64          `json:"place_type_id"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
	CoverImageUrl     string         `json:"cover_image_url"`
	ProfileImageUrl   string         `json:"profile_image_url"`
	ResturantBranchID int64          `json:"resturant_branch_id"`
	PreferenceMatch   []int64        `json:"preference_match"`
}

func (server *Server) CreateNewPlace(ctx *gin.Context) {
	var req createPlaceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	OpeningHours, _ := time.Parse("2006-01-02T15:04:05Z", req.OpeningHours)
	ClosingHours, _ := time.Parse("2006-01-02T15:04:05Z", req.ClosingHours)
	CreatedAt, _ := time.Parse("2006-01-02T15:04:05Z", req.CreatedAt)
	UpdatedAt, _ := time.Parse("2006-01-02T15:04:05Z", req.UpdatedAt)
	arg := db.CreatePlaceParams{
		Name:              req.Name,
		Description:       pgtype.Text{String: req.Description, Valid: true},
		OpeningHours:      OpeningHours,
		ClosingHours:      ClosingHours,
		Rating:            req.Rating,
		TicketCategory:    req.TicketCategory,
		TicketPrice:       req.TicketPrice,
		LocationID:        req.LocationID,
		PlaceTypeID:       req.PlaceTypeID,
		CreatedAt:         CreatedAt,
		UpdatedAt:         UpdatedAt,
		CoverImageUrl:     req.CoverImageUrl,
		ProfileImageUrl:   req.ProfileImageUrl,
		ResturantBranchID: req.ResturantBranchID,
		PreferenceMatch:   req.PreferenceMatch,
	}

	place, err := server.SQLStore.CreatePlace(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, utils.SuccessResponse(place))
}
