package ping

import (
	"fmt"
	"io"
	
	"github.com/rodkranz/botIRC/pkg/command"
)

func New() Ping { return Ping{} }

type Ping struct{}

func (p Ping) Name() string {
	return "ping"
}

func (p Ping) Description() string {
	return "Ping in channel of server"
}

func (p Ping) Run(pl command.Payload, w io.Writer) {
	command.Message{
		Text:  fmt.Sprintf("JOIN #%s", pl.Text),
		Parse: command.ParseStyleCommand,
	}.Send(w)
}

func (p Ping) Verify(pl command.Payload) bool {
	return pl.TriggerWord == "PING"
}
