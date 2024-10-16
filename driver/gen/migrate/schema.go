// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "words", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "network", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "private_key", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "balance", Type: field.TypeOther, SchemaType: map[string]string{"mysql": "decimal(36,18)"}},
		{Name: "balance_update_time", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "token_info", Type: field.TypeJSON, Nullable: true},
		{Name: "create_time", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "create_date", Type: field.TypeTime, Nullable: true},
		{Name: "is_transfer", Type: field.TypeInt8, Nullable: true, Default: 0},
		{Name: "total_token_value", Type: field.TypeOther, SchemaType: map[string]string{"mysql": "decimal(36,18)"}},
		{Name: "trx_mode", Type: field.TypeEnum, Nullable: true, Enums: []string{"transfer", "lock"}, Default: "transfer"},
		{Name: "trx_addr_type", Type: field.TypeEnum, Nullable: true, Enums: []string{"single", "multi"}, Default: "single"},
		{Name: "trx_priv_addr", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "trx_priv_pkey", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "aes_type", Type: field.TypeInt8, Nullable: true, Default: 1},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UserTable,
	}
)

func init() {
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
