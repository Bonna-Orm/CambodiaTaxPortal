package models

import (
	"time"

	"github.com/goravel/framework/support/carbon"
)

type Sale struct {
	Id                            uint            `gorm:"primaryKey" json:"id"`
	No                            string          `gorm:"column:no" json:"no"`
	Date                          time.Time       `gorm:"column:date" json:"date"`
	InvoiceNo                     string          `gorm:"column:invoice_no" json:"invoice_no"`
	TypeOfCustomer                string          `gorm:"column:type_of_customer" json:"type_of_customer"`
	TaxIdentificationNo           string          `gorm:"column:tax_identification_no" json:"tax_identification_no"`
	CustomerNameKh                string          `gorm:"column:customer_namekh" json:"customer_namekh"`
	CustomerName                  string          `gorm:"column:customer_name" json:"customer_name"`
	TypeOfGoods                   string          `gorm:"column:type_of_goods" json:"type_of_goods"`
	TotalAmountInclVat            float64         `gorm:"column:total_amount_Incl_vat" json:"total_amount_incl_vat"`
	TotalAmountExclVat            float64         `gorm:"column:total_amount_excl_vat" json:"total_amount_excl_vat"`
	SpecificTaxCertainMerchandise float64         `gorm:"column:specific_tax_certain_merchandise" json:"specific_tax_certain_merchandise"`
	SpecificTaxCertainService     float64         `gorm:"column:specific_tax_certain_service" json:"specific_tax_certain_service"`
	PublicLightingTax             float64         `gorm:"column:public_lighting_tax" json:"public_lighting_tax"`
	AccommodationTax              float64         `gorm:"column:accommodation_tax" json:"accommodation_tax"`
	PrepaymentOfTaxOnIncomeRate   float64         `gorm:"column:prepayment_of_tax_on_income_rate" json:"prepayment_of_tax_on_income_rate"`
	Sector                        string          `gorm:"column:sector" json:"sector"`
	TreasuryCreditNoteNo          string          `gorm:"column:treasury_credit_note_no" json:"treasury_credit_note_no"`
	Description                   string          `gorm:"column:description" json:"description"`
	CreatedAt                     carbon.DateTime `json:"created_at"`
	UpdatedAt                     carbon.DateTime `json:"updated_at"`
}
