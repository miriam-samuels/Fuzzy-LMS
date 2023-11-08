-- name: GetAllApplications :many
SELECT * FROM applications;

-- name: GetApplicationsByBorrower :many
SELECT * FROM applications WHERE borrowerId = $1;

-- name: CreateApplication :one
INSERT INTO applications (id,loanId, borrowerId,type,term,amount,purpose) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;