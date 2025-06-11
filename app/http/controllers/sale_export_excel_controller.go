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
	startDate := ctx.Request().Query("start_date")
	endDate := ctx.Request().Query("end_date")

	var sales []models.Sale
	query := facades.Orm().Query()

	// Optional: Validate date format
	const layout = "2006-01-02"
	if startDate != "" {
		if _, err := time.Parse(layout, startDate); err != nil {
			return ctx.Response().Status(400).Json(http.Json{
				"error": "Invalid start date format. Use YYYY-MM-DD.",
			})
		}
	}

	if endDate != "" {
		if _, err := time.Parse(layout, endDate); err != nil {
			return ctx.Response().Status(400).Json(http.Json{
				"error": "Invalid end date format. Use YYYY-MM-DD.",
			})
		}
	}

	// Filter by date if provided
	if startDate != "" && endDate != "" {
		query = query.Where("date >= ? AND date <= ?", startDate, endDate)
	} else if startDate != "" {
		query = query.Where("date >= ?", startDate)
	} else if endDate != "" {
		query = query.Where("date <= ?", endDate)
	}
	if err := query.Get(&sales); err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"message": "Failed to fetch sales",
			"error":   err.Error(),
		})

	}
	// Test Return Data Json with Postman
	//return ctx.Response().Json(http.StatusOK, sales)

	// if err := facades.Orm().Query().Find(&sales); err != nil {
	// 	return ctx.Response().Status(500).Json(http.Json{
	// 		"message": "Failed to fetch sales",
	// 		"error":   err.Error(),
	// 	})
	// }

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
		file.SetCellStyle("Sheet1", "A2", "T3", headerStyle)
	}

	// Merge cells for grouped and single headers
	file.MergeCell("sheet1", "A2", "A3")
	file.MergeCell("sheet1", "B2", "B3")
	file.MergeCell("sheet1", "C2", "C3")
	file.MergeCell("sheet1", "D2", "G2")
	file.MergeCell("sheet1", "H2", "H3")
	file.MergeCell("sheet1", "I2", "I3")
	file.MergeCell("sheet1", "J2", "J3")
	file.MergeCell("sheet1", "K2", "K3")
	file.MergeCell("sheet1", "L2", "L3")
	file.MergeCell("sheet1", "M2", "M3")
	file.MergeCell("sheet1", "N2", "N3")
	file.MergeCell("sheet1", "O2", "O3")
	file.MergeCell("sheet1", "P2", "P3")
	file.MergeCell("sheet1", "Q2", "Q3")
	file.MergeCell("sheet1", "R2", "R3")
	file.MergeCell("sheet1", "S2", "S3")
	file.MergeCell("sheet1", "T2", "T3")

	// Set main headers (row 1)
	file.SetCellValue("Sheet1", "A2", "ល.រ*\nNo.*")
	file.SetCellValue("Sheet1", "B2", "កាលបរិច្ឆេទ*\nDate*")
	file.SetCellValue("Sheet1", "C2", "លេខវិក្កយបត្រ ឬ ប្រតិវេទន៍គយ*\nInvoice no./customs declaration no.")
	file.SetCellValue("Sheet1", "D2", "អ្នកទិញ\nCustomer")
	file.SetCellValue("Sheet1", "H2", "ប្រភេទផ្គត់ផ្គង់ទំនិញ* ឬសេវាកម្ម\nType of goods supplied / services rendered*")
	file.SetCellValue("Sheet1", "I2", "តម្លៃសរុបជាប់ អតប*\nTotal amount include VAT*")
	file.SetCellValue("Sheet1", "J2", "តម្លៃសរុបមិនជាប់ អតប ឬ អតប អត្រា ០%\nTotal amount exclude VAT / VAT 0%")
	file.SetCellValue("Sheet1", "K2", "អាករពិសេសលើទំនិញមួយចំនួន\nSpecific tax on certain merchandise")
	file.SetCellValue("Sheet1", "L2", "អាករពិសេសលើសេវាមួយចំនួន\nSpecific tax on certain services")
	file.SetCellValue("Sheet1", "M2", "អាករបំភ្លឺសាធារណៈ\nPublic Lighting Tax")
	file.SetCellValue("Sheet1", "N2", "អាករលើការស្នាក់នៅ\nAccommodation Tax")
	file.SetCellValue("Sheet1", "O2", "អត្រាប្រាក់រំដោះពន្ធលើប្រាក់ចំណូល\nPrepayment of Tax on Income Rate")
	file.SetCellValue("Sheet1", "P2", "វិស័យ\nSector")
	file.SetCellValue("Sheet1", "Q2", "លេខឥណទានរតនាគារជាតិ\nTreasury credit note no.")
	file.SetCellValue("Sheet1", "R2", "បរិយាយ*\nDescription*")
	file.SetCellValue("Sheet1", "S2", "Create At")
	file.SetCellValue("Sheet1", "T2", "Update At")

	// Set sub-headers for "Customer" (row 2)
	file.SetCellValue("Sheet1", "D3", "ប្រភេទ*\nType of customer*")
	file.SetCellValue("Sheet1", "E3", "លេខសម្គាល់*\nTax identification no.")
	file.SetCellValue("Sheet1", "F3", "ឈ្មោះ(ខ្មែរ)\nName (Khmer)")
	file.SetCellValue("Sheet1", "G3", "ឈ្មោះ(ឡាតាំង)\nName (Latin)")

	// Set row heights for better appearance
	file.SetRowHeight("Sheet1", 2, 50)
	file.SetRowHeight("Sheet1", 3, 50)

	// Fill data rows
	for i, sale := range sales {
		row := i + 4
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

	filename := "export_sale_" + time.Now().Format("20060102_150405") + ".xlsx"
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
