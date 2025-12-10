package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	DICT = 'd'
	END  = 'e'
)

type Value struct {
	typ   string
	key   string
	bulk  string
	value []Value
}

type Resp struct {
	reader *bufio.Reader
}

func newResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) ReadLine() (byte, error) {
	line, err := r.reader.ReadByte()
	return line, err
}

func (r *Resp) Read() {

	line, err := r.ReadLine()
	if err != nil {
		return
	}

	switch line {
	case DICT:
		r.DecodeDictionary()
		return
	}

}

func (r *Resp) DecodeDictionary() {
	v := Value{}
	v.typ = string(DICT)

	for {
		key, err := r.readLine()
		if err != nil {
			return
		}
		value, err := r.readLine()
		if err != nil {
			return
		}

		val := Value{}

		val.typ = string(DICT)
		val.key = key
		val.bulk = value

		v.value = append(v.value, val)
		fmt.Println(v)
	}
}

func (r *Resp) readLine() (string, error) {
	len, err := r.readInteger()
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

func (r *Resp) readInteger() (line []byte, err error) {
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
