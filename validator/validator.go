package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/pokala15/parkingLot/entity"
)

func RegisterCustomValidators(v *validator.Validate) {
	_ = v.RegisterValidation("vehicleSlotValidation", validateVehicleSlots)
}

func validateVehicleSlots(fl validator.FieldLevel) bool {
	slots, ok := fl.Field().Interface().(map[entity.VehicleType]uint32)
	if !ok || len(slots) == 0 {
		return false
	}
	for vehicleType, slotNo := range slots {
		if vehicleType == entity.Truck {
			return 0 < slotNo && slotNo < 2
		} else if vehicleType == entity.Bike {
			return 0 < slotNo && slotNo < 3
		}
		return 0 < slotNo
	}
	return true
}
