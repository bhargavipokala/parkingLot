-- name: AddFloor :one
INSERT INTO floor (
    id,
    parking_lot_id
) VALUES (
    $1, $2
) RETURNING *;