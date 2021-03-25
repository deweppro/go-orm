package plugins

import "github.com/deweppro/go-logger"

var (
	//StdOutLog simple stdout debug log
	StdOutLog logger.Logger = func() logger.Logger {
		l := logger.New()
		l.SetLevel(logger.LevelDebug)
		l.SetOutput(StdOutWriter)
		return l
	}()
)
