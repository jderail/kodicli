package kodicli

const (
	// InputBack RPC method to go back in GUI
	InputBack = "Input.Back"
	// InputContextMenu RPC method to show context menu
	InputContextMenu = "Input.ContextMenu"
	// InputDown RPC method to navigate down in GUI
	InputDown = "Input.Down"
	// InputUp RPC method to navigate up in GUI
	InputUp = "Input.Up"
	// InputLeft RPC method to navigate left in GUI
	InputLeft = "Input.Left"
	// InputRight RPC method to navigate right in GUI
	InputRight = "Input.Right"
	// InputHome RPC method to go home in GUI
	InputHome = "Input.Home"
	// InputInfo RPC method to show info dialog
	InputInfo = "Input.Info"
	// InputSelect RPC method to select current item in GUI
	InputSelect = "Input.Select"

	inputSendText = "Input.SendText"
)

type sendTextInputRequest struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// SendInputRequest sends an input request to trigger actions in GUI
func (k *KodiRPCClient) SendInputRequest(i string) error {
	req := RPCRequest{
		ID:      k.nextID,
		Method:  i,
		Args:    nil,
		JSONRPC: jsonRPCVersion,
	}

	res, err := k.doRequest(&req)
	if err != nil {
		return err
	}

	if !checkResultOk(res.Result) {
		return ErrUnexpectedResponse
	}

	return nil
}

// SendTextInputRequest sends RPC request to input text
func (k *KodiRPCClient) SendTextInputRequest(text string) error {
	req := RPCRequest{
		ID:     k.nextID,
		Method: inputSendText,
		Args: sendTextInputRequest{
			Text: text,
			Done: true,
		},
		JSONRPC: jsonRPCVersion,
	}

	res, err := k.doRequest(&req)
	if err != nil {
		return err
	}

	if !checkResultOk(res.Result) {
		return ErrUnexpectedResponse
	}

	return nil
}
