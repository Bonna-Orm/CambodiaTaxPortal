package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250612074623CreatePurchaseTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250612074623CreatePurchaseTable) Signature() string {
	return "20250612074623_create_purchase_table"
}

// Up Run the migrations.
func (r *M20250612074623CreatePurchaseTable) Up() error {
	if !facades.Schema().HasTable("purchase") {
		return facades.Schema().Create("purchase", func(table schema.Blueprint) {
			table.ID()
			table.String("no").Nullable()
			table.Date("date_purchase").Nullable()
			table.String("invoice_no").Nullable()
			table.String("type_of_supplier").Nullable()
			table.String("tax_identification_no").Nullable()
			table.String("supplier_namekh").Nullable()
			table.String("supplier_name").Nullable()
			table.String("type_of_goods_supplier").Nullable()
			table.Decimal("total_amount_incl_vat").Nullable()
			table.Decimal("total_amount_excl_vat").Nullable()
			table.String("description").Nullable()
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250612074623CreatePurchaseTable) Down() error {
	return facades.Schema().DropIfExists("purchase")
}
