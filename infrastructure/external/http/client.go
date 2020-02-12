package http

import (
	iHttp "awawe/usecase/infrastructure/external/http"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type client struct {
	Method          string
	Protocol        string
	Host            string
	Endpoint        string
	Headers         map[string]string
	BuffData        *bytes.Buffer
	ResponseSuccess interface{}
	ResponseError   interface{}
	ClientTimeout   time.Duration
}

func NewHttpClient() iHttp.Client {
	return &client{}
}

func (h *client) SetMethod(method string) iHttp.Client {
	h.Method = method
	return h
}

func (h *client) SetProtocol(protocol string) iHttp.Client {
	h.Protocol = protocol
	return h
}

func (h *client) SetHost(host string) iHttp.Client {
	h.Host = host
	return h
}

func (h *client) SetEndpoint(endpoint string) iHttp.Client {
	h.Endpoint = endpoint
	return h
}

func (h *client) SetHeaders(headers map[string]string) iHttp.Client {
	h.Headers = headers
	return h
}

func (h *client) SetBuffData(buffData *bytes.Buffer) iHttp.Client {
	h.BuffData = buffData
	return h
}

func (h *client) SetResponse(success, error interface{}) iHttp.Client {
	h.ResponseSuccess = success
	h.ResponseError = error
	return h
}

func (h *client) Call() (*http.Response, error) {
	var data *bytes.Buffer

	if h.BuffData == nil {
		data = &bytes.Buffer{}
	} else {
		data = h.BuffData
	}

	request, err := http.NewRequest(h.Method, h.getUrlPath(), data)
	if err != nil {
		return nil, err
	}

	for key, value := range h.Headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * h.ClientTimeout,
	}

	resp, err := client.Do(request)

	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()
	return h.handleResponse(resp)
}

func (h *client) getUrlPath() string {
	return h.Protocol + "://" + h.Host + h.Endpoint
}

func (h *client) handleResponse(resp *http.Response) (*http.Response, error) {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return resp, err
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted || resp.StatusCode == http.StatusCreated {
		if h.ResponseSuccess != nil {
			err = json.Unmarshal(body, h.ResponseSuccess)
		}
	} else {
		if h.ResponseError != nil {
			err = json.Unmarshal(body, h.ResponseError)
		}
		return resp, errors.New(resp.Status)
	}

	if err != nil {
		return resp, err
	}

	return resp, nil
}
