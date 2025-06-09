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
	facades.Route().Get("/sale/{id}", saleController.Show)
	facades.Route().Put("/sale/{id}", saleController.Update)
	facades.Route().Delete("/sale/{id}", saleController.Destroy)
	facades.Route().Get("/export", controllers.NewExportController().ExportExcel)
	facades.Route().Get("/dashboard", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("dashboard.tmpl", map[string]any{})
	})
}
