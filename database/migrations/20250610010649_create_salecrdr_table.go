package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250610010649CreateSalecrdrTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250610010649CreateSalecrdrTable) Signature() string {
	return "20250610010649_create_salecrdr_table"
}

// Up Run the migrations.
func (r *M20250610010649CreateSalecrdrTable) Up() error {
	if !facades.Schema().HasTable("sale_cr_drs") {
		return facades.Schema().Create("sale_cr_drs", func(table schema.Blueprint) {
			table.ID()
			table.String("no", 50).Nullable()
			table.String("invoice_no", 50).Nullable()
			table.Date("date_cr_dr").Nullable()
			table.String("type_of_supplier", 100).Nullable()
			table.String("tax_identification_no", 50).Nullable()
			table.String("type_of_adjustment", 100).Nullable()
			table.Date("date_of_adjustment").Nullable()
			table.String("credit_debit_note_no", 100).Nullable()
			table.Decimal("total_amount_Incl_vat").Nullable()
			table.Decimal("total_amount_excl_vat").Nullable()
			table.Decimal("specific_tax_certain_merchandise").Nullable()
			table.Decimal("specific_tax_certain_service").Nullable()
			table.Decimal("public_lighting_tax").Nullable()
			table.Decimal("accommodation_tax").Nullable()
			table.Decimal("prepayment_of_tax_on_income_rate").Nullable()
			table.String("description", 100).Nullable()
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250610010649CreateSalecrdrTable) Down() error {
	return facades.Schema().DropIfExists("sale_cr_drs")
}
