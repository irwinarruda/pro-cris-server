package models

type Coordinate struct {
	Latitude  float64 `json:"latitude" validate:"latitude"`
	Longitude float64 `json:"longitude" validate:"longitude"`
}
