package main

import (
	"strconv"
)

type Response struct {
	status	int
	meta	string
	body	[]byte
}

func NewResponse(status int, meta string) Response {
	response := Response{status, meta, []byte("")}
	return response
}

func (response *Response) AddBodyFromBytes(data []byte) {
	response.body = data
}

func (response *Response) AddBodyFromString(data string) {
	response.body = []byte(data)
}

func (response Response) Format() []byte {
	output := []byte(strconv.Itoa(response.status) + " " + response.meta + "\r\n")
	output = append(output, response.body...)
	return output
}