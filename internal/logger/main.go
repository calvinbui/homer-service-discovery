package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func SetLevel(level string) error {
	l, err := zerolog.ParseLevel(strings.ToLower(level))
	if err != nil {
		return err
	}

	zerolog.SetGlobalLevel(l)

	return nil
}

func formatMsgWithErr(msg string, err error) string {
	return msg + ": " + err.Error()
}

func Trace(msg string) {
	log.Debug().Msg(msg)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Warn(msg string, err error) {
	if err != nil {
		log.Warn().Msg(formatMsgWithErr(msg, err))
	} else {
		log.Warn().Msg(msg)
	}
}

func Error(msg string, err error) {
	log.Error().Msg(formatMsgWithErr(msg, err))
}

func Fatal(msg string, err error) {
	log.Fatal().Msg(formatMsgWithErr(msg, err))
}
