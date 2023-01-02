package catalog

import "errors"

//  alias
type DatabaseId = int
type SchemaId = int

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
	return &DatabaseCatalog{id, name, 1, make(map[SchemaId]*SchemaCatalog), make(map[string]SchemaId)}
}

func (db *DatabaseCatalog) addSchema(name string) (SchemaId, error) {
	if _, exist := db.nameIdMapping[name]; exist {
		return 0, errors.New("name already exist")
	}
	schemaId := db.nextSchemeId
	schema := SchemaCatalog{id: schemaId, name: name}
	db.schemas[schemaId] = &schema
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

type SchemaCatalog struct {
	id   SchemaId
	name string
}

type attribute struct {
}
