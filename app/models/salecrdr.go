package models

import "time"

type SaleCrDr struct {
	Id                            uint      `gorm:"primaryKey" json:"id"`
	No                            string    `gorm:"column:no" json:"no"`
	InvoiceNo                     string    `gorm:"column:invoice_no" json:"invoice_no"`
	DateCrDr                      time.Time `gorm:"column:date_cr_dr" json:"date_cr_dr"`
	TypeOfSupplier                string    `gorm:"column:type_of_supplier" json:"type_of_supplier"`
	TaxIdentificationNo           string    `gorm:"column:tax_identification_no" json:"tax_identification_no"`
	TypeOfAdjustment              string    `gorm:"column:type_of_adjustment" json:"type_of_adjustment"`
	DateOfAdjustment              time.Time `gorm:"column:date_of_adjustment" json:"date_of_adjustment"`
	CreditDebitNoteNo             string    `gorm:"column:credit_debit_note_no" json:"credit_debit_note_no"`
	TotalAmountInclVat            float64   `gorm:"column:total_amount_Incl_vat" json:"total_amount_Incl_vat"`
	TotalAmountExclVat            float64   `gorm:"column:total_amount_excl_vat" json:"total_amount_excl_vat"`
	SpecificTaxCertainMerchandise float64   `gorm:"column:specific_tax_certain_merchandise" json:"specific_tax_certain_merchandise"`
	SpecificTaxCertainService     float64   `gorm:"column:specific_tax_certain_service" json:"specific_tax_certain_service"`
	PublicLightingTax             float64   `gorm:"column:public_lighting_tax" json:"public_lighting_tax"`
	AccommodationTax              float64   `gorm:"column:accommodation_tax" json:"accommodation_tax"`
	PrepaymentOfTaxOnIncomeRate   float64   `gorm:"column:prepayment_of_tax_on_income_rate" json:"prepayment_of_tax_on_income_rate"`
	Description                   string    `gorm:"column:description" json:"description"`
	CreatedAt                     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (SaleCrDr) TableName() string {
	return "sale_cr_drs"
}
