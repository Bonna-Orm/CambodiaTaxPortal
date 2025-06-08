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

	// Define a style for header alignment (centered)
	headerStyle, err := file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
		Font: &excelize.Font{
			Bold: true,
		},
	})
	if err == nil {
		// Apply the style to the header area (A1:T2)
		file.SetCellStyle("Sheet1", "A1", "T2", headerStyle)
	}

	// Merge cells for grouped and single headers
	file.MergeCell("Sheet1", "A1", "A2") // No.
	file.MergeCell("Sheet1", "B1", "B2") // Date
	file.MergeCell("Sheet1", "C1", "C2") // Invoice no./customs declaration no.
	file.MergeCell("Sheet1", "D1", "G1") // Customer group
	file.MergeCell("Sheet1", "H1", "H2") // Type of goods supplied / services rendered*
	file.MergeCell("Sheet1", "I1", "I2") // Total amount include VAT*
	file.MergeCell("Sheet1", "J1", "J2") // Total amount exclude VAT / VAT 0%
	file.MergeCell("Sheet1", "K1", "K2") // Specific tax on certain merchandise
	file.MergeCell("Sheet1", "L1", "L2") // Specific tax on certain services
	file.MergeCell("Sheet1", "M1", "M2") // Public Lighting Tax
	file.MergeCell("Sheet1", "N1", "N2") // Accommodation Tax
	file.MergeCell("Sheet1", "O1", "O2") // Prepayment of Tax on Income Rate
	file.MergeCell("Sheet1", "P1", "P2") // Sector
	file.MergeCell("Sheet1", "Q1", "Q2") // Treasury credit note no.
	file.MergeCell("Sheet1", "R1", "R2") // Description*
	file.MergeCell("Sheet1", "S1", "S2") // Create At
	file.MergeCell("Sheet1", "T1", "T2") // Update At

	// Set main headers (row 1)
	file.SetCellValue("Sheet1", "A1", "No.")
	file.SetCellValue("Sheet1", "B1", "Date")
	file.SetCellValue("Sheet1", "C1", "Invoice no./customs declaration no.")
	file.SetCellValue("Sheet1", "D1", "Customer")
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

	// Set sub-headers for "Customer" (row 2)
	file.SetCellValue("Sheet1", "D2", "Type of customer*")
	file.SetCellValue("Sheet1", "E2", "Tax identification no.")
	file.SetCellValue("Sheet1", "F2", "Name (Khmer)")
	file.SetCellValue("Sheet1", "G2", "Name (Latin)")

	// Fill data rows
	for i, sale := range sales {
		row := i + 3
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), sale.No)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), sale.InvoiceNo)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), sale.Date)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), sale.TypeOfCustomer) // Adjust if you have CustomerType
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), sale.TaxIdentificationNo)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), sale.CustomerNameKh)                // Name (Khmer) - fill if available
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), sale.CustomerName)                  // Name (Latin) - adjust if needed
		file.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), sale.TypeOfGoods)                   // Type of goods/services - fill if available
		file.SetCellValue("Sheet1", fmt.Sprintf("I%d", row), sale.TotalAmountInclVat)            // Total amount include VAT
		file.SetCellValue("Sheet1", fmt.Sprintf("J%d", row), sale.TotalAmountExclVat)            // Exclude VAT - fill if available
		file.SetCellValue("Sheet1", fmt.Sprintf("K%d", row), sale.SpecificTaxCertainMerchandise) // Specific tax merchandise
		file.SetCellValue("Sheet1", fmt.Sprintf("L%d", row), sale.SpecificTaxCertainService)     // Specific tax services
		file.SetCellValue("Sheet1", fmt.Sprintf("M%d", row), sale.PublicLightingTax)             // Public Lighting Tax
		file.SetCellValue("Sheet1", fmt.Sprintf("N%d", row), sale.AccommodationTax)              // Accommodation Tax
		file.SetCellValue("Sheet1", fmt.Sprintf("O%d", row), sale.PrepaymentOfTaxOnIncomeRate)   // Prepayment of Tax on Income Rate
		file.SetCellValue("Sheet1", fmt.Sprintf("P%d", row), sale.Sector)                        // Sector
		file.SetCellValue("Sheet1", fmt.Sprintf("Q%d", row), sale.TreasuryCreditNoteNo)          // Treasury credit note no.
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
