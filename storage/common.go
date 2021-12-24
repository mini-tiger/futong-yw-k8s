package storage

import "gorm.io/gorm"

// GetMysqlResultToMapSlice convert the mysql result into map slice
func GetMysqlResultToMapSlice(qc *gorm.DB) ([]map[string]interface{}, error) {
	rows, err := qc.Rows()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)
	fieldNameSlice, _ := rows.Columns()
	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columnValues := make([]interface{}, len(fieldNameSlice))
		columnPointers := make([]interface{}, len(fieldNameSlice))
		for i, _ := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range fieldNameSlice {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		for k, v := range m {
			switch vt := v.(type) {
			case []uint8:
				m[k] = string(vt)
			}
		}
		result = append(result, m)
	}
	return result, nil
}
