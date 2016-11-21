package skt

import (
	"reflect"
	"strings"
)

// DDLStruct is DDL struct
type DDLStruct struct {
	StructName   string
	CreateString string
	DropString   string
}

// GetDDL get create string
func GetDDL(i interface{}) DDLStruct {
	var ds DDLStruct
	re := reflect.TypeOf(i).Elem()
	ds.StructName = strings.ToLower(re.Name())
	var rf []string
	for n := 0; n < re.NumField(); n++ {
		tag := re.Field(n).Tag
		// 判断是否有核心结构，有则加入核心所有字段
		if re.Field(n).Name == "CK" {
			recore := reflect.TypeOf(&CK{}).Elem()
			for n := 0; n < recore.NumField(); n++ {
				tagcore := recore.Field(n).Tag
				rf = append(rf, strings.ToLower(recore.Field(n).Name)+" "+tagcore.Get("db"))
			}
			continue
		}
		rf = append(rf, strings.ToLower(re.Field(n).Name)+" "+tag.Get("db"))
	}
	str := strings.Join(rf, ",")
	ds.CreateString = "CREATE TABLE " + ds.StructName + "(" + str + ")"
	ds.DropString = "DROP TABLE " + ds.StructName
	return ds
}

// CreateTable create table
func (db *DB) CreateTable(ss ...interface{}) (int64, error) {
	var i int64
	for _, s := range ss {
		gd := GetDDL(s)
		_, err := db.Exec(gd.CreateString)
		if err != nil {
			return i, err
		}
		i++
	}
	return i, nil
}

// DropTable drop table
func (db *DB) DropTable(ss ...interface{}) (int64, error) {
	var i int64
	for _, s := range ss {
		gd := GetDDL(s)
		_, err := db.Exec(gd.DropString)
		if err != nil {
			return i, err
		}
		i++
	}
	return i, nil
}
