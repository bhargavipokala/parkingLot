-- name: CreateTicket :one
INSERT INTO ticket (
    id,
    vehicle_colour,
    vehicle_type,
    registration_no,
    status,
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;


-- name: GetTicketById :one
SELECT * FROM ticket WHERE id = $1;