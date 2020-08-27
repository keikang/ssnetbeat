package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/keikang0905/ssnetbeat/beater"
)

func main() {
	err := beat.Run("ssnetbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
