package skt

import (
	"reflect"
	"strings"
	"time"
)

// DMLStruct is DML struct
type DMLStruct struct {
	StructName   string
	InsertString string
	DeleteString string
	UpdateString string
	RowID        int64
	InsertData   []interface{}
	UpdateData   []interface{}
}

// GetDMLI is get DML string and data
func GetDMLI(i interface{}) DMLStruct {
	var ds DMLStruct
	rv := reflect.ValueOf(i).Elem()
	rt := rv.Type()
	ds.StructName = strings.ToLower(rt.Name())
	var its, itx []string
	for n := 0; n < rv.NumField(); n++ {
		// 判断值是否为空
		if !IsBlank(rv.Field(n)) {
			// 判断是否为结构体
			if rv.Field(n).Kind() == reflect.Struct {
				rs := reflect.Indirect(rv.Field(n))
				if !IsBlank(rs.FieldByName("ID")) {
					ds.RowID = rs.FieldByName("ID").Int()
				}
				continue
			}
			fn := strings.ToLower(rt.Field(n).Name)
			if fn == "id" {
				ds.RowID = rv.Field(n).Int()
			} else {
				its = append(its, fn)
				itx = append(itx, "?")
				ds.InsertData = append(ds.InsertData, rv.Field(n).Interface())
			}
		}
	}
	// 自动加创建时间
	if _, b := rt.FieldByName("Created"); b {
		its = append(its, "created")
		itx = append(itx, "?")
		ds.InsertData = append(ds.InsertData, time.Now())
	}

	itstr := strings.Join(its, ",")
	itxstr := strings.Join(itx, ",")

	ds.InsertString = "INSERT INTO " + ds.StructName + "(" + itstr + ") VALUES(" + itxstr + ")"
	ds.DeleteString = "DELETE FROM " + ds.StructName + " WHERE id=?"
	return ds
}

// GetDMLU is get DML string and data
func GetDMLU(i interface{}, ss ...string) DMLStruct {
	var ds DMLStruct
	rv := reflect.ValueOf(i).Elem()
	rt := rv.Type()
	ds.StructName = strings.ToLower(rt.Name())
	var ius []string
	// 设置必须更新为空的字段
	for _, s := range ss {
		ius = append(ius, strings.ToLower(s)+"=?")
		ds.UpdateData = append(ds.UpdateData, reflect.Zero(rv.FieldByName(s).Type()).Interface())
	}
	for n := 0; n < rv.NumField(); n++ {
		// 判断值是否为空
		if !IsBlank(rv.Field(n)) {
			// 判断是否为结构体
			if rv.Field(n).Kind() == reflect.Struct {
				rs := reflect.Indirect(rv.Field(n))
				if !IsBlank(rs.FieldByName("ID")) {
					ds.RowID = rs.FieldByName("ID").Int()
				}
				continue
			}
			fn := strings.ToLower(rt.Field(n).Name)
			if fn == "id" {
				ds.RowID = rv.Field(n).Int()
			} else {

				ius = append(ius, fn+"=?")
				ds.UpdateData = append(ds.UpdateData, rv.Field(n).Interface())
			}
		}
	}
	// 自动加更新时间
	if _, b := rt.FieldByName("Updated"); b {
		ius = append(ius, "updated=?")
		ds.UpdateData = append(ds.UpdateData, time.Now())
	}
	iustr := strings.Join(ius, ",")
	ds.UpdateData = append(ds.UpdateData, ds.RowID)
	ds.UpdateString = "UPDATE " + ds.StructName + " SET " + iustr + " WHERE id=?"
	return ds
}

// GetDMLD is get DML string and data
func GetDMLD(i interface{}) DMLStruct {
	var ds DMLStruct
	rv := reflect.ValueOf(i).Elem()
	rt := rv.Type()
	ds.StructName = strings.ToLower(rt.Name())
	for n := 0; n < rv.NumField(); n++ {
		// 判断值是否为空
		if !IsBlank(rv.Field(n)) {
			// 判断是否为结构体
			if rv.Field(n).Kind() == reflect.Struct {
				rs := reflect.Indirect(rv.Field(n))
				if !IsBlank(rs.FieldByName("ID")) {
					ds.RowID = rs.FieldByName("ID").Int()
				}
				continue
			}
			fn := strings.ToLower(rt.Field(n).Name)
			if fn == "id" {
				ds.RowID = rv.Field(n).Int()
			}
		}
	}

	ds.DeleteString = "DELETE FROM " + ds.StructName + " WHERE id=?"
	return ds
}

// Insert insert data
func (db *DB) Insert(i interface{}) (int64, error) {
	gm := GetDMLI(i)
	stmt, err := db.Prepare(gm.InsertString)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(gm.InsertData...)
	if err != nil {
		return 0, err
	}
	num, _ := res.RowsAffected()
	return num, err
}

// Delete data
func (db *DB) Delete(i interface{}) (int64, error) {
	gm := GetDMLD(i)
	stmt, err := db.Prepare(gm.DeleteString)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(gm.RowID)
	if err != nil {
		return 0, err
	}
	num, _ := res.RowsAffected()
	return num, err
}

// Update data
func (db *DB) Update(i interface{}, ss ...string) (int64, error) {
	gm := GetDMLU(i, ss...)
	// fmt.Println(gm.UpdateString, gm.UpdateData)
	stmt, err := db.Prepare(gm.UpdateString)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(gm.UpdateData...)
	if err != nil {
		return 0, err
	}
	num, _ := res.RowsAffected()
	return num, err
}
