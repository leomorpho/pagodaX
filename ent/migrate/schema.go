// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FileStoragesColumns holds the columns for the "file_storages" table.
	FileStoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "bucket_name", Type: field.TypeString},
		{Name: "object_key", Type: field.TypeString},
		{Name: "original_file_name", Type: field.TypeString, Nullable: true},
		{Name: "file_size", Type: field.TypeInt64, Nullable: true},
		{Name: "content_type", Type: field.TypeString, Nullable: true},
		{Name: "file_hash", Type: field.TypeString, Nullable: true},
	}
	// FileStoragesTable holds the schema information for the "file_storages" table.
	FileStoragesTable = &schema.Table{
		Name:       "file_storages",
		Columns:    FileStoragesColumns,
		PrimaryKey: []*schema.Column{FileStoragesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "filestorage_bucket_name_object_key",
				Unique:  false,
				Columns: []*schema.Column{FileStoragesColumns[3], FileStoragesColumns[4]},
			},
		},
	}
	// PasswordTokensColumns holds the columns for the "password_tokens" table.
	PasswordTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "password_token_user", Type: field.TypeInt},
	}
	// PasswordTokensTable holds the schema information for the "password_tokens" table.
	PasswordTokensTable = &schema.Table{
		Name:       "password_tokens",
		Columns:    PasswordTokensColumns,
		PrimaryKey: []*schema.Column{PasswordTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "password_tokens_users_user",
				Columns:    []*schema.Column{PasswordTokensColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FileStoragesTable,
		PasswordTokensTable,
		UsersTable,
	}
)

func init() {
	PasswordTokensTable.ForeignKeys[0].RefTable = UsersTable
}
