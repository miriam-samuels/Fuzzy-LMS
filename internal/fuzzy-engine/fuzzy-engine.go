package fis

import (
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
	// fmt.Printf("%v \n %v \n%v \n%v \n%v \n", t, s, c, i, ctl)

	// fuzzify each input
	inputs.EmploymentTerm = t.fuzzify()
	inputs.CreditScore = s.fuzzify()
	inputs.CriminalRecord = c.fuzzify()
	inputs.Income = i.fuzzify()
	inputs.Collateral = ctl.fuzzify()

	// pass fuzified input into inference engine
	bad, avg, good := inputs.inference()

	//  defuzzify
	output := defuzzify(bad, avg, good)

	return output
}
