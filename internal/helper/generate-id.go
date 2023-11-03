package helper

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func GenerateLoanID() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random 6-digit number.
	min := 100000 // Smallest 6-digit number
	max := 999999 // Largest 6-digit number
	id := rand.Intn(max-min+1) + min

	return "Loan" + strconv.Itoa(id)
}
