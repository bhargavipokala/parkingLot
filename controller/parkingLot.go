package controller

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	q "github.com/pokala15/parkingLot/db/sqlc"
)

type ParkingLot interface {
	// db.Querier
	Create(ctx context.Context, floorSize uint32, slotSize uint32) error
	// AddFloor()
	// AddSlot()
	// ParkVehicle()
	// UnParkVehicle()
	// Display()
}

type parkingLot struct {
	db          *sql.DB
	transaction q.Transaction
}

func NewParkingLot(db *sql.DB) *parkingLot {
	return &parkingLot{
		db:          db,
		transaction: q.NewTransaction(db),
	}
}

func (p *parkingLot) Create(ctx context.Context, floorSize uint32, slotSize uint32) error {
	return p.transaction.ExecTxn(ctx, func(qr *q.Queries) error {
		id, err := uuid.NewRandom()
		if err != nil {
			return fmt.Errorf("couldn't create parkingLot id: %s", err)
		}
		for i := 1; i <= int(floorSize); i++ {
			_, err := p.transaction.AddFloor(ctx, q.AddFloorParams{
				ID:           int32(i),
				ParkingLotID: id.String(),
			})
			if err != nil {
				return fmt.Errorf("couldn't add floor: %s ", err)
			}
			for j := 1; j <= int(slotSize); j++ {
				_, err := p.transaction.AddSlot(ctx, q.AddSlotParams{
					ID:          int32(j),
					FloorID:     int32(i),
					VehicleType: getVehicleType(j),
					IsBooked:    false,
				})
				if err != nil {
					return fmt.Errorf("couldn't add slot: %s", err)
				}
			}
		}
		return nil
	})
}

func getVehicleType(slot int) q.VehicleType {
	switch slot {
	case 1:
		return q.VehicleTypeTruck
	case 2, 3:
		return q.VehicleTypeBike
	default:
		return q.VehicleTypeCar
	}
}
