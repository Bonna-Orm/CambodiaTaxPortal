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
			table.Date("sale_date").Nullable()
			table.String("customs_no", 50).Nullable()
			table.String("customer", 100).Nullable()
			table.Decimal("amount").Nullable()
			table.String("description", 255).Nullable()
			table.String("tid", 50).Nullable()
			table.String("passport_no", 50).Nullable()
			table.String("nationality", 50).Nullable()
			table.String("gender", 10).Nullable()
			table.Date("dob").Nullable()
			table.String("phone", 30).Nullable()
			table.String("email", 100).Nullable()
			table.Decimal("salary").Nullable()
			table.Decimal("fringe_benefit").Nullable()
			table.Timestamps()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250606045825CreateSaleTable) Down() error {
	return facades.Schema().DropIfExists("sales")
}
