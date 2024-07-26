package pretty

import (
	"fmt"
	"reflect"
	"strings"
)

const prettyPrintIndentNum = 2
const prettyPrintIndentChar = " "

func Value(values ...interface{}) []byte {
	str := ""
	for k, s := range values {
		str += splitOne(1, s, false)
		if k < (len(values) - 1) {
			str += strings.Repeat("-", 80)
		}
	}
	return []byte(str)
}

func splitOne(deep int, s interface{}, needIndent bool) string {
	currPrefixSpaces := strings.Repeat(prettyPrintIndentChar, deep*prettyPrintIndentNum)
	parentPrefixSpaces := ""
	if len(currPrefixSpaces) != 0 {
		parentPrefixSpaces = strings.Repeat(prettyPrintIndentChar, (deep-1)*prettyPrintIndentNum)
	}

	bufStr := ""

	if s == nil {
		if needIndent == true {
			bufStr += fmt.Sprint(currPrefixSpaces)
		}
		bufStr += fmt.Sprintln(s)
		return bufStr
	}
	switch reflect.TypeOf(s).Kind() {
	case reflect.Ptr:
		bufStr += splitOne(deep, reflect.ValueOf(s).Elem().Interface(), false)
	case reflect.Struct:
		value := reflect.ValueOf(s)
		valueType := reflect.TypeOf(s)
		bufStr += fmt.Sprintln("{")
		for i := 0; i < value.NumField(); i++ {
			fieldName := currPrefixSpaces + valueType.Field(i).Name + ": "
			bufStr += fmt.Sprintf("%c[0;0;32m%s%c[0m", 0x1B, fieldName, 0x1B)
			if value.Field(i).CanInterface() {
				bufStr += splitOne(deep+1, value.Field(i).Interface(), false)
			} else {
				bufStr += fmt.Sprintf("%+v\n", value.Field(i))
			}
		}
		bufStr += fmt.Sprintln(parentPrefixSpaces + "}")
	case reflect.Array, reflect.Slice:
		value := reflect.ValueOf(s)

		switch value.Type().Elem().Kind() {
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			buf := sliceIntType(value)
			bufStr += fmt.Sprint(string(buf))
		case reflect.String:
			buf := sliceStrType(value)
			bufStr += fmt.Sprint(string(buf))
		default:
			if value.Len() <= 0 {
				bufStr += fmt.Sprintln("[]")
				break
			}
			bufStr += fmt.Sprintln("[")
			for i := 0; i < value.Len(); i++ {
				bufStr += splitOne(deep+1, value.Index(i).Interface(), true)
			}
			bufStr += fmt.Sprintln(parentPrefixSpaces + "]")
		}
	case reflect.Map:
		value := reflect.ValueOf(s)
		if needIndent == true {
			bufStr += fmt.Sprint(parentPrefixSpaces)
		}
		bufStr += fmt.Sprintln("{")
		for _, v := range value.MapKeys() {
			switch v.Interface().(type) {
			case string:
				fieldName := currPrefixSpaces + (v.Interface().(string)) + ": "
				bufStr += fmt.Sprintf("%c[0;0;32m%s%c[0m", 0x1B, fieldName, 0x1B)
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				fieldName := currPrefixSpaces + fmt.Sprintf("%v", v.Interface()) + ": "
				bufStr += fmt.Sprintf("%c[0;0;32m%s%c[0m", 0x1B, fieldName, 0x1B)
			default:
				fieldName := currPrefixSpaces + v.String() + ": "
				bufStr += fmt.Sprintf("%c[0;0;32m%s%c[0m", 0x1B, fieldName, 0x1B)
			}

			if value.MapIndex(v).CanInterface() {
				bufStr += splitOne(deep+1, value.MapIndex(v).Interface(), false)
			} else {
				bufStr += fmt.Sprintln(value.MapIndex(v))
			}

		}
		bufStr += fmt.Sprintln(parentPrefixSpaces + "}")
	case reflect.String:
		if needIndent == true {
			bufStr += fmt.Sprint(currPrefixSpaces)
		}
		value := reflect.ValueOf(s)
		if value.String() == "" {
			bufStr += fmt.Sprintln(`""`)
		} else {
			bufStr += fmt.Sprintf("\"%v\"\n", value.String())
		}
	default:
		if needIndent == true {
			bufStr += fmt.Sprint(currPrefixSpaces)
		}
		bufStr += fmt.Sprintln(s)
	}
	return bufStr
}

func sliceIntType(value reflect.Value) []byte {
	buf := make([]byte, 0)

	buf = append(buf, '[')
	for i := 0; i < value.Len(); i++ {
		if i != value.Len() && i != 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, []byte(fmt.Sprintf("%v", value.Index(i)))...)
	}
	buf = append(buf, ']', '\n')

	return buf
}

func sliceStrType(value reflect.Value) []byte {
	buf := make([]byte, 0)

	buf = append(buf, '[')
	for i := 0; i < value.Len(); i++ {
		if i != value.Len() && i != 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, '"')
		buf = append(buf, []byte(fmt.Sprintf("%v", value.Index(i)))...)
		buf = append(buf, '"')
	}
	buf = append(buf, ']', '\n')

	return buf
}
