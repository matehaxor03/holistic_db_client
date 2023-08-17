package mysql

import (
	json "github.com/matehaxor03/holistic_json/json"
	validate "github.com/matehaxor03/holistic_validator/validate"
	"strings"
)

type MySQL struct {
	GetCreateDatabaseSQL func(verify *validate.Validator, database_name string, character_set *string, collate *string,  options json.Map) (*strings.Builder, json.Map, []error)
	GetCreateRecordSQL func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options json.Map) (*strings.Builder, json.Map, []error)
	GetCreateTableSQL func(verify *validate.Validator, table_name string, table_data json.Map, options json.Map) (*strings.Builder, json.Map, []error)
	GetDatabaseExistsSQL func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetDropDatabaseIfExistsSQL func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetDropDatabaseSQL func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetDropTableIfExistsSQL func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetDropTableSQL func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetTableNamesSQL func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetTableSchemaSQL func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error)
	MapTableSchemaFromDB func(verify *validate.Validator, table_name string, json_array json.Array) (*json.Map, []error)
	GetTableSchemaAdditionalSQL func(verify *validate.Validator, database_name string, table_name string, options json.Map) (*strings.Builder, json.Map, []error)
	MapAdditionalSchemaFromDBToMap func(json_array json.Array) (*json.Map, []error)
	GetSelectRecordsSQL func(verify *validate.Validator, table_name string, table_data json.Map, select_fields *json.Array, filters *json.Array, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64, options json.Map) (*strings.Builder, json.Map, []error)
	GetTableExistsSQL func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error)
	GetUpdateRecordSQL func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options json.Map) (*strings.Builder, json.Map, []error)
}

func NewMySQL() (*MySQL) {
	get_create_database_sql := newCreateDatabaseSQL()
	get_create_record_sql := newCreateRecordSQL()
	get_create_table_sql := newCreateTableSQL()
	get_database_exists_sql := newDatabaseExistsSQL()
	drop_database_sql := newDropDatabaseSQL()
	drop_table_sql := newDropTableSQL()
	get_table_names_sql := newTableNamesSQL()
	get_table_schema_sql := newTableSchemaSQL()
	get_table_schema_additional_sql := newTableSchemaAdditionalSQL()
	get_select_records_sql := newSelectRecordsSQL()
	get_table_exists_sql := newTableExistsSQL() 
	get_update_record_sql := newUpdateRecordSQL()

	return &MySQL{
		GetCreateDatabaseSQL: func(verify *validate.Validator, database_name string, character_set *string, collate *string,  options json.Map) (*strings.Builder, json.Map, []error) {	
			return get_create_database_sql.GetCreateDatabaseSQL(verify, database_name, character_set, collate, options)
		},
		GetCreateRecordSQL: func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_create_record_sql.GetCreateRecordSQL(verify, table_name, table_schema, valid_columns, record_data, options)
		},
		GetCreateTableSQL: func(verify *validate.Validator, table_name string, table_data json.Map, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_create_table_sql.GetCreateTableSQL(verify, table_name, table_data, options)
		},
		GetDatabaseExistsSQL: func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_database_exists_sql.GetDatabaseExistsSQL(verify, database_name, options)
		},
		GetDropDatabaseIfExistsSQL: func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return drop_database_sql.GetDropDatabaseIfExistsSQL(verify, database_name, options)
		},
		GetDropDatabaseSQL: func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return drop_database_sql.GetDropDatabaseSQL(verify, database_name, options)
		},
		GetDropTableIfExistsSQL: func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return drop_table_sql.GetDropTableIfExistsSQL(verify, table_name, options)
		},
		GetDropTableSQL: func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return drop_table_sql.GetDropTableSQL(verify, table_name, options)
		},
		GetTableNamesSQL: func(verify *validate.Validator, database_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_table_names_sql.GetTableNamesSQL(verify, database_name, options)
		},
		GetTableSchemaSQL: func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_table_schema_sql.GetTableSchemaSQL(verify, table_name, options)
		},
		MapTableSchemaFromDB: func(verify *validate.Validator, table_name string, json_array json.Array) (*json.Map, []error) {
			return get_table_schema_sql.MapTableSchemaFromDB(verify, table_name, json_array)
		},
		GetTableSchemaAdditionalSQL: func(verify *validate.Validator, database_name string, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_table_schema_additional_sql.GetTableSchemaAdditionalSQL(verify, database_name, table_name, options)
		},
		MapAdditionalSchemaFromDBToMap: func(json_array json.Array) (*json.Map, []error) {
			return get_table_schema_additional_sql.MapAdditionalSchemaFromDBToMap(json_array)
		},
		GetSelectRecordsSQL: func(verify *validate.Validator, table_name string, table_data json.Map, select_fields *json.Array, filters *json.Array, group_by *json.Array, order_by *json.Array, limit *uint64, offset *uint64, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_select_records_sql.GetSelectRecordsSQL(verify, table_name, table_data, select_fields, filters, group_by, order_by, limit, offset, options)
		},
		GetTableExistsSQL: func(verify *validate.Validator, table_name string, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_table_exists_sql.GetTableExistsSQL(verify, table_name, options)
		},
		GetUpdateRecordSQL: func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options json.Map) (*strings.Builder, json.Map, []error) {
			return get_update_record_sql.GetUpdateRecordSQL(verify, table_name, table_schema, valid_columns, record_data, options)
		},
	}
}
