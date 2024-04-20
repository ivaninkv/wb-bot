package logger

import (
	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog"
	"log/slog"
	"os"
)

var (
	hostname, _ = os.Hostname()
	zlog        = zerolog.New(os.Stdout).
			With().
			Str("host", hostname).
			Logger()

	Log = slog.New(
		slogzerolog.Option{Logger: &zlog}.NewZerologHandler(),
	)
)
