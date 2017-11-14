package command

import (
	"regexp"
	"strings"
)

type Payload struct {
	Username    string
	Channel     string
	Text        string
	TriggerWord string
	BotName     string
	Robot       string
	Type        string
}

func ParseLine(line string) (pl Payload) {
	if len(line) == 0 {
		return
	}
	
	switch {
	case strings.HasPrefix(line, ":"):
		regexLn := regexp.MustCompile(`\:([^\s]+)[\ ]([^\s]+)[\ ]([^\s]+)\ \:(.*)`)
		cutLine := regexLn.FindStringSubmatch(line)
		if len(cutLine) >= 5 {
			pl.Username = cutLine[1]
			pl.Type = cutLine[2]
			pl.Channel = cutLine[3]
			pl.Text = cutLine[4]
			
			regServer := regexp.MustCompile(`(.*)!~`)
			cutServer := regServer.FindStringSubmatch(pl.Username)
			if len(cutServer) >= 1 {
				pl.Username = cutServer[1]
			}
			
			if cutLine[4][0:1] == "!" {
				regCmd := regexp.MustCompile(`([^\s]+)[\ ](.*)`)
				cutCmd := regCmd.FindStringSubmatch(pl.Text[1:])
				if len(cutCmd) >= 3 {
					pl.TriggerWord = strings.ToUpper(cutCmd[1])
					pl.Text = cutCmd[2]
				}
			}
		}
		break
	default:
		regexLn := regexp.MustCompile(`(.*)\ \:(.*)`)
		cutLine := regexLn.FindStringSubmatch(line)
		if len(cutLine) >= 3 {
			pl.TriggerWord = strings.ToUpper(cutLine[1])
			pl.Text = cutLine[2]
		}
	}
	
	return pl
}
