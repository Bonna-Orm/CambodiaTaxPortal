package routes

import (
	"CambodiaTaxPortal/app/http/controllers"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	saleController := controllers.NewSaleController()
	facades.Route().Get("/sale", saleController.Index)
	facades.Route().Post("/sale", saleController.Store)
	facades.Route().Get("/export_sale", controllers.NewExportController().ExportExcel)
	facades.Route().Get("/sales", controllers.NewSaleController().Filter)
	facades.Route().Get("/dashboard", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("dashboard.tmpl", map[string]any{})
	})

	saleCrDrController := controllers.NewSaleCrDrController()
	facades.Route().Get("/salecrdr", saleCrDrController.Index)
	facades.Route().Post("/salecrdr", saleCrDrController.Store)
	facades.Route().Get("/export_salecrdr", controllers.NewSaleCrDrExportController().SaleCrDrExportExcel)
	facades.Route().Get("/salescrdr", controllers.NewSaleCrDrController().FilterSaleCrDr)
}
