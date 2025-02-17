// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	AddFloor(ctx context.Context, arg AddFloorParams) (Floor, error)
	AddSlot(ctx context.Context, arg AddSlotParams) (Slot, error)
	CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error)
	GetAllFreeSlotsByVehicleType(ctx context.Context, vehicleType VehicleType) ([]GetAllFreeSlotsByVehicleTypeRow, error)
	GetFreeSlotByVehicleType(ctx context.Context, vehicleType VehicleType) (Slot, error)
	GetOccupiedSlotsByVehicleType(ctx context.Context, vehicleType VehicleType) ([]GetOccupiedSlotsByVehicleTypeRow, error)
	GetTicketById(ctx context.Context, id string) (Ticket, error)
}

var _ Querier = (*Queries)(nil)
