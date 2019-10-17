package kodicli

import (
	"testing"
)

func TestSendInputRequest(t *testing.T) {
	host, port := initHTTP(t, "OK", func(req *RPCRequest) bool {
		if req.Method != InputHome {
			return false
		}
		if req.Args != nil {
			return false
		}
		return true
	})

	kodi, err := New(host, port, "kodi", "")
	if err != nil {
		t.Fatalf("Cannot initialize kodicli with error : %v", err.Error())
	}

	err = kodi.SendInputRequest(InputHome)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestSendTextInputRequest(t *testing.T) {
	const testText = "test test"

	host, port := initHTTP(t, "OK", func(req *RPCRequest) bool {

		if req.Method != inputSendText {
			return false
		}

		args, ok := req.Args.(map[string]interface{})
		if !ok {
			return false
		}
		done, ok := args["done"]
		if !ok {
			return false
		}
		if done, ok := done.(bool); !ok || !done {
			return false
		}

		text, ok := args["text"]
		if !ok {
			return false
		}
		if text, ok := text.(string); !ok || text != testText {
			return false
		}

		return true
	})

	kodi, err := New(host, port, "kodi", "")
	if err != nil {
		t.Fatalf("Cannot initialize kodicli with error : %v", err.Error())
	}

	err = kodi.SendTextInputRequest(testText)
	if err != nil {
		t.Fatal(err.Error())
	}
}
