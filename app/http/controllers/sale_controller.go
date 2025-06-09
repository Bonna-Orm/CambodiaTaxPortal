package controllers

import (
	"CambodiaTaxPortal/app/models"

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
	return ctx.Response().Redirect(http.StatusFound, "/sale")
}

// Show retrieves a sale by ID
func (r *SaleController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var sale models.Sale
	if err := facades.Orm().Query().Where("id", id).First(&sale); err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{"error": "Sale not found"})
	}
	return ctx.Response().Json(http.StatusOK, sale)
}

// Update modifies an existing sale by ID
func (r *SaleController) Update(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var sale models.Sale
	if err := facades.Orm().Query().Where("id", id).First(&sale); err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{"error": "Sale not found"})
	}
	if err := ctx.Request().Bind(&sale); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": "Invalid sale data: " + err.Error()})
	}
	if err := facades.Orm().Query().Save(&sale); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": "Failed to update sale: " + err.Error()})
	}
	return ctx.Response().Json(http.StatusOK, sale)
}

// Destroy deletes a sale by ID
func (r *SaleController) Destroy(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	_, err := facades.Orm().Query().Where("id", id).Delete(&models.Sale{})
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": "Failed to delete sale: " + err.Error()})
	}
	return ctx.Response().Json(http.StatusOK, http.Json{"message": "Sale deleted"})
}

// CrudPage displays sales in CRUD page
func (r *SaleController) CrudPage(ctx http.Context) http.Response {
	var sales []models.Sale
	if err := facades.Orm().Query().Find(&sales); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": "Failed to fetch sales: " + err.Error()})
	}
	return ctx.Response().View().Make("sale.tmpl", map[string]any{
		"Sales": sales,
	})
}

// StoreOrDelete handles create or delete based on _method override
func (r *SaleController) StoreOrDelete(ctx http.Context) http.Response {
	if ctx.Request().Input("_method") == "DELETE" {
		id := ctx.Request().Route("id")
		_, err := facades.Orm().Query().Where("id", id).Delete(&models.Sale{})
		if err != nil {
			return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": "Failed to delete sale: " + err.Error()})
		}
		return ctx.Response().Redirect(http.StatusFound, "/sale")
	}

	var sale models.Sale
	if err := ctx.Request().Bind(&sale); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{"error": "Invalid sale data: " + err.Error()})
	}
	if err := facades.Orm().Query().Create(&sale); err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{"error": "Failed to create sale: " + err.Error()})
	}
	return ctx.Response().Redirect(http.StatusFound, "/sale")
}

// Edit displays the edit page for a sale
func (r *SaleController) Edit(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	var sale models.Sale
	err := facades.Orm().Query().Where("id", id).First(&sale)
	if err != nil {
		return ctx.Response().Json(http.StatusNotFound, http.Json{"error": "Sale not found"})
	}
	return ctx.Response().View().Make("edit.tmpl", http.Json{
		"Sale": sale,
	})
}
