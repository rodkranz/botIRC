package command

import (
	"io"

	log "gopkg.in/clog.v1"
)

type ParseStyle string

var (
	ParseStyleCommand = ParseStyle("command")
)

type Message struct {
	Channel  string
	Username string
	Text     string
	Type     string
	Parse    ParseStyle
}

func (m Message) String() string {
	return m.Text + "\r\n"
}

func (m Message) Bytes() []byte {
	return []byte(m.String())
}

func (m Message) Send(writer io.Writer) {
	log.Trace("<- %s", m.Text)
	writer.Write(m.Bytes())
}
