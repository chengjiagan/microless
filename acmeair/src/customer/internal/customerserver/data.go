package customerserver

import (
	"microless/acmeair/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	CustomerId      primitive.ObjectID `json:"customer_id" bson:"_id,omitempty"`
	Username        string             `json:"username" bson:"username"`
	Status          string             `json:"status" bson:"status"`
	TotalMiles      int                `json:"total_miles" bson:"total_miles"`
	MilesYTD        int                `json:"miles_ytd" bson:"miles_ytd"`
	Address         *Address           `json:"address" bson:"address"`
	PhoneNumber     string             `json:"phone_number" bson:"phone_number"`
	PhoneNumberType string             `json:"phone_number_type" bson:"phone_number_type"`
}

type Address struct {
	StreetAddress1 string `json:"street_address1" bson:"street_address1"`
	StreetAddress2 string `json:"street_address2" bson:"street_address2"`
	City           string `json:"city" bson:"city"`
	StateProvince  string `json:"state_province" bson:"state_province"`
	Country        string `json:"country" bson:"country"`
	PostalCode     string `json:"postal_code" bson:"postal_code"`
}

func CustomerFromProto(customer *proto.CustomerInfo) *Customer {
	return &Customer{
		Username:        customer.Username,
		Status:          customer.Status.String(),
		TotalMiles:      int(customer.TotalMiles),
		MilesYTD:        int(customer.MilesYtd),
		Address:         AddressFromProto(customer.Address),
		PhoneNumber:     customer.PhoneNumber,
		PhoneNumberType: customer.PhoneNumberType.String(),
	}
}

func (c *Customer) toProto() *proto.CustomerInfo {
	status := proto.MemberShipStatus(proto.MemberShipStatus_value[c.Status])
	phoneType := proto.PhoneType(proto.PhoneType_value[c.PhoneNumberType])
	return &proto.CustomerInfo{
		CustomerId:      c.CustomerId.Hex(),
		Username:        c.Username,
		Status:          status,
		TotalMiles:      int32(c.TotalMiles),
		MilesYtd:        int32(c.MilesYTD),
		Address:         c.Address.toProto(),
		PhoneNumber:     c.PhoneNumber,
		PhoneNumberType: phoneType,
	}
}

func AddressFromProto(address *proto.AddressInfo) *Address {
	return &Address{
		StreetAddress1: address.StreetAddress1,
		StreetAddress2: address.StreetAddress2,
		City:           address.City,
		StateProvince:  address.StateProvince,
		Country:        address.Country,
		PostalCode:     address.PostalCode,
	}
}

func (a *Address) toProto() *proto.AddressInfo {
	return &proto.AddressInfo{
		StreetAddress1: a.StreetAddress1,
		StreetAddress2: a.StreetAddress2,
		City:           a.City,
		StateProvince:  a.StateProvince,
		Country:        a.Country,
		PostalCode:     a.PostalCode,
	}
}
