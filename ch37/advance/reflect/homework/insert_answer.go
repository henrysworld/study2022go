package homework

import (
	"errors"
	"reflect"
	"strings"
)

var errInvalidEntity = errors.New("invalid entity")

func InsertStmt(entity interface{}) (string, []interface{}, error) {
	if entity == nil {
		return "", nil, errInvalidEntity
	}
	// var name string = ""

	val := reflect.ValueOf(entity)
	typ := val.Type()

	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return "", nil, errInvalidEntity
	}

	sb := strings.Builder{}
	//insert into datibiao () values ()
	sb.WriteString("INSERT INTO ")
	sb.WriteRune('`')
	sb.WriteString(typ.Name())
	sb.WriteRune('`')
	sb.WriteString(" (")
	fields, values := fieldNameAndValues(val)

	for i, field := range fields {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteRune('`')
		sb.WriteString(field)
		sb.WriteRune('`')

	}

	sb.WriteString(") ")
	sb.WriteString("VALUES (")

	args := make([]interface{}, 0, len(values))
	for i, field := range fields {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString("?")
		args = append(args, values[field])
	}
	if len(args) == 0 {
		return "", nil, errInvalidEntity
	}
	sb.WriteString(");")

	return sb.String(), args, nil

	// return
}

func fieldNameAndValues(val reflect.Value) ([]string, map[string]interface{}) {
	typ := val.Type()

	fieldNum := typ.NumField()
	fields := make([]string, 0, fieldNum)
	values := make(map[string]interface{}, fieldNum)
	for i := 0; i < fieldNum; i++ {
		field := typ.Field(i)
		value := val.Field(i)
		if field.Type.Kind() == reflect.Struct && field.Anonymous {
			subFields, subValues := fieldNameAndValues(value)
			for _, subField := range subFields {
				if _, ok := values[subField]; ok {
					//忽略重复字段
					continue
				}
				fields = append(fields, subField)
				values[subField] = subValues[subField]
			}
			continue
		}

		fields = append(fields, field.Name)
		values[field.Name] = value.Interface()
	}

	return fields, values
}
