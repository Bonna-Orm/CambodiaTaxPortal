package models

import (
	"time"
)

type Purchase struct {
	Id                  uint      `gorm:"primaryKey" json:"id"`
	No                  string    `gorm:"column:no" json:"no"`
	DatePurchase        time.Time `gorm:"column:date_purchase" json:"date_purchase"`
	InvoiceNo           string    `gorm:"column:invoice_no" json:"invoice_no"`
	TypeOfSupplier      string    `gorm:"column:type_of_supplier" json:"type_of_supplier"`
	TaxIdentificationNo string    `gorm:"column:tax_identification_no" json:"tax_identification_no"`
	SupplierNameKh      string    `gorm:"column:supplier_namekh" json:"supplier_namekh"`
	SupplierName        string    `gorm:"column:supplier_name" json:"supplier_name"`
	TypeOfGoodsSupplier string    `gorm:"column:type_of_goods_supplier" json:"type_of_goods_supplier"`
	TotalAmountInclVat  float64   `gorm:"column:total_amount_incl_vat" json:"total_amount_incl_vat"`
	TotalAmountExclVat  float64   `gorm:"column:total_amount_excl_vat" json:"total_amount_excl_vat"`
	Description         string    `gorm:"column:description" json:"description"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func (Purchase) TableName() string {
	return "purchase"
}
