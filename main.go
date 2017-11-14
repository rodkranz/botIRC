package main

import (
	"fmt"
	"os"
	"runtime"
	
	log "gopkg.in/clog.v1"
	"gopkg.in/urfave/cli.v2"
	
	"github.com/rodkranz/botIRC/cmd"
	"github.com/rodkranz/botIRC/pkg/setting"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := log.New(log.CONSOLE, log.ConsoleConfig{})
	if err != nil {
		fmt.Printf("error to start logs: %v", err)
		os.Exit(1)
	}
}

func main() {
	app := cli.App{
		Name:    setting.AppName,
		Version: setting.AppVer,
		Usage:   "Bot Golang",
		Authors: []*cli.Author{
			{
				Name:  "Rodrigo Lopes",
				Email: "dev.rodrigo.lopes@gmail.com",
			},
		},
		Before: setting.Bootstrap,
		Commands: []*cli.Command{
			cmd.BotCmd,
		},
	}
	
	app.Flags = append(app.Flags, []cli.Flag{}...)
	if len(os.Args) == 1 {
		os.Args = append(os.Args, cmd.BotCmd.Name)
	}
	
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
