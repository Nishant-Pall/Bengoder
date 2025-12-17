package main

import (
	"bufio"
	"io"
	"strconv"
)

type Value map[string]interface{}

const (
	DICT        = 'd'
	INT         = 'i'
	BYTE_STRING = "bs"
)

type Resp struct {
	reader *bufio.Reader
}

func newResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) ReadChar() (byte, error) {
	line, err := r.reader.ReadByte()
	return line, err
}

func (r *Resp) Decode() (Value, error) {

	char, err := r.ReadChar()
	if err != nil {
		return Value{}, nil
	}

	switch char {
	case DICT:
		return r.DecodeDictionary()
	}
	return Value{}, nil
}

func (r *Resp) DecodeDictionary() (Value, error) {
	value := Value{}
	key, err := r.readBaseString()

	if err != nil {
		return nil, err
	}

	peek, err := r.reader.Peek(1)

	if err != nil {
		return nil, err
	}

	if string(peek) == "i" {
		r.reader.ReadByte()
		byteArr, err := r.readInteger()

		if err != nil {
			return nil, err
		}

		intVal, err := strconv.Atoi(string(byteArr))

		value[key] = intVal
	} else if string(peek) == "d" {
		val, err := r.Decode()
		if err != nil {
			return nil, err
		}

		value[key] = val
	} else {
		val, err := r.readBaseString()
		if err != nil {
			return nil, err
		}

		value[key] = val
	}

	return value, nil
}

func (r *Resp) readBaseString() (string, error) {
	len, err := r.readLength()
	if err != nil {
		return "", err
	}
	var index = 0

	length, err := strconv.ParseInt(string(len), 10, 64)
	line := make([]byte, int(length))

	for index < int(length) {
		byt, err := r.reader.ReadByte()
		if err != nil {
			return "", err
		}
		line = append(line, byt)
		index++
	}

	return string(line), nil
}

func (r *Resp) readInteger() (integer []byte, err error) {
	for {
		char, err := r.reader.ReadByte()
		if err != nil {
			return nil, err
		}

		if string(char) == "e" {
			break
		}

		integer = append(integer, char)
	}
	return integer, nil
}

func (r *Resp) readLength() (line []byte, err error) {
	for {
		char, err := r.reader.ReadByte()
		if err != nil {
			return nil, err
		}

		if char == 58 {
			break
		}
		line = append(line, char)
	}
	return line, nil
}
