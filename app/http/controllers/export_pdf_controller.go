package controllers

import (
	"CambodiaTaxPortal/app/models"
	"bytes"
	"fmt"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/jung-kurt/gofpdf/v2"
)

func (r *ExportController) ExportPDF(ctx http.Context) http.Response {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Sales Report")

	// Example: Add table headers
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	headers := []string{"No.", "Date", "Invoice No.", "Customer", "Amount"}
	for _, h := range headers {
		pdf.CellFormat(38, 10, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Example: Add data rows (replace with your DB data)
	var sales []models.Sale
	if err := facades.Orm().Query().Find(&sales); err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"message": "Failed to fetch sales",
			"error":   err.Error(),
		})
	}
	for _, sale := range sales {
		pdf.CellFormat(38, 10, fmt.Sprintf("%v", sale.No), "1", 0, "C", false, 0, "")
		pdf.CellFormat(38, 10, sale.Date, "1", 0, "C", false, 0, "")
		pdf.CellFormat(38, 10, sale.InvoiceNo, "1", 0, "C", false, 0, "")
		pdf.CellFormat(38, 10, sale.CustomerNameKh, "1", 0, "C", false, 0, "")
		pdf.CellFormat(38, 10, fmt.Sprintf("%v", sale.TotalAmountExclVat), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"message": "Failed to generate PDF",
			"error":   err.Error(),
		})
	}

	filename := "sales_" + time.Now().Format("20060102_150405") + ".pdf"
	ctx.Response().Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	return ctx.Response().Data(
		http.StatusOK,
		"application/pdf",
		buf.Bytes(),
	)
}
