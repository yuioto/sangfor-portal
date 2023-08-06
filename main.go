package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func checkPortal(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Warn().Msg("Request unsuccessful.")
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Warn().Msg("Reading response failed.")
		return
	}

	if response.StatusCode >= 300 && response.StatusCode < 400 && strings.Contains(string(body), "portal") {
		log.Info().Msg("Portal detection succeed.")
	} else {
		log.Warn().Msg("Portal detection fail.")
		log.Info().Msg("Network is online.")
	}
}

func main() {
	// set log style, look like: 2006-01-02T15:04:05Z07:00 INFO Msg Str
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	check_url := "https://ping.archlinux.org/nm-check.txt"

	log.Info().Msg("アトリは、高性能ですから!")

	checkPortal(check_url)
}
