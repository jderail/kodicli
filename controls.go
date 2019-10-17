package kodicli

var (
	controlPlayPause = "Player.PlayPause"
	controlStop      = "Player.Stop"
)

// playPauseArgs structure contains informations for a Player.PlayPause RPC request
type playPauseArgs struct {
	PlayerID int `json:"playerid"`
}

// SendPlayPauseRequest sends a request to Play|Pause the player of
// the kodi configured in the RPC client
func (k *KodiRPCClient) SendPlayPauseRequest() (*int, error) {
	req := RPCRequest{
		ID:      k.nextID,
		Method:  controlPlayPause,
		Args:    playPauseArgs{PlayerID: 1},
		JSONRPC: jsonRPCVersion,
	}

	res, err := k.doRequest(&req)
	if err != nil {
		return nil, err
	}

	resMap, ok := res.Result.(map[string]interface{})
	if !ok {
		return nil, ErrUnableToParseResponse
	}

	if s, ok := resMap["speed"]; ok {
		if f64, ok := s.(float64); ok {
			speed := int(f64)
			return &speed, nil
		}
	}

	return nil, ErrUnexpectedResponse
}

// stopArgs structure contains informations for a Player.Stop RPC request
type stopArgs struct {
	PlayerID int `json:"playerid"`
}

// SendStopRequest sends a request to Stop the player of
// the kodi configured in the RPC Client
func (k *KodiRPCClient) SendStopRequest() error {
	req := RPCRequest{
		ID:      k.nextID,
		Method:  controlStop,
		Args:    stopArgs{PlayerID: 1},
		JSONRPC: jsonRPCVersion,
	}

	res, err := k.doRequest(&req)
	if err != nil {
		return err
	}

	if !checkResultOk(res.Result) {
		return ErrUnableToParseResponse
	}

	return nil
}
