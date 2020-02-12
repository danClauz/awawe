package mock

import (
	iHttp "awawe/usecase/infrastructure/external/http"
	"bytes"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type HttpClientMock struct {
	mock.Mock
}

func (m *HttpClientMock) SetMethod(method string) iHttp.Client {
	args := m.Called(method)
	return args.Get(0).(iHttp.Client)
}

func (m *HttpClientMock) SetProtocol(protocol string) iHttp.Client {
	args := m.Called(protocol)
	return args.Get(0).(iHttp.Client)
}

func (m *HttpClientMock) SetHost(host string) iHttp.Client {
	args := m.Called(host)
	return args.Get(0).(iHttp.Client)
}

func (m *HttpClientMock) SetEndpoint(endpoint string) iHttp.Client {
	args := m.Called(endpoint)
	return args.Get(0).(iHttp.Client)
}

func (m *HttpClientMock) SetHeaders(headers map[string]string) iHttp.Client {
	args := m.Called(headers)
	return args.Get(0).(iHttp.Client)
}

func (m *HttpClientMock) SetBuffData(buffData *bytes.Buffer) iHttp.Client {
	args := m.Called(buffData)
	return args.Get(0).(iHttp.Client)
}

func (m *HttpClientMock) SetResponse(success, error interface{}) iHttp.Client {
	args := m.Called(success, error)
	return args.Get(0).(iHttp.Client)
}

func (m *HttpClientMock) Call() (*http.Response, error) {
	args := m.Called()
	return args.Get(0).(*http.Response), args.Error(1)
}
