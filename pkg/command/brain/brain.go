package brain

import (
	"fmt"
	"io"
	
	"github.com/aichaos/rivescript-go"
	log "gopkg.in/clog.v1"
	
	"github.com/rodkranz/botIRC/pkg/command"
	"github.com/rodkranz/botIRC/pkg/setting"
)

type Brain struct {
	brubs *rivescript.RiveScript
}

func New(path string) Brain {
	brubs := rivescript.New(rivescript.WithUTF8())
	
	err := brubs.LoadDirectory(path)
	if err != nil {
		panic(err)
	}
	
	brubs.SortReplies()
	
	return Brain{
		brubs: brubs,
	}
}

func (b Brain) Name() string {
	return "brain"
}

func (b Brain) Description() string {
	return "AI for bot"
}

func (b Brain) Run(p command.Payload, w io.Writer) {
	if p.Text == "" {
		return
	}
	
	reply, err := b.brubs.Reply(setting.Auth.Nickname, p.Text)
	if err != nil {
		log.Warn("Brubs broken: %s", err.Error())
		return
	}
	
	command.Message{
		Text:  fmt.Sprintf("%s %s :%s", p.Type, p.Username, reply),
		Parse: command.ParseStyleCommand,
	}.Send(w)
}

func (b Brain) Verify(p command.Payload) bool {
	return p.Channel == setting.Auth.Nickname && p.Username != setting.Auth.Nickname
}
