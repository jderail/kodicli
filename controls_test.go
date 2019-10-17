package kodicli

import (
	"testing"
)

type playPauseResponse struct {
	Speed int `json:"speed"`
}

func TestSendPlayPauseRequest(t *testing.T) {
	mockedResponse := playPauseResponse{Speed: 1}
	host, port := initHTTP(t, mockedResponse, func(req *RPCRequest) bool {
		if req.Method != controlPlayPause {
			return false
		}
		args, ok := req.Args.(map[string]interface{})
		if !ok {
			return false
		}
		playerID, ok := args["playerid"]
		if !ok {
			return false
		}
		if playerID, ok := playerID.(float64); !ok || playerID != 1 {
			return false
		}
		return true
	})

	kodi, err := New(host, port, "kodi", "")
	if err != nil {
		t.Fatalf("Cannot initialize kodicli with error : %v", err.Error())
	}

	speed, err := kodi.SendPlayPauseRequest()
	if err != nil {
		t.Fatal(err.Error())
	}

	if *speed != 1 {
		t.Fatalf("Expected returned speed to be 1 but got %v", speed)
	}
}

func TestSendStopRequest(t *testing.T) {

	host, port := initHTTP(t, "OK", func(req *RPCRequest) bool {

		if req.Method != controlStop {
			return false
		}

		args, ok := req.Args.(map[string]interface{})
		if !ok {
			return false
		}
		playerID, ok := args["playerid"]
		if !ok {
			return false
		}
		if playerID, ok := playerID.(float64); !ok || playerID != 1 {
			return false
		}

		return true
	})

	kodi, err := New(host, port, "kodi", "")
	if err != nil {
		t.Fatalf("Cannot initialize kodicli with error : %v", err.Error())
	}

	err = kodi.SendStopRequest()
	if err != nil {
		t.Fatal(err.Error())
	}
}
