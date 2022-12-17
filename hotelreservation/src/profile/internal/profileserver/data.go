package profileserver

import (
	"microless/hotelreservation/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	HotelId     primitive.ObjectID `json:"hotel_id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Description string             `json:"description" bson:"description"`
	Address     *Address           `json:"address" bson:"address"`
	Images      []*Image           `json:"images" bson:"images"`
	RoomNumber  int32              `json:"room_number" bson:"room_number"`
}

type Address struct {
	StreetNumber string  `json:"street_number" bson:"street_number"`
	StreetName   string  `json:"street_name" bson:"street_name"`
	City         string  `json:"city" bson:"city"`
	State        string  `json:"state" bson:"state"`
	Country      string  `json:"country" bson:"country"`
	PostalCode   string  `json:"postal_code" bson:"postal_code"`
	Lat          float64 `json:"lat" bson:"lat"`
	Lon          float64 `json:"lon" bson:"lon"`
}

type Image struct {
	Url     string `json:"url" bson:"url"`
	Default bool   `json:"default" bson:"default"`
}

func (o *Hotel) toProto() *proto.Hotel {
	images := make([]*proto.Image, len(o.Images))
	for i, img := range o.Images {
		images[i] = img.toProto()
	}

	return &proto.Hotel{
		HotelId:     o.HotelId.Hex(),
		Name:        o.Name,
		PhoneNumber: o.PhoneNumber,
		Description: o.Description,
		Address:     o.Address.toProto(),
		Images:      images,
		RoomNumber:  o.RoomNumber,
	}
}

func hotelFromProto(hotel *proto.Hotel) *Hotel {
	if hotel == nil {
		return nil
	}

	images := make([]*Image, len(hotel.Images))
	for i, img := range hotel.Images {
		images[i] = imageFromProto(img)
	}

	return &Hotel{
		Name:        hotel.Name,
		PhoneNumber: hotel.PhoneNumber,
		Description: hotel.Description,
		Address:     addressFromProto(hotel.Address),
		Images:      images,
		RoomNumber:  hotel.RoomNumber,
	}
}

func (o *Address) toProto() *proto.Address {
	return &proto.Address{
		StreetNumber: o.StreetNumber,
		StreetName:   o.StreetName,
		City:         o.City,
		State:        o.State,
		Country:      o.Country,
		PostalCode:   o.PostalCode,
		Lat:          o.Lat,
		Lon:          o.Lon,
	}
}

func addressFromProto(addr *proto.Address) *Address {
	if addr == nil {
		return nil
	}

	return &Address{
		StreetNumber: addr.StreetNumber,
		StreetName:   addr.StreetName,
		City:         addr.City,
		State:        addr.State,
		Country:      addr.Country,
		PostalCode:   addr.PostalCode,
		Lat:          addr.Lat,
		Lon:          addr.Lon,
	}
}

func (o *Image) toProto() *proto.Image {
	return &proto.Image{
		Url:     o.Url,
		Default: o.Default,
	}
}

func imageFromProto(img *proto.Image) *Image {
	if img == nil {
		return nil
	}

	return &Image{
		Url:     img.Url,
		Default: img.Default,
	}
}
