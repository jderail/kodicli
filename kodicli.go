package kodicli

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const jsonRPCVersion = "2.0"

// ErrUnableToParseResponse error is returned when RPC client is unable
// to parse response from Kodi
var ErrUnableToParseResponse = errors.New("Unable to parse Kodi response")

// ErrUnexpectedResponse error is returned when the response received
// does not contains expected content
var ErrUnexpectedResponse = errors.New("Unexpected response from Kodi")

// KodiRPCClient structure contains informations necessary to
// contact Kodi HTTP server
type KodiRPCClient struct {
	endpoint string
	auth     string
	nextID   int
}

// RPCRequest structure represents RPC request object
// according to JSON-RPC 2.0 specifications
type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Args    interface{} `json:"params,omitempty"`
	ID      int         `json:"id"`
}

// RPCResponse structure contains RPC response object
// according to JSON-RPC 2.0 specifications
type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Error   *RPCError   `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// RPCError structure contains informations about an RPC error object
// according to JSON-RPC 2.0 specifications
type RPCError struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

func checkResultOk(result interface{}) bool {
	asStr, ok := result.(string)
	if !ok {
		return false
	}
	return asStr == "OK"
}

// New initialize a new Kodi RPC Client
func New(host string, port int, user string, password string) (*KodiRPCClient, error) {

	endpoint, err := url.Parse("http://" + host + ":" + strconv.Itoa(port) + "/jsonrpc")
	if err != nil {
		return nil, errors.New("Unable to create kodi URL : " + err.Error())
	}

	auth := user + ":" + password
	auth = base64.StdEncoding.EncodeToString([]byte(auth))

	return &KodiRPCClient{
		endpoint: endpoint.String(),
		auth:     auth,
		nextID:   0,
	}, nil

}

func (k *KodiRPCClient) doRequest(rpc *RPCRequest) (*RPCResponse, error) {
	asJSON, err := json.Marshal(rpc)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, k.endpoint, bytes.NewBuffer(asJSON))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Basic "+k.auth)
	req.Header.Set("User-Agent", "kodicli")
	cli := http.Client{Timeout: time.Second * 20}
	response, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	var result RPCResponse
	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, ErrUnableToParseResponse
	}

	if result.Error != nil {
		return nil, errors.New("Operation failed with message : " + result.Error.Message)
	}

	return &result, nil

}
