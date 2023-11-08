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
  has_collateral = $16, 
  collateral = $17, 
  collateral_docs = $18, 
  kin = $19, 
  guarantor = $20, 
  nin_slip = $21, 
  nin = $22, 
  bvn = $23, 
  bank_name = $24, 
  account = $25, 
  identification = $26, 
  loan_ids = $27 
WHERE id = $28 RETURNING *;
