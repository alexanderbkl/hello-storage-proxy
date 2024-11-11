package event

import "github.com/sirupsen/logrus"

// TextFormatter for log messages.
var TextFormatter = &logrus.TextFormatter{
	DisableColors: false,
	FullTimestamp: true,
}

// Log is the global default logger.
var Log = logrus.New()

func init() {
	Log.SetFormatter(TextFormatter)
	//level is set to debug
	//Log.SetLevel(logrus.DebugLevel)
}
