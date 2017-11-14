package nick

import (
	"fmt"
	"io"
	
	"github.com/rodkranz/botIRC/pkg/command"
)

func New() Nick { return Nick{} }

type Nick struct{}

func (n Nick) Name() string {
	return "nick"
}

func (n Nick) Description() string {
	return "Identificate nick in server"
}

func (n Nick) Run(p command.Payload, w io.Writer) {
	command.Message{
		Text:  fmt.Sprintf("NICK %s", p.Text),
		Parse: command.ParseStyleCommand,
	}.Send(w)
}

func (n Nick) Verify(command.Payload) bool {
	return false
}
