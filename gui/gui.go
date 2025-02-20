// Package to capture user input and display it in a meaningful way
package gui

import (
	"fmt"
	"strings"

	"github.com/pkg/term"
)

type GUI struct {
	RedisAdress      string
	userInputHistory []string
}

func NewGUI(redisAdress string) GUI {
	return GUI{
		RedisAdress: redisAdress,
	}
}

// Starts a GUI instance
func (gui *GUI) Run() {
	gui.Render()
	for {
		key := captureUserInput()
		readableKey := string(key)
		gui.userInputHistory = append(gui.userInputHistory, readableKey)
		switch key {
		// CTRL + C = Exit
		// NOTE: This maybe should have confirmation or need to be done twice
		case ctrl_c:
			return
		case backspace:
			if len(gui.userInputHistory) > 0 {
				gui.userInputHistory = gui.userInputHistory[0 : len(gui.userInputHistory)-2]
			}
		}
		gui.Render()
	}
}

// Render should get called on each update event (key press/user action)
// and paint the UI
func (gui *GUI) Render() {
	// Code to clear the current line before each paint
	fmt.Print("\033[2K")
	fmt.Printf("\r[%s] > %s", gui.RedisAdress, strings.Join(gui.userInputHistory, ""))
}

// keycodes from raw input
var (
	ctrl_c    byte = 3
	backspace byte = 127
)

func captureUserInput() byte {
	// NOTE: I have no idea if this works on windows/linux/mac...
	t, err := term.Open("/dev/tty")
	if err != nil {
		// TODO: Not sure how to handle this error
		panic(err)
	}
	term.RawMode(t)

	// read in 3 bytes at a time
	var action int
	bytes := make([]byte, 3)
	action, _ = t.Read(bytes)

	t.Restore()
	t.Close()

	if action == 3 {
		return bytes[2]
	} else {
		return bytes[0]
	}
}
