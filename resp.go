package main

import (
	"bufio"
	"io"
	"strconv"
)

type Value map[string]any

const (
	DICT = "d"
	INT  = "i"
	LIST = "l"
)

type Resp struct {
	reader *bufio.Reader
}

func newResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) PeekChar() (string, error) {
	line, err := r.reader.Peek(1)
	return string(line), err
}

func (r *Resp) Decode() (any, error) {
	char, err := r.PeekChar()
	if err != nil {
		return Value{}, nil
	}

	switch char {
	case LIST:
		return r.decodeList()
	case DICT:
		return r.decodeDictionary()
	case INT:
		return r.readInteger()
	default:
		return r.readBaseString()
	}
}

func (r *Resp) decodeList() ([]any, error) {
	r.reader.ReadByte()

	var value []any
	for {
		val, err := r.Decode()
		if err != nil {
			return nil, err
		}

		value = append(value, val)

		peek, err := r.reader.Peek(1)
		if err != nil {
			return nil, err
		}

		if string(peek) == "e" {
			r.reader.ReadByte()
			break
		}
	}
	return value, nil
}

func (r *Resp) decodeDictionary() (Value, error) {
	value := Value{}

	r.reader.ReadByte()
	for {
		peek, err := r.reader.Peek(1)
		if err != nil {
			return nil, err
		}

		if string(peek) == "e" {
			r.reader.ReadByte()
			break
		}

		key, err := r.readBaseString()
		if err != nil {
			return nil, err
		}

		val, err := r.Decode()
		if err != nil {
			return nil, err
		}
		value[key] = val
	}
	return value, nil
}

func (r *Resp) readBaseString() (string, error) {
	length, err := r.readLength()
	if err != nil {
		return "", err
	}

	line := make([]byte, int(length))

	io.ReadFull(r.reader, line)

	return string(line), nil
}

func (r *Resp) readInteger() (int, error) {
	r.reader.ReadByte()

	byteArr, err := r.readUntilDelim(byte('e'))
	if err != nil {
		return 0, nil
	}

	intVal, err := strconv.Atoi(string(byteArr))
	if err != nil {
		return 0, err
	}

	return intVal, nil
}

func (r *Resp) readLength() (int64, error) {
	line, err := r.readUntilDelim(byte(':'))

	if err != nil {
		return 0, err
	}

	length, err := strconv.ParseInt(string(line), 10, 64)

	if err != nil {
		return 0, err
	}

	return length, nil
}

func (r *Resp) readUntilDelim(delimiter byte) ([]byte, error) {
	byteArr, err := r.reader.ReadBytes(delimiter)

	if err != nil {
		return nil, err
	}

	if len(byteArr) > 0 && byteArr[len(byteArr)-1] == delimiter {
		byteArr = byteArr[:len(byteArr)-1]
	}

	return byteArr, nil
}
