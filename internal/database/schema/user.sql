-- name: CheckUser :one
SELECT 1 FROM users WHERE email=$1;

-- name: CreateUser :one
INSERT INTO users (id, firstname, lastname, email, password, role) 
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: CreateBorrower :many
INSERT INTO borrowers (id, firstname, lastname, email) 
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByEmail :one
SELECT id, firstname, lastname, email, password, role FROM users 
WHERE email= $1;

-- name: GetUserById :one
SELECT id, firstname, lastname, email, role FROM users 
WHERE id = $1;

-- name: UpdateBorrower :exec
UPDATE borrowers 
SET 
  phone = $1, 
  birth_date = $2, 
  gender = $3, 
  nationality = $4, 
  state_origin = $5, 
  address = $6, 
  passport = $7, 
  signature = $8, 
  job = $9, 
  job_term = $10, 
  income = $11, 
  deck = $12, 
  has_criminal_record = $13, 
  offences = $14, 
  jail_time = $15, 
  kin = $16, 
  guarantor = $17, 
  nin = $18, 
  bvn = $19, 
  bank_name = $20, 
  account_number = $21, 
  identification = $22, 
  progress = $23,
  credit_score = $24
WHERE id = $25 RETURNING *;
