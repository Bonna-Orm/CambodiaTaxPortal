package controllers

import (
	"CambodiaTaxPortal/app/models"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type SaleCrDrController struct{}

// NewSaleCrDrController returns a new SaleCrDrController instance
func NewSaleCrDrController() *SaleCrDrController {
	return &SaleCrDrController{}
}

// Index retrieves all sale credit/debit records
func (c *SaleCrDrController) Index(ctx http.Context) http.Response {
	var saleCrDrs []models.SaleCrDr
	if err := facades.Orm().Query().Find(&saleCrDrs); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to fetch sale credit/debit records: " + err.Error(),
		})
	}
	// return ctx.Response().Json(http.StatusOK, http.Json{
	// 	"data": saleCrDrs,
	// })
	return ctx.Response().View().Make("salecrdr.tmpl", map[string]any{
		"SaleCrDrs": saleCrDrs,
	})
}

// Store creates a new sale credit/debit record
func (c *SaleCrDrController) Store(ctx http.Context) http.Response {
	var saleCrDr models.SaleCrDr
	if err := ctx.Request().Bind(&saleCrDr); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "Invalid sale credit/debit record data: " + err.Error(),
		})
	}
	if err := facades.Orm().Query().Create(&saleCrDr); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to create sale credit/debit record: " + err.Error(),
		})
	}

	return ctx.Response().Redirect(http.StatusFound, "/salecrdr")
}

// Filter retrieves sale credit/debit records based on optional filters
func (c *SaleCrDrController) FilterSaleCrDr(ctx http.Context) http.Response {
	startDate := ctx.Request().Query("start_date")
	endDate := ctx.Request().Query("end_date")

	var saleCrDrs []models.SaleCrDr
	query := facades.Orm().Query()

	// Optional: Validate date format
	const layout = "2006-01-02"
	if startDate != "" {
		if _, err := time.Parse(layout, startDate); err != nil {
			return ctx.Response().Status(400).Json(http.Json{
				"error": "Invalid start date format. Use YYYY-MM-DD.",
			})
		}
	}

	if endDate != "" {
		if _, err := time.Parse(layout, endDate); err != nil {
			return ctx.Response().Status(400).Json(http.Json{
				"error": "Invalid end date format. Use YYYY-MM-DD.",
			})
		}
	}

	// Filter by date if provided
	if startDate != "" && endDate != "" {
		query = query.Where("date_cr_dr >= ? AND date_cr_dr <= ?", startDate, endDate)
	} else if startDate != "" {
		query = query.Where("date_cr_dr >= ?", startDate)
	} else if endDate != "" {
		query = query.Where("date_cr_dr <= ?", endDate)
	}

	if err := query.Get(&saleCrDrs); err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"message": "Failed to fetch sale credit/debit records",
			"error":   err.Error(),
		})
	}
	if len(saleCrDrs) == 0 {
		return ctx.Response().Status(404).Json(http.Json{
			"message": "No sale credit/debit records found for the specified date range",
		})
	}

	// Test return data json with postman
	//return ctx.Response().Json(http.StatusOK, saleCrDrs)
	// Pass filter values back to the template for sticky form fields
	return ctx.Response().View().Make("salecrdr.tmpl", map[string]any{
		"SaleCrDrs": saleCrDrs,
		"StartDate": startDate,
		"EndDate":   endDate,
	})
}
