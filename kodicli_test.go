package kodicli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestCheckResultOk(t *testing.T) {

	var ok interface{} = "OK"
	if !checkResultOk(ok) {
		t.Error("OK test failed should have returned true")
	}

	var nok interface{} = "NOK"
	if checkResultOk(nok) {
		t.Error("NOK test failed should have returned false")
	}
}

func respondErrorMessage(message string, w http.ResponseWriter) {
	res := RPCResponse{
		JSONRPC: jsonRPCVersion,
		ID:      0,
		Error: &RPCError{
			Code:    -1,
			Message: message,
			Data:    nil,
		},
	}

	b, _ := json.Marshal(res)
	w.Write(b)
}

func initHTTP(t *testing.T, result interface{}, valid func(req *RPCRequest) bool) (host string, port int) {

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, passwd, ok := r.BasicAuth()
		if !ok || user != "kodi" || passwd != "" {
			respondErrorMessage("Failed to authenticate client request", w)
			return
		}

		defer r.Body.Close()
		rb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondErrorMessage("Failed to read request body", w)
			return
		}
		var req RPCRequest
		json.Unmarshal(rb, &req)
		if !valid(&req) {
			respondErrorMessage("Request does not match validation", w)
			return
		}

		res := RPCResponse{
			JSONRPC: jsonRPCVersion,
			Result:  result,
			ID:      0,
			Error:   nil,
		}
		b, _ := json.Marshal(res)

		w.Write(b)
	}))

	addr, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatal("Error occured while initializing test server")
	}

	host = addr.Hostname()
	port, err = strconv.Atoi(addr.Port())
	if err != nil {
		t.Fatal("Error occured while initializing test server")
	}

	return
}
