package catalog

import "testing"

func Test_normal(t *testing.T) {
	db := ConstructDataBase(0, "test_db")
	schemaId, _ := db.addSchema("test_sc")
	schema, _ := db.getSchema(schemaId)
	var cols []ColumnCatalog
	col1 := ColumnCatalog{0, "id", ColumnDesc{false, true, 1}}
	col2 := ColumnCatalog{0, "name", ColumnDesc{false, false, 2}}
	col3 := ColumnCatalog{0, "ct", ColumnDesc{false, false, 3}}
	cols = append(cols, col1, col2, col3)
	schema.addTable("test_tb", cols)

	// validate
	if len(db.schemas) != 1 {
		t.Errorf("schema size = %v, want %v", len(db.schemas), 1)
	}
}
