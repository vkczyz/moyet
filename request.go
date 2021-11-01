package main

import (
	"bytes"
	"net/url"
)

type Request struct {
	url	url.URL
}

func NewRequest(data []byte) (Request, error) {
	path, err := url.Parse(string(bytes.Split(data, []byte("\r\n"))[0]))
	if err != nil {
		return Request{*path}, err
	} else {
		return Request{*path}, nil
	}
}