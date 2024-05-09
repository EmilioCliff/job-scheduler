-- name: CreateJob :one
INSERT INTO jobs (
    description, client_name, client_address, client_email, price, start_date, end_date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetJob :one
SELECT * FROM jobs
WHERE id = $1 LIMIT 1;

-- name: ListJobs :many
SELECT * FROM jobs;

-- name: UpdateJob :one
UPDATE jobs
    set description = $1,
    client_name = $2,
    client_address = $3,
    client_email = $4,
    price = $5,
    start_date = $6,
    end_date = $7
WHERE id = $8
RETURNING *;

-- name: DeleteJob :exec
DELETE FROM jobs
WHERE id = $1;