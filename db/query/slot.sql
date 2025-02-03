-- name: AddSlot :one
INSERT INTO slot (
    id,
    floor_id,
    vehicle_type,
    is_booked
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetFreeSlotByVehicleType :one
SELECT * FROM slot 
where vehicle_type = $1 and is_booked = false 
order by floor_id, id
limit 1;

-- name: GetAllFreeSlotsByVehicleType :many
SELECT floor_id, count(*) as free_slots FROM slot 
where vehicle_type = $1 and is_booked = false 
group by 1
having free_slots > 0;

-- name: GetOccupiedSlotsByVehicleType :many
SELECT floor_id, count(*) as free_slots FROM slot 
where vehicle_type = $1 and is_booked = true 
group by 1;