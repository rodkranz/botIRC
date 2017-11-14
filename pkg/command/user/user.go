package user

import (
	"fmt"
	"io"
	
	"github.com/rodkranz/botIRC/pkg/command"
)

func New() User { return User{} }

type User struct{}

func (u User) Name() string {
	return "user"
}

func (u User) Description() string {
	return "Identificate user in server"
}

func (u User) Run(p command.Payload, w io.Writer) {
	command.Message{
		Text:  fmt.Sprintf("USER %s", p.Text),
		Parse: command.ParseStyleCommand,
	}.Send(w)
}

func (u User) Verify(command.Payload) bool {
	return false
}