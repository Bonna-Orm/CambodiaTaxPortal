package models

import "github.com/goravel/framework/support/carbon"

type Sale struct {
	Id                            uint    `gorm:"primaryKey"`
	No                            string  `gorm:"column:no"`
	Date                          string  `gorm:"column:date"`
	InvoiceNo                     string  `gorm:"column:invoice_no"`
	TypeOfCustomer                string  `gorm:"column:type_of_customer"`
	TaxIdentificationNo           string  `gorm:"column:tax_identification_no"`
	CustomerNameKh                string  `gorm:"column:customer_namekh"`
	CustomerName                  string  `gorm:"column:customer_name"`
	TypeOfGoods                   string  `gorm:"column:type_of_goods"`
	TotalAmountInclVat            float64 `gorm:"column:total_amount_Incl_vat"`
	TotalAmountExclVat            float64 `gorm:"column:total_amount_excl_vat"`
	SpecificTaxCertainMerchandise float64 `gorm:"column:specific_tax_certain_merchandise"`
	SpecificTaxCertainService     float64 `gorm:"column:specific_tax_certain_service"`
	PublicLightingTax             float64 `gorm:"column:public_lighting_tax"`
	AccommodationTax              float64 `gorm:"column:accommodation_tax"`
	PrepaymentOfTaxOnIncomeRate   float64 `gorm:"column:prepayment_of_tax_on_income_rate"`
	Sector                        string  `gorm:"column:sector"`
	TreasuryCreditNoteNo          string  `gorm:"column:treasury_credit_note_no"`
	Description                   string  `gorm:"column:description"`
	CreatedAt                     carbon.DateTime
	UpdatedAt                     carbon.DateTime
}
