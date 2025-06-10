package controllers

import (
	"CambodiaTaxPortal/app/models"

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
