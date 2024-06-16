// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package places

import (
	"time"
)

type DateObj struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}

type DetailsPage struct {
	ID            *int            `json:"id,omitempty"`
	Name          *string         `json:"name,omitempty"`
	Image         *string         `json:"image,omitempty"`
	Date          *DateObj        `json:"date,omitempty"`
	Time          *DateObj        `json:"time,omitempty"`
	Description   *string         `json:"description,omitempty"`
	Location      *LocationObj    `json:"location,omitempty"`
	Tags          []*TagObj       `json:"tags,omitempty"`
	Features      []*FeatureObj   `json:"features,omitempty"`
	SimilarEvents []*SimEventsObj `json:"similarEvents,omitempty"`
}

type FeatureObj struct {
	ID    *int    `json:"id,omitempty"`
	Icon  *string `json:"icon,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type LocationObj struct {
	City *string  `json:"city,omitempty"`
	Lat  *float64 `json:"Lat,omitempty"`
	Lng  *float64 `json:"Lng,omitempty"`
}

type Query struct {
}

type SimEventsObj struct {
	ID       *int         `json:"id,omitempty"`
	Name     *string      `json:"name,omitempty"`
	Image    *string      `json:"image,omitempty"`
	Location *LocationObj `json:"location,omitempty"`
	Rating   *float64     `json:"rating,omitempty"`
	Date     *DateObj     `json:"date,omitempty"`
}

type TagObj struct {
	ID    *int    `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}
