package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250606045825CreateSaleTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250606045825CreateSaleTable) Signature() string {
	return "20250606045825_create_sale_table"
}

// Up Run the migrations.
func (r *M20250606045825CreateSaleTable) Up() error {
	if !facades.Schema().HasTable("sales") {
		return facades.Schema().Create("sales", func(table schema.Blueprint) {
			table.ID()
			table.String("no", 50).Nullable()
			table.String("invoice_no", 50).Nullable()
			table.Date("date").Nullable()
			table.String("type_of_customer", 100).Nullable()
			table.String("tax_identification_no", 50).Nullable()
			table.String("customer_namekh", 100).Nullable()
			table.String("customer_name", 100).Nullable()
			table.String("type_of_goods", 100).Nullable()
			table.Decimal("total_amount_Incl_vat").Nullable()
			table.Decimal("total_amount_excl_vat").Nullable()
			table.Decimal("specific_tax_certain_merchandise").Nullable()
			table.Decimal("specific_tax_certain_service").Nullable()
			table.Decimal("public_lighting_tax").Nullable()
			table.Decimal("accommodation_tax").Nullable()
			table.Decimal("prepayment_of_tax_on_income_rate").Nullable()
			table.String("sector", 100).Nullable()
			table.String("treasury_credit_note_no", 50).Nullable()
			table.String("description", 255).Nullable()
			table.Timestamps()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250606045825CreateSaleTable) Down() error {
	return facades.Schema().DropIfExists("sales")
}
