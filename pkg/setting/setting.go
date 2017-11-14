package setting

import (
	"fmt"
	
	"gopkg.in/ini.v1"
	"gopkg.in/urfave/cli.v2"
)

var (
	AppName string
	AppVer  string
	
	Auth struct {
		Username string
		Nickname string
		Password string
	}
	
	Server struct {
		Host     string
		Port     int
		Channels []string
	}
	
	AI struct {
		Brain string
	}
	
	LogFile string
	
	Cfg *ini.File
)

func Bootstrap(_ *cli.Context) error {
	//B Load configurations
	Cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Fatal(0, "Fail to parse 'config/app.ini': %v", err)
		return err
	}
	
	// Auth configurations
	sec := Cfg.Section("AUTHENTICATION")
	Auth.Username = sec.Key("username").MustString("GoLangBot")
	Auth.Nickname = sec.Key("nickname").MustString("gobot")
	Auth.Password = sec.Key("password").MustString("")
	
	// Server configurations
	sec = Cfg.Section("SERVER")
	Server.Host = sec.Key("hostname").MustString("irc.freenode.net")
	Server.Port = sec.Key("port").MustInt(6667)
	scs := sec.Key("channels").Strings(",")
	Server.Channels = make([]string, 0, len(scs))
	for _, c := range scs {
		Server.Channels = append(Server.Channels, c)
	}
	
	AI.Brain = Cfg.Section("AI").Key("brain").MustString("brain")
	
	logFile := Cfg.Section("LOG").Key("output").MustString("log/log_%s.log")
	LogFile = fmt.Sprintf(logFile, Server.Host)
	
	return nil
}
