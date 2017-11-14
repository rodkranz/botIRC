package join

import (
	"fmt"
	"io"
	
	"github.com/rodkranz/botIRC/pkg/command"
)

func New() Join { return Join{} }

type Join struct{}

func (j Join) Name() string {
	return "join"
}

func (j Join) Description() string {
	return "Join in channel of server"
}

func (j Join) Run(p command.Payload, w io.Writer) {
	command.Message{
		Text: fmt.Sprintf("JOIN #%s", p.Text),
		Parse: command.ParseStyleCommand,
	}.Send(w)
}

func (j Join) Verify(command.Payload) bool {
	return false
}