/*
 * Copyright (c) 2020 Mikhail Knyazhev <markus621@gmail.com>.
 * All rights reserved. Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package orm

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	tagKey = "orm"
)

var (
	typeTableName = reflect.TypeOf(TableName(""))
)

type (
	TableName string
	ormModel  struct {
		Table  TableName
		Path   string
		Type   reflect.Type
		Origin reflect.Value
		Fields []ormModelField
	}
	ormModelField struct {
		Name  string
		Col   string
		Val   interface{}
		Empty bool
	}
)

func parseModel(v interface{}) *ormModel {
	val := reflect.ValueOf(v)
	ref := val.Type()

	switch ref.Kind() {
	case reflect.Struct:
		return decodeType(ref, val)
	case reflect.Ptr:
		return decodeType(ref.Elem(), val.Elem())
	}

	return nil
}

func decodeType(t reflect.Type, v reflect.Value) *ormModel {
	mod := &ormModel{
		Path:   t.PkgPath() + ":" + t.Name(),
		Type:   t,
		Fields: make([]ormModelField, 0),
	}

	for i := 0; i < t.NumField(); i++ {
		typ := t.Field(i)
		val := v.FieldByName(typ.Name)
		tag := typ.Tag.Get(tagKey)

		if typ.Type.AssignableTo(typeTableName) {
			mod.Table = TableName(tag)
			continue
		}

		if len(tag) == 0 {
			continue
		}

		tags := strings.Split(tag, ";")

		field := ormModelField{
			Name:  typ.Name,
			Col:   tags[0],
			Val:   getValueByType(val),
			Empty: val.IsZero(),
		}

		mod.Fields = append(mod.Fields, field)
	}

	return mod
}

func getValueByType(v reflect.Value) interface{} {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	case reflect.String:
		return v.String()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Bool:
		return v.Bool()
	default:
		panic(fmt.Sprintf("unknow type - %v", v.Kind()))
	}
}
