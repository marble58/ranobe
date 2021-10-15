package server

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Logger(path string) zerolog.Logger {
	return zerolog.New(&lumberjack.Logger{
		Filename:   path, // File name
		MaxSize:    10,   // Size in MB before file gets rotated
		MaxBackups: 5,    // Max number of files kept before being overwritten
		MaxAge:     30,   // Max number of days to keep the files
		Compress:   true, // Whether to compress log files using gzip
	})

}

