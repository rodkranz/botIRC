package timer

import (
	"fmt"
	"io"
	"time"
	
	"github.com/rodkranz/botIRC/pkg/command"
)

func New() Timer { return Timer{} }

type Timer struct{}

func (t Timer) Name() string {
	return "timer"
}

func (t Timer) Description() string {
	return "show the time to user in channel"
}

func (t Timer) Run(p command.Payload, w io.Writer) {
	command.Message{
		Text:  fmt.Sprintf("%s %s :%s", p.Type, p.Channel, time.Now().Format("Monday Jan 2006 15:04:05")),
		Parse: command.ParseStyleCommand,
	}.Send(w)
}

func (t Timer) Verify(p command.Payload) bool {
	available := []string{
		"!time",
		"!timer",
		"!hora",
		"!horas",
	}
	
	for _, word := range available {
		if word == p.Text {
			return true
		}
	}
	
	return false
}
