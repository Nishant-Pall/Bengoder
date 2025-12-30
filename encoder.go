package bengoder

import (
	"reflect"
	"strconv"
)

const (
	SLICE   = "slice"
	MAP     = "map"
	STRING  = "string"
	INTEGER = "int"
)

func Encode(rep any) []byte {
	t := reflect.TypeOf(rep)

	kind := t.Kind().String()

	switch kind {
	case STRING:
		return encodeString(rep)
	case INTEGER:
		return encodeInteger(rep)
	case SLICE:
		return encodeSlice(rep)
	case MAP:
		return encodeMap(rep)
	}

	return []byte{}
}

func encodeSlice(rep any) []byte {
	var bytes []byte
	bytes = append(bytes, 'l')

	for _, val := range rep.([]any) {
		bytes = append(bytes, Encode(val)...)
	}
	bytes = append(bytes, 'e')
	return bytes
}

func encodeMap(rep any) []byte {
	var bytes []byte
	bytes = append(bytes, 'd')

	for key, val := range rep.(map[string]any) {
		encodedKey := Encode(key)
		encodedValue := Encode(val)
		bytes = append(bytes, encodedKey...)
		bytes = append(bytes, encodedValue...)
	}
	bytes = append(bytes, 'e')
	return bytes
}

func encodeString(rep any) []byte {
	var bytes []byte
	len := len(rep.(string))

	bytes = append(bytes, []byte(strconv.Itoa(len))...)
	bytes = append(bytes, ':')
	bytes = append(bytes, []byte(rep.(string))...)
	return bytes
}

func encodeInteger(rep any) []byte {
	var bytes []byte

	bytes = append(bytes, 'i')
	bytes = append(bytes, []byte(strconv.Itoa(rep.(int)))...)
	bytes = append(bytes, 'e')
	return bytes
}
