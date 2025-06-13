package routes

import (
	"CambodiaTaxPortal/app/http/controllers"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	// Welcome Page route
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	// Dashboard route
	facades.Route().Get("/dashboard", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("dashboard.tmpl", map[string]any{})
	})

	// Sale routes
	saleController := controllers.NewSaleController()
	facades.Route().Get("/sale", saleController.Index)
	facades.Route().Post("/sale", saleController.Store)
	facades.Route().Get("/export_sale", controllers.NewSaleExportController().SaleExportExcel)
	facades.Route().Get("/sales", controllers.NewSaleController().Filter)

	// Sale Credit/Debit routes
	saleCrDrController := controllers.NewSaleCrDrController()
	facades.Route().Get("/salecrdr", saleCrDrController.Index)
	facades.Route().Post("/salecrdr", saleCrDrController.Store)
	facades.Route().Get("/export_salecrdr", controllers.NewSaleCrDrExportController().SaleCrDrExportExcel)
	facades.Route().Get("/salecrdr_filter", controllers.NewSaleCrDrController().FilterSaleCrDr)

	// Purchase routes
	purchaseController := controllers.NewPurchaseController()
	facades.Route().Get("/purchase", purchaseController.Index)
	facades.Route().Post("/purchase", purchaseController.Store)
	facades.Route().Get("/purchase_filter", purchaseController.FilterPurchase)
	facades.Route().Get("/export_purchase", controllers.NewPurchaseExportController().PurchaseExportExcel)
}
