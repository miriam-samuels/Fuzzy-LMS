package dashboard

import (
	"net/http"
	"time"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/dashboard"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/loan"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

const (
	weekly  = "weekly"
	monthly = "monthly"
	yearly  = "yearly"
)

// get data needed to populate dashboard
func GetDashboardData(w http.ResponseWriter, r *http.Request) {
	var labels, points interface{} // graph plots
	var err error
	currentUser := r.Context().Value(types.AuthCtxKey{}).(types.AuthCtxKey) // get user aking the request
	period := r.URL.Query().Get("period")

	// check graph periods
	switch period {
	case weekly:
		labels, points, err = getLoanChartPlotWeekly()
		if err != nil {
			helper.SendResponse(w, http.StatusBadRequest, false, "request failed", nil, err)
			return
		}
	case monthly:
		labels, points, err = getLoanChartPlotMonthly()
		if err != nil {
			helper.SendResponse(w, http.StatusBadRequest, false, "request failed", nil, err)
			return
		}
	case yearly:
		labels, points, err = getLoanChartPlot5Yrs()
		if err != nil {
			helper.SendResponse(w, http.StatusBadRequest, false, "request failes", nil, err)
			return
		}
	}

	//  get application count per status
	rejected := GetApplicationsCount(currentUser, "rejected")
	approved := GetApplicationsCount(currentUser, "approved")
	pending := GetApplicationsCount(currentUser, "pending")

	// get loans
	rows, err := loan.GetLoans(currentUser, "")

	// slice to store all loan applications
	loans := []loan.Loan{}
	// process query
	for rows.Next() {
		var loan loan.Loan
		err := rows.Scan(&loan.ID, &loan.LoanID, &loan.BorrowerId, &loan.Type, &loan.Term, &loan.Amount, &loan.Purpose, &loan.Status, &loan.Creditworthiness, &loan.HasCollateral, &loan.CollateralDocs, &loan.Collateral, &loan.CreatedAt)
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error getting loans", nil, err)
			return
		}

		loans = append(loans, loan)
	}

	// Form response object
	res := map[string]interface{}{
		"labels":       labels,
		"points":       points,
		"rejected":     rejected,
		"approved":     approved,
		"pending":      pending,
		"applications": loans,
	}
	helper.SendResponse(w, http.StatusOK, true, "plots successfully fetched", res)

}

func GetApplicationsCount(currentUser types.AuthCtxKey, status string) int {
	var count int
	row := dashboard.GetApplicationsCount(currentUser, status)
	row.Scan(&count)
	return count
}

func getLoanChartPlotWeekly() ([]time.Time, []int, error) {
	var days []time.Time
	var applications []int

	rows, err := dashboard.GetWeeklyPlotPoints()
	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var day time.Time
		var application int

		err = rows.Scan(&day, &application)
		if err != nil {
			return nil, nil, err
		}

		days = append(days, day)
		applications = append(applications, application)

	}

	return days, applications, nil
}

func getLoanChartPlotMonthly() ([]string, []int, error) {
	var months []string
	var applications []int

	rows, err := dashboard.GetMonthlyPlotPoints()
	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var month time.Time
		var application int

		err = rows.Scan(&month, &application)
		if err != nil {
			return nil, nil, err
		}

		months = append(months, month.Month().String())
		applications = append(applications, application)

	}

	return months, applications, nil
}

func getLoanChartPlot5Yrs() ([]int, []int, error) {
	var days []int
	var applications []int

	rows, err := dashboard.GetPlotPoints5Yrs()
	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var day time.Time
		var application int

		err = rows.Scan(&day, &application)
		if err != nil {
			return nil, nil, err
		}

		days = append(days, day.Year())
		applications = append(applications, application)

	}

	return days, applications, nil
}
