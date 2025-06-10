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

type SaleCrDrExportController struct{}

func NewSaleCrDrExportController() *ExportController {
	return &ExportController{}
}

func (r *ExportController) SaleCrDrExportExcel(ctx http.Context) http.Response {

	var salecrdrs []models.SaleCrDr
	if err := facades.Orm().Query().Find(&salecrdrs); err != nil {
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
			Bold:   true,
			Family: "Khmer OS Siemreap", // Supports Khmer and English
			Size:   11,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4CAF50"}, // Green background
			Pattern: 1,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})
	if err == nil {
		// Apply the style to the header area (A1:T2)
		file.SetCellStyle("Sheet1", "A2", "P3", headerStyle)
	}

	// Merge cells for grouped and single headers
	file.MergeCell("sheet1", "A2", "A3")
	file.MergeCell("sheet1", "B2", "E2")
	file.MergeCell("sheet1", "F2", "P2")
	// Set main headers (row 2)
	file.SetCellValue("Sheet1", "A2", "ល.រ*\nNo.*")
	file.SetCellValue("Sheet1", "B2", "អ្នកទិញ\nCustomer")
	file.SetCellValue("Sheet1", "F2", "ព័ត៌មានវិក្កយបត្រថ្មី\nNew Invoice Detail Information")
	// Set sub-headers for "Customer" (row 3)
	file.SetCellValue("Sheet1", "B3", "កាលបរិច្ឆេទ*\nDate*")
	file.SetCellValue("Sheet1", "C3", "លេខវិក្កយបត្រ ឬ ប្រតិវេទន៍គយ*\nInvoice no. / customs declaration no.*")
	file.SetCellValue("Sheet1", "D3", "ប្រភេទ*\nType of supplier*")
	file.SetCellValue("Sheet1", "E3", "លេខសម្គាល់*\nTax identification no.*")
	file.SetCellValue("Sheet1", "F3", "ប្រភេទការកែតម្រូវ*\nType of adjustment*")
	file.SetCellValue("Sheet1", "G3", "កាលបរិច្ឆេទ*\nDate*")
	file.SetCellValue("Sheet1", "H3", "លេខសម្គាល់ \nCredit/ Debit Note no.*")
	file.SetCellValue("Sheet1", "I3", "តម្លៃសរុបជាប់ អតប*\nTotal amount include VAT*")
	file.SetCellValue("Sheet1", "J3", "តម្លៃសរុបមិនជាប់ អតប ឬ អតប អត្រា ០%\nTotal amount exclude VAT / VAT 0%")
	file.SetCellValue("Sheet1", "K3", "អាករពិសេសលើទំនិញមួយចំនួន\nSpecific Tax on certain merchandise")
	file.SetCellValue("Sheet1", "L3", "អាករពិសេសលើសេវាមួយចំនួន\nSpecific tax on certain services")
	file.SetCellValue("Sheet1", "M3", "អាករបំភ្លឺសាធារណៈ\nPublic Lighting Tax")
	file.SetCellValue("Sheet1", "N3", "អាករលើការស្នាក់នៅ\nAccommodation Tax")
	file.SetCellValue("Sheet1", "O3", "អត្រាប្រាក់រំដោះពន្ធលើប្រាក់ចំណូល\nPrepayment of Tax on Income Rate")
	file.SetCellValue("Sheet1", "P3", "បរិយាយ*\nDescription*")

	// Set row heights for better appearance
	file.SetRowHeight("Sheet1", 2, 50)
	file.SetRowHeight("Sheet1", 3, 50)

	// Fill data rows
	for i, salecrdr := range salecrdrs {
		row := i + 4
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), salecrdr.No)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), salecrdr.DateCrDr)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), salecrdr.InvoiceNo)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), salecrdr.TypeOfSupplier)
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), salecrdr.TaxIdentificationNo)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row), salecrdr.TypeOfAdjustment)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row), salecrdr.DateOfAdjustment)
		file.SetCellValue("Sheet1", fmt.Sprintf("H%d", row), salecrdr.CreditDebitNoteNo)
		file.SetCellValue("Sheet1", fmt.Sprintf("I%d", row), salecrdr.TotalAmountInclVat)
		file.SetCellValue("Sheet1", fmt.Sprintf("J%d", row), salecrdr.TotalAmountExclVat)
		file.SetCellValue("Sheet1", fmt.Sprintf("K%d", row), salecrdr.SpecificTaxCertainMerchandise)
		file.SetCellValue("Sheet1", fmt.Sprintf("L%d", row), salecrdr.SpecificTaxCertainService)
		file.SetCellValue("Sheet1", fmt.Sprintf("M%d", row), salecrdr.PublicLightingTax)
		file.SetCellValue("Sheet1", fmt.Sprintf("N%d", row), salecrdr.AccommodationTax)
		file.SetCellValue("Sheet1", fmt.Sprintf("O%d", row), salecrdr.PrepaymentOfTaxOnIncomeRate)
		file.SetCellValue("Sheet1", fmt.Sprintf("P%d", row), salecrdr.Description)
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
