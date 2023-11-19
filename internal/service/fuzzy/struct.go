package service

type CreditScore struct {
	score uint16
}

type Collateral struct {
	hasCollateral bool
	Collateral    string
}

type Income struct {
	amount float64
}

type Criminal struct {
	hasCriminalRec bool
	Offences       []string
}

type Employment struct {
	term uint8
}

type FISInput struct {
	CreditScore    float32
	Collateral     float32
	Income         float32
	CriminalRecord float32
	EmploymentTerm float32
}

type FISOutput struct {
	Creditworthiness uint
}
