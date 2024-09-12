-- name: CreateDonation :one
INSERT INTO donations (food_type, quantity, expiration, location)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetDonations :many
SELECT id, food_type, quantity, expiration, location FROM donations;
