package catalog

import "errors"

//  alias
type DatabaseId = int32
type SchemaId = int32
type TableId = int32
type ColumnId = int32
type DataType = int32

// DatabaseCatalog :Definition for DatabaseCatalog
type DatabaseCatalog struct {
	id   DatabaseId
	name string
	// next SchemeId. this id will increase when schema add to the database catalog
	nextSchemeId SchemaId
	schemas      map[SchemaId]*SchemaCatalog
	// simple mapping for name to id, check duplicate name
	nameIdMapping map[string]SchemaId
}

func ConstructDataBase(id DatabaseId, name string) *DatabaseCatalog {
	return &DatabaseCatalog{id, name, 0, make(map[SchemaId]*SchemaCatalog), make(map[string]SchemaId)}
}

func (db *DatabaseCatalog) addSchema(name string) (SchemaId, error) {
	if _, exist := db.nameIdMapping[name]; exist {
		return -1, errors.New("name already exist")
	}
	schemaId := db.nextSchemeId
	schema := constructSchema(schemaId, name)
	db.schemas[schemaId] = schema
	db.nextSchemeId++
	return schemaId, nil
}

func (db *DatabaseCatalog) getSchema(id SchemaId) (*SchemaCatalog, bool) {
	if _, exist := db.schemas[id]; !exist {
		return nil, false
	}
	return db.schemas[id], true
}

func (db DatabaseCatalog) delSchema(id SchemaId) error {
	if _, exist := db.schemas[id]; !exist {
		return errors.New("schema not exist")
	}
	schema := db.schemas[id]
	delete(db.schemas, id)
	delete(db.nameIdMapping, schema.name)
	return nil
}

// SchemaCatalog :Definition for Schema
type SchemaCatalog struct {
	id          SchemaId
	name        string
	nextTableId TableId
	tables      map[TableId]*TableCatalog
	// simple mapping for name to id, check duplicate name
	nameIdMapping map[string]TableId
}

func constructSchema(id SchemaId, name string) *SchemaCatalog {
	return &SchemaCatalog{id, name, 0, make(map[TableId]*TableCatalog), make(map[string]TableId)}
}

func (schema *SchemaCatalog) addTable(name string, cols []ColumnCatalog) (TableId, error) {
	if _, exist := schema.nameIdMapping[name]; exist {
		return -1, errors.New("table already exist")
	}
	tableId := schema.nextTableId
	schema.nextTableId++
	table := constructTable(tableId, name, cols)
	schema.tables[tableId] = table
	schema.nameIdMapping[name] = tableId
	return tableId, nil
}

func (schema *SchemaCatalog) getTable(id TableId) (*TableCatalog, error) {
	if _, exist := schema.tables[id]; !exist {
		return nil, errors.New("table not exist")
	}
	return schema.tables[id], nil
}

func (schema *SchemaCatalog) delTable(id TableId) error {
	if _, exist := schema.tables[id]; !exist {
		return errors.New("table not exist")
	}
	table := schema.tables[id]
	delete(schema.tables, id)
	delete(schema.nameIdMapping, table.name)
	return nil
}

// TableCatalog :Definition for Table
type TableCatalog struct {
	id   TableId
	name string
	cols []ColumnCatalog
}

func constructTable(id TableId, name string, cols []ColumnCatalog) *TableCatalog {
	// generate column id
	var colId ColumnId
	colId = 0
	for _, col := range cols {
		col.id = colId
		colId++
	}
	return &TableCatalog{id, name, cols}
}

func (tb *TableCatalog) getColumn(id ColumnId) (*ColumnCatalog, error) {
	for _, col := range tb.cols {
		if col.id == id {
			return &col, nil
		}
	}
	return nil, errors.New("column not exist")
}

type ColumnCatalog struct {
	id   ColumnId
	name string
	desc ColumnDesc
}

type ColumnDesc struct {
	nullable bool
	primary  bool
	dataType DataType
}
