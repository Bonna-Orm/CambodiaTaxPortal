package models

import "github.com/goravel/framework/support/carbon"

type Sale struct {
	Id            uint        `gorm:"primaryKey"`
	No            string      `gorm:"column:no"`
	InvoiceNo     string      `gorm:"column:invoice_no"`
	CustomsNo     string      `gorm:"column:customs_no"`
	Customer      string      `gorm:"column:customer"`
	Amount        float64     `gorm:"column:amount"`
	SaleDate      carbon.Date `gorm:"column:sale_date"`
	Description   string      `gorm:"column:description"`
	TID           string      `gorm:"column:tid"`
	PassportNo    string      `gorm:"column:passport_no"`
	Nationality   string      `gorm:"column:nationality"`
	Gender        string      `gorm:"column:gender"`
	DOB           carbon.Date `gorm:"column:dob"`
	Phone         string      `gorm:"column:phone"`
	Email         string      `gorm:"column:email"`
	Salary        float64     `gorm:"column:salary"`
	FringeBenefit float64     `gorm:"column:fringe_benefit"`
	CreatedAt     carbon.DateTime
	UpdatedAt     carbon.DateTime
}
