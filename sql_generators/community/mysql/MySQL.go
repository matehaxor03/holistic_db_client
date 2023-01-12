package mysql

import (
	json "github.com/matehaxor03/holistic_json/json"
	validate "github.com/matehaxor03/holistic_db_client/validate"
)

type MySQL struct {
	GetCreateDatabaseSQL func(verify *validate.Validator, database_name string, character_set *string, collate *string,  options *json.Map) (*string, *json.Map, []error)
	GetCreateRecordSQL func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options *json.Map) (*string, *json.Map, []error)
	GetCreateTableSQL func(verify *validate.Validator, table_name string, table_data json.Map, options *json.Map) (*string, *json.Map, []error)
	GetDatabaseExistsSQL func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error)
	GetDropDatabaseIfExistsSQL func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error)
	GetDropDatabaseSQL func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error)
}

func NewMySQL() (*MySQL) {
	get_create_database_sql := newCreateDatabaseSQL()
	get_create_record_sql := newCreateRecordSQL()
	get_create_table_sql := newCreateTableSQL()
	get_database_exists_sql := newDatabaseExistsSQL()
	drop_database_sql := newDropDatabaseSQL()

	return &MySQL{
		GetCreateDatabaseSQL: func(verify *validate.Validator, database_name string, character_set *string, collate *string,  options *json.Map) (*string, *json.Map, []error) {	
			return get_create_database_sql.GetCreateDatabaseSQL(verify, database_name, character_set, collate, options)
		},
		GetCreateRecordSQL: func(verify *validate.Validator, table_name string, table_schema json.Map, valid_columns map[string]bool, record_data json.Map, options *json.Map) (*string, *json.Map, []error) {
			return get_create_record_sql.GetCreateRecordSQL(verify, table_name, table_schema, valid_columns, record_data, options)
		},
		GetCreateTableSQL: func(verify *validate.Validator, table_name string, table_data json.Map, options *json.Map) (*string, *json.Map, []error) {
			return get_create_table_sql.GetCreateTableSQL(verify, table_name, table_data, options)
		},
		GetDatabaseExistsSQL: func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
			return get_database_exists_sql.GetDatabaseExistsSQL(verify, database_name, options)
		},
		GetDropDatabaseIfExistsSQL: func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
			return drop_database_sql.GetDropDatabaseIfExistsSQL(verify, database_name, options)
		},
		GetDropDatabaseSQL: func(verify *validate.Validator, database_name string, options *json.Map) (*string, *json.Map, []error) {
			return drop_database_sql.GetDropDatabaseSQL(verify, database_name, options)
		},

	}
}
