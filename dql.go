package skt

import (
	"fmt"
	"time"
)

// Select data
func (db *DB) Select(query string, args ...interface{}) ([]map[string]string, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	var field []string
	field = append(field, columns...)
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(columns))
	var vmapSlice []map[string]string
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		vmap := make(map[string]string, len(columns))
		for i, v := range values {
			switch x := v.(type) {
			case nil:
				vmap[field[i]] = ""
			case []uint8:
				vmap[field[i]] = string(x)
			case time.Time:
				vmap[field[i]] = x.Format(DATETIME)
			default:
				vmap[field[i]] = fmt.Sprint(x)
			}
		}
		vmapSlice = append(vmapSlice, vmap)
	}
	return vmapSlice, nil
}
