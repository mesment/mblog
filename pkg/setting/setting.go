package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	EnableConsole     bool
	ConsoleJsonFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJsonFormat    bool
	FileLevel         string
	FileLocation      string
	FileName          string
	DebugOpen         bool
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadLog()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RunMode").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HttpPort").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("ReadTimeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WriteTimeout").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JwtSecret").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PageSize").MustInt(10)
}

func LoadLog() {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		log.Fatalf("Fail to get section 'log': %v", err)
	}

	EnableConsole = sec.Key("EnableConsole").MustBool(false)
	ConsoleJsonFormat = sec.Key("ConsoleJsonFormat").MustBool(false)
	ConsoleLevel    = sec.Key("ConsoleLevel").MustString("info")
	EnableFile        = sec.Key("EnableFile").MustBool(false)
	FileJsonFormat    = sec.Key("FileJsonFormat").MustBool(false)
	FileLevel        = sec.Key("FileLevel").MustString("info")
	FileLocation     = sec.Key("FileLocation").MustString("./runtime/logs")
	FileName         = sec.Key("FileName").MustString("default.log")
	DebugOpen       = sec.Key("DebugOpen").MustBool(false)
}
