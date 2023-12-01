package dashboard

import (
	"net/http"
	"time"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/dashboard"
)

func getLoanChartPlotWeekly(w http.ResponseWriter) error {
	var days []time.Time
	var applications []int

	rows, err := dashboard.GetWeeklyPlotPoints()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var day time.Time
		var application int

		err = rows.Scan(&day, &application)
		if err != nil {
			return err
		}

		days = append(days, day)
		applications = append(applications, application)

	}
	// Form response object
	res := map[string]interface{}{
		"labels": days,
		"points": applications,
	}

	helper.SendResponse(w, http.StatusOK, true, "plots successfully fetched", res)
	return nil
}

func getLoanChartPlotMonthly(w http.ResponseWriter) error {
	var months []string
	var applications []int

	rows, err := dashboard.GetMonthlyPlotPoints()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var month time.Time
		var application int

		err = rows.Scan(&month, &application)
		if err != nil {
			return err
		}

		months = append(months, month.Month().String())
		applications = append(applications, application)

	}
	// Form response object
	res := map[string]interface{}{
		"labels": months,
		"points": applications,
	}

	helper.SendResponse(w, http.StatusOK, true, "plots successfully fetched", res)
	return nil
}

func getLoanChartPlot5Yrs(w http.ResponseWriter) error {
	var days []int
	var applications []int

	rows, err := dashboard.GetPlotPoints5Yrs()
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var day time.Time
		var application int

		err = rows.Scan(&day, &application)
		if err != nil {
			return err
		}

		days = append(days, day.Year())
		applications = append(applications, application)

	}

	// Form response object
	res := map[string]interface{}{
		"labels": days,
		"points": applications,
	}

	helper.SendResponse(w, http.StatusOK, true, "plots successfully fetched", res)
	return nil
}

const (
	weekly  = "weekly"
	monthly = "monthly"
	yearly  = "yearly"
)

func GetLoanStats(w http.ResponseWriter, r *http.Request) {
	period := r.URL.Query().Get("period")

	switch period {
	case weekly:
		err := getLoanChartPlotWeekly(w)
		if err != nil {
			helper.SendResponse(w, http.StatusBadRequest, false, "request failed", nil, err)
		}
	case monthly:
		err := getLoanChartPlotMonthly(w)
		if err != nil {
			helper.SendResponse(w, http.StatusBadRequest, false, "request failed", nil, err)
		}
	case yearly:
		err := getLoanChartPlot5Yrs(w)
		if err != nil {
			helper.SendResponse(w, http.StatusBadRequest, false, "request failed", nil, err)
		}
	}

}
