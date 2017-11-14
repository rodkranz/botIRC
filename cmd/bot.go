package cmd

import (
	"fmt"
	"os"
	"time"
	
	"gopkg.in/urfave/cli.v2"
	
	"github.com/rodkranz/botIRC/pkg/command"
	"github.com/rodkranz/botIRC/pkg/command/brain"
	"github.com/rodkranz/botIRC/pkg/command/join"
	"github.com/rodkranz/botIRC/pkg/command/nick"
	"github.com/rodkranz/botIRC/pkg/command/ping"
	"github.com/rodkranz/botIRC/pkg/command/timer"
	"github.com/rodkranz/botIRC/pkg/command/user"
	"github.com/rodkranz/botIRC/pkg/irc"
	"github.com/rodkranz/botIRC/pkg/setting"
)

var BotCmd = &cli.Command{
	Name:   "bot",
	Usage:  "start bot",
	Action: runBot,
}

func runBot(c *cli.Context) error {
	// Create Log file
	f, err := os.Create(setting.LogFile)
	if err != nil {
		return err
	}
	defer f.Close()
	
	// Start new client
	mIrc := irc.New()
	if err := mIrc.Connect(); err != nil {
		return err
	}
	
	// Register commands
	mIrc.AddHandler(
		nick.New(), user.New(), join.New(), ping.New(),
		timer.New(), brain.New(setting.AI.Brain),
	)
	
	// Execute commands hardcore
	mIrc.Execute("nick", command.Payload{Text: setting.Auth.Nickname})
	mIrc.Execute("user", command.Payload{Text: fmt.Sprintf("%s 8 * :%s", setting.Auth.Username, setting.Auth.Nickname)})
	
	// Set five seconds to enter in channels
	time.AfterFunc(5*time.Second, func() {
		for _, chn := range setting.Server.Channels {
			mIrc.Execute("join", command.Payload{Text: chn})
		}
	})
	
	// Middleware running in each line that server sends to us
	mIrc.Middleware = func(ln string) string {
		return ln
	}
	
	// Listen commands from server.
	if err := mIrc.Listen(); err != nil {
		return err
	}
	
	return nil
}
