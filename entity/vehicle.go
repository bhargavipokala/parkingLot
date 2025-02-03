package entity

import db "github.com/pokala15/parkingLot/db/sqlc"

type Vehicle struct {
	vehicleType    db.VehicleType
	registrationNo string
	color          string
}
