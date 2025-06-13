package controllers

import (
	"CambodiaTaxPortal/app/models"

	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type SaleController struct{}

// NewSaleController returns a new SaleController instance
func NewSaleController() *SaleController {
	return &SaleController{}
}

// Index retrieves all sales
func (r *SaleController) Index(ctx http.Context) http.Response {
	var sales []models.Sale
	if err := facades.Orm().Query().Find(&sales); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to fetch sales: " + err.Error(),
		})
	}
	//return ctx.Response().Json(http.StatusOK, sales)
	return ctx.Response().View().Make("sale.tmpl", map[string]any{
		"Sales": sales})
}

// Store creates a new sale
func (r *SaleController) Store(ctx http.Context) http.Response {
	var sale models.Sale
	if err := ctx.Request().Bind(&sale); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "Invalid sale data: " + err.Error(),
		})
	}
	if err := facades.Orm().Query().Create(&sale); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to create sale: " + err.Error(),
		})
	}
	// Directly To Page Sale
	//return ctx.Response().Redirect(http.StatusFound, "/sale")

	return ctx.Response().Json(http.StatusCreated, http.Json{
		"message": "Sale created successfully",
	})
}

// Edit retrieves a sale by ID for editing
func (r *SaleController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Bind("id")

	var sale models.Sale
	if err := facades.Orm().Query().Where("id = ?", id).First(&sale); err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"error": "Sale not found: " + err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Sale Edit successfully",
		"Sale":    sale,
	})
}

// Filter retrieves sales by date
// func (r *SaleController) Filter(ctx http.Context) http.Response {
// 	startDate := ctx.Request().Query("start_date")
// 	endDate := ctx.Request().Query("end_date")

// 	var sales []models.Sale
// 	query := facades.Orm().Query()

// 	if startDate != "" && endDate != "" {
// 		query = query.Where("date >= ? AND date <= ?", startDate, endDate)
// 	} else if startDate != "" {
// 		query = query.Where("date >= ?", startDate)
// 	} else if endDate != "" {
// 		query = query.Where("date <= ?", endDate)
// 	}

// 	if err := query.Get(&sales); err != nil {
// 		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
// 			"error": "Failed to fetch sales: " + err.Error(),
// 		})
// 	}

// 	return ctx.Response().View().Make("sale.tmpl", map[string]any{
// 		"Sales": sales,
// 	})
// }

// Filter retrieves sales by date range
func (r *SaleController) Filter(ctx http.Context) http.Response {
	startDate := ctx.Request().Query("start_date")
	endDate := ctx.Request().Query("end_date")

	var sales []models.Sale
	query := facades.Orm().Query()

	// Validate date format (YYYY-MM-DD)
	const layout = "2006-01-02"
	if startDate != "" {
		if _, err := time.Parse(layout, startDate); err != nil {
			return ctx.Response().Json(http.StatusBadRequest, http.Json{
				"error": "Invalid start_date format. Use YYYY-MM-DD.",
			})
		}
	}
	if endDate != "" {
		if _, err := time.Parse(layout, endDate); err != nil {
			return ctx.Response().Json(http.StatusBadRequest, http.Json{
				"error": "Invalid end_date format. Use YYYY-MM-DD.",
			})
		}
	}

	// Build query
	if startDate != "" && endDate != "" {
		query = query.Where("date >= ? AND date <= ?", startDate, endDate)
	} else if startDate != "" {
		query = query.Where("date >= ?", startDate)
	} else if endDate != "" {
		query = query.Where("date <= ?", endDate)
	}

	if err := query.Get(&sales); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Failed to fetch sales: ",
			"error":   err.Error(),
		})
	}
	if len(sales) == 0 {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "No sales found for the specified date range",
		})
	}

	// Pass filter values back to the template for sticky form fields
	return ctx.Response().View().Make("sale.tmpl", map[string]any{
		"Sales":     sales,
		"StartDate": startDate,
		"EndDate":   endDate,
	})
	// Test return json with postman
	//return ctx.Response().Json(http.StatusOK, sales)
}
