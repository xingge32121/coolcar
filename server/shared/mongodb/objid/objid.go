package objid

import (
	"coolcar/shared/id"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// fromId converts an id to ObjectID
func FromID(id fmt.Stringer) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

//MustFromId converts an id to ObjectID
func MustFromId(id fmt.Stringer) primitive.ObjectID {
	objId, err := primitive.ObjectIDFromHex(id.String())

	if err != nil {
		panic(err)
	}
	return objId
}

func ToAccountID(oid primitive.ObjectID) id.AccountId {
	return id.AccountId(oid.Hex())
}

func ToTripID(oid primitive.ObjectID) id.TripId {
	return id.TripId(oid.Hex())
}
