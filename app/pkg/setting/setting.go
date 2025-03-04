package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

type project struct {
	Name 	  string
	Desc 	  string
	Logo 	  string
	Developer string
}

func init() {
	var err error
	Cfg, err = ini.Load("app/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'app/conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadProject() *project {
	sec, err := Cfg.GetSection("project")
	if err != nil {
		log.Fatalf("Fail to get section 'project': %v", err)
	}
	return &project{
		Name:sec.Key("PROJECT_NAME").String(),
		Desc:sec.Key("PROJECT_DESCRIPTION").String(),
		Logo:sec.Key("PROJECT_LOGO").String(),
		Developer:sec.Key("PROJECT_DEVELOPERS").String(),
	}
}