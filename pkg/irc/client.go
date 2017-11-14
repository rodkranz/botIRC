package irc

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"strings"

	log "gopkg.in/clog.v1"

	"github.com/rodkranz/botIRC/pkg/command"
	"github.com/rodkranz/botIRC/pkg/setting"
)

type Handler interface {
	Name() string
	Description() string
	Run(command.Payload, io.Writer)
	Verify(command.Payload) bool
}

type Client struct {
	Conn       net.Conn
	Handlers   map[string][]Handler
	Middleware func(ln string) string
}

func New() *Client {
	return &Client{
		Handlers: make(map[string][]Handler),
	}
}

func (c *Client) AddHandler(commands ...Handler) {
	for _, cmd := range commands {
		cmdName := strings.ToUpper(cmd.Name())
		c.Handlers[cmdName] = append(c.Handlers[cmdName], cmd)
		log.Info("Command %s registered successfully", cmdName)
	}
}

func (c *Client) Execute(command string, p command.Payload) error {
	commands, found := c.Handlers[strings.ToUpper(command)]
	if !found {
		log.Warn("Command %s not found", command)
		return fmt.Errorf("no command was found :(")
	}

	for _, cmd := range commands {
		cmd.Run(p, c.Conn)
	}

	return nil
}

func (c *Client) Connect() (err error) {
	dial := fmt.Sprintf("%s:%d", setting.Server.Host, setting.Server.Port)

	c.Conn, err = net.Dial("tcp", dial)
	if err != nil {
		log.Warn("Unable to connect to IRC Server: %s", err.Error())
		return fmt.Errorf("unable to connect to IRC Server: %s", err.Error())
	}

	return nil
}

func (c *Client) Listen() error {
	defer c.Conn.Close()

	reader := bufio.NewReader(c.Conn)
	tp := textproto.NewReader(reader)

	quit := make(chan error, 1)
	go func() {
		c.readLine(tp, quit)
	}()

	return <-quit
}

func (c *Client) readLine(r *textproto.Reader, quit chan error) {
	for {
		line, err := r.ReadLine()
		if err != nil {
			log.Warn("Error to parse line", err)
			quit <- err
			return
		}
		line = c.Middleware(line)

		log.Trace("-> %s", line)
		pl := command.ParseLine(line)

		log.Info("## %#v", pl)

		for _, handler := range c.Handlers {
			for _, h := range handler {
				if ok := h.Verify(pl); ok {
					h.Run(pl, c.Conn)
				}
			}
		}
	}
	close(quit)
}
