package http

import (
	"bytes"
	"net/http"
)

type Client interface {
	SetMethod(method string) Client
	SetProtocol(protocol string) Client
	SetHost(host string) Client
	SetEndpoint(endpoint string) Client
	SetHeaders(headers map[string]string) Client
	SetBuffData(buffData *bytes.Buffer) Client
	SetResponse(success, error interface{}) Client
	Call() (*http.Response, error)
}
