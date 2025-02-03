package mapper

import (
	db "github.com/pokala15/parkingLot/db/sqlc"
)

type CreateParkingLotRequest struct {
	FloorSize uint32 `json:"floor_size" binding:"required"`
	SlotSize  uint32 `json:"slot_size" binding:"required"`
}

type AddFloorRequest struct {
	ParkingLotId string `json:"parkinglot_id" binding:"required"`
	FloorSize    uint32 `json:"floor_size" binding:"required"`
}

type AddSlotRequest struct {
	ParkingLotId string                    `json:"parkinglot_id" binding:"required"`
	FloorNo      uint32                    `json:"floor_no" binding:"required"`
	Slots        map[db.VehicleType]uint32 `json:"slots" binding:"required,vehicleSlotValidation"`
}

type ParkVehicleRequest struct {
	VehicleType db.VehicleType `json:"vehicle_type" binding:"required"`
	RegNo       string         `json:"registration_number" binding:"required"`
	Colour      string         `json:"colour" binding:"required,oneof=RED,BLUE,WHITE,BLACK"`
}

type UnParkVehicleRequest struct {
	TicketId string `json:"ticket_id" binding:"required"`
}

type DisplayRequest struct {
	DisplayType DisplayType    `json:"display_type" binding:"required"`
	VehicleType db.VehicleType `json:"vehicle_type" binding:"required"`
}

type DisplayType uint32

const (
	FreeCount DisplayType = iota
	FreeSlots
	OccupiedSlots
)
