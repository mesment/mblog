package logger

import (
	"github.com/mesment/mblog/pkg/setting"
	"github.com/mesment/sparrow/pkg/xlog"
)

var instance  *xlog.Logger
func Init() *xlog.Logger {
	conf := xlog.Config{
		EnableConsole:setting.EnableConsole,
		ConsoleJSONFormat:setting.ConsoleJsonFormat,
		ConsoleLevel: setting.ConsoleLevel,
		EnableFile:setting.EnableFile,
		FileJSONFormat:setting.FileJsonFormat,
		Dir:  setting.FileLocation,
		Name: setting.FileName,
		Level: setting.FileLevel,
		AddCaller:true,
		CallerSkip:1,
		Debug:setting.DebugOpen,
	}
	instance = conf.Build()
	return instance
}

func NewLogger() *xlog.Logger {
	return instance
}