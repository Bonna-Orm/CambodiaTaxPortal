package controllers

import (
	"CambodiaTaxPortal/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"time"
)

type PurchaseController struct{}

func NewPurchaseController() *PurchaseController {
	return &PurchaseController{}
}

// Index retrieves all purchases
func (c *PurchaseController) Index(ctx http.Context) http.Response {
	var purchases []models.Purchase
	if err := facades.Orm().Query().Find(&purchases); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to fetch purchases: " + err.Error(),
		})
	}
	// Test Return JSON with Postman
	//return ctx.Response().Json(http.StatusOK, purchases)
	return ctx.Response().View().Make("purchase.tmpl", map[string]any{
		"Purchases": purchases,
	})
}

// Post creates a new purchase
func (c *PurchaseController) Store(ctx http.Context) http.Response {
	var purchase models.Purchase
	if err := ctx.Request().Bind(&purchase); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "Invalid input: " + err.Error(),
		})
	}

	if err := facades.Orm().Query().Create(&purchase); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Failed to create purchase: " + err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusCreated, http.Json{
		"message": "Purchase created successfully",
	})
}

// Edit retrieves a purchase by ID for editing
// Delete removes a purchase by ID
// Filter retrieves purchases by date range
func (c *PurchaseController) FilterPurchase(ctx http.Context) http.Response {
	startDate := ctx.Request().Query("start_date")
	endDate := ctx.Request().Query("end_date")
	var purchases []models.Purchase
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
		query = query.Where("date_purchase >= ? AND date_purchase <= ?", startDate, endDate)
	} else if startDate != "" {
		query = query.Where("date_purchase >= ?", startDate)
	} else if endDate != "" {
		query = query.Where("date_purchase <= ?", endDate)
	}

	if err := query.Get(&purchases); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"message": "Failed to fetch purchases: ",
			"error":   err.Error(),
		})
	}
	if len(purchases) == 0 {
		return ctx.Response().Json(http.StatusNotFound, http.Json{
			"message": "No purchases found for the specified date range",
		})
	}

	// Pass filter values back to the template for sticky form fields
	// Test return json with postman
	//return ctx.Response().Json(http.StatusOK, purchases)
	return ctx.Response().View().Make("purchase.tmpl", map[string]any{
		"Purchases": purchases,
		"StartDate": startDate,
		"EndDate":   endDate,
	})
}
