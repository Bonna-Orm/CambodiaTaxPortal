package controllers

import (
	"bytes"
	"fmt"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/xuri/excelize/v2"

	"CambodiaTaxPortal/app/models"
)

type ExportController struct{}

func NewExportController() *ExportController {
	return &ExportController{}
}

func (r *ExportController) ExportExcel(ctx http.Context) http.Response {

	var sales []models.Sale
	if err := facades.Orm().Query().Find(&sales); err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"message": "Failed to fetch sales",
			"error":   err.Error(),
		})
	}

	file := excelize.NewFile()
	index, err := file.NewSheet("Sheet1")
	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"message": "Failed to create sheet",
			"error":   err.Error(),
		})
	}

	// Set headers (Row 1)
	file.SetCellValue("Sheet1", "A1", "No.")
	file.SetCellValue("Sheet1", "B1", "Date")
	file.SetCellValue("Sheet1", "C1", "Invoice no./customs declaration no.")
	file.SetCellValue("Sheet1", "D1", "Type of customer*")
	file.SetCellValue("Sheet1", "E1", "Tax identification no.")
	file.SetCellValue("Sheet1", "F1", "Name (Khmer)")
	file.SetCellValue("Sheet1", "G1", "Name (Latin)")
	file.SetCellValue("Sheet1", "H1", "Type of goods supplied / services rendered*")
	file.SetCellValue("Sheet1", "I1", "Total amount include VAT*")
	file.SetCellValue("Sheet1", "J1", "Total amount exclude VAT / VAT 0%")
	file.SetCellValue("Sheet1", "K1", "Specific tax on certain merchandise")
	file.SetCellValue("Sheet1", "L1", "Specific tax on certain services")
	file.SetCellValue("Sheet1", "M1", "Public Lighting Tax")
	file.SetCellValue("Sheet1", "N1", "Accommodation Tax")
	file.SetCellValue("Sheet1", "O1", "Prepayment of Tax on Income Rate")
	file.SetCellValue("Sheet1", "P1", "Sector")
	file.SetCellValue("Sheet1", "Q1", "Treasury credit note no.")
	file.SetCellValue("Sheet1", "R1", "Description*")
	file.SetCellValue("Sheet1", "S1", "Create At")
	file.SetCellValue("Sheet1", "T1", "Update At")

	// Fill data rows
	for i, sale := range sales {
		row := i + 2
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), sale.No)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), sale.SaleDate.String())
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), sale.InvoiceNo)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), sale.Customer) // Adjust if you have CustomerType
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), sale.TID)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), "")            // Name (Khmer) - fill if available
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), sale.Customer) // Name (Latin) - adjust if needed
		file.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), "")            // Type of goods/services - fill if available
		file.SetCellValue("Sheet1", fmt.Sprintf("I%d", row), sale.Amount)
		file.SetCellValue("Sheet1", fmt.Sprintf("J%d", row), "") // Exclude VAT - fill if available
		file.SetCellValue("Sheet1", fmt.Sprintf("K%d", row), "") // Specific tax merchandise
		file.SetCellValue("Sheet1", fmt.Sprintf("L%d", row), "") // Specific tax services
		file.SetCellValue("Sheet1", fmt.Sprintf("M%d", row), "") // Public Lighting Tax
		file.SetCellValue("Sheet1", fmt.Sprintf("N%d", row), "") // Accommodation Tax
		file.SetCellValue("Sheet1", fmt.Sprintf("O%d", row), "") // Prepayment of Tax on Income Rate
		file.SetCellValue("Sheet1", fmt.Sprintf("P%d", row), "") // Sector
		file.SetCellValue("Sheet1", fmt.Sprintf("Q%d", row), "") // Treasury credit note no.
		file.SetCellValue("Sheet1", fmt.Sprintf("R%d", row), sale.Description)
		file.SetCellValue("Sheet1", fmt.Sprintf("S%d", row), sale.CreatedAt.String())
		file.SetCellValue("Sheet1", fmt.Sprintf("T%d", row), sale.UpdatedAt.String())
	}

	file.SetActiveSheet(index)

	// Write to buffer (no temp file needed)
	var buf bytes.Buffer
	if err := file.Write(&buf); err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"message": "Failed to write file",
			"error":   err.Error(),
		})
	}

	filename := "export_" + time.Now().Format("20060102_150405") + ".xlsx"
	ctx.Response().Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	return ctx.Response().Data(
		http.StatusOK,
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		buf.Bytes(),
	)
}

// func itoa(i int) string {
// 	return fmt.Sprintf("%d", i)
// }
