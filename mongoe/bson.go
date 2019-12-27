package mongoe

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// 忽略 bson.M 中的空值
func OmitEmpty(m bson.M) bson.M {
	noEmptyM := bson.M{}
	for k, v := range m {
		value := reflect.ValueOf(v)
		if !value.IsZero() {
			noEmptyM[k] = v
		}
	}
	return noEmptyM
}

// convert struct to bson list
func ToBsonList(structs interface{}) []bson.M {
	var record []bson.M
	v := reflect.ValueOf(structs)

	for i := 0; i < v.Len(); i++ {
		s := v.Index(i)
		bson := ToBson(s.Interface())
		record = append(record, bson)
	}
	return record
}

// convert struct to bson
func ToBson(structure interface{}) bson.M {
	result := make(bson.M)
	t := reflect.TypeOf(structure)
	v := reflect.ValueOf(structure)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.IsZero() {
			continue
		}

		tag := t.Field(i).Tag
		name := t.Field(i).Name
		key := tag.Get("bson")

		switch field.Kind() {
		case reflect.Int, reflect.Int64:
			v := field.Int()
			result[key] = v
			break
		case reflect.String:
			v := field.String()
			result[key] = v
			break
		case reflect.Slice:
			v := field.Interface()
			result[key] = v
			break
		// ObjectID
		case reflect.Array:
			f := reflect.Indirect(v).FieldByName(name)
			result[key] = f.Interface()
			break
		case reflect.Struct:
			v := getField(structure, name)
			result[key] = v
			break
		}
	}

	return result
}

// get struct field value
func getField(v interface{}, field string) interface{} {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	fieldValue := f.Interface()

	switch fieldValue.(type) {
	case time.Time:
		return f.Interface()
	default:
		return getFieldString(fieldValue)
	}
}

func getFieldString(fieldValue interface{}) string {
	switch v := fieldValue.(type) {
	case int64:
		return strconv.FormatInt(v, 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int:
		return strconv.FormatInt(int64(v), 10)
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	default:
		fmt.Println("value", v)
		return ""
	}
}
