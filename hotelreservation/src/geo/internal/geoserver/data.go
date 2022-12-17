package geoserver

import "go.mongodb.org/mongo-driver/bson/primitive"

type HotelLocation struct {
	HotelId  primitive.ObjectID `json:"hotel_id" bson:"hotel_id"`
	Location GeoJson            `json:"location" bson:"location"`
}

type GeoJson struct {
	Type string `json:"type" bson:"type"`
	// list the longitude first, and then latitude
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
