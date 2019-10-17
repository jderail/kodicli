# Kodicli

Kodicli provide a Kodi RPC client to perform basic controls of the GUI and the player. It comes with the following functions.

* **SendTextInputRequest**

Send text to Kodi.

* **SendPlayPauseRequest**

Send a request to Play/Pause the player.

* **SendStopRequest**

Send a request to Stop the player.

* **SendInputRequest**

Interact with GUI. The following variables are provided as possible parameter.

| RPC Method        | Description                | Variable         |
|-------------------|----------------------------|------------------|
| Input.Up          | Move up in GUI             | InputUp          |
| Input.Down        | Move down in GUI           | InputDown        |
| Input.Left        | Move left in GUI           | InputLeft        |
| Input.Right       | Move right in GUI          | InputRight       |
| Input.Select      | Select current item in GUI | InputSelect      |
| Input.Home        | Go to home window in GUI   | InputHome        |
| Input.Back        | Go back in GUI             | InputBack        |
| Input.ContextMenu | Open context dialog        | InputContextMenu |

## Example

```golang
package main

import (
	"github.com/jderail/kodicli"
)

func main() {

	kodi, err := kodicli.New("192.168.1.125", 8080, "kodi", "")
	if err != nil {
		panic("Failed to initialize Kodicli")
	}
	_, err = kodi.SendPlayPauseRequest()
	if err != nil {
		panic("Failed to Play/Pause")
	}

}

```