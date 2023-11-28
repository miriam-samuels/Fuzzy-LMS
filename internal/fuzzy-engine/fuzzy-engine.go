package fis

import (
	"fmt"

	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/loan"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/user"
)

func AccessCreditworthiness(brw user.Borrower, application loan.Loan) float64 {
	//  variable to store fuzzified inputs
	var inputs FISInput

	//  get inputs to fuzzify
	t := Employment{brw.JobTerm}
	s := CreditScore{brw.CreditScore}
	c := Criminal{brw.HasCriminalRec, brw.Offences}
	i := Income{brw.Income}
	ctl := Collateral{application.HasCollateral, application.Collateral}

	// fuzzify each input
	inputs.EmploymentTerm = t.fuzzify()
	inputs.CreditScore = s.fuzzify()
	inputs.CriminalRecord = c.fuzzify()
	inputs.Income = i.fuzzify()
	inputs.Collateral = ctl.fuzzify()

	fmt.Printf("%#v",inputs)

	// pass fuzified input into inference engine
	output := inputs.inference()

	return output
}
