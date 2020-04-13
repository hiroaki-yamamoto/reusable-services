package log

import (
	"log"

	"go.uber.org/zap"
)

// ZapDevLogger is a log based on zap.Logger for Development
var ZapDevLogger IZap

// ZapProdLogger is a log based on zap.Logger for Production
var ZapProdLogger IZap

func init() {
	var err error
	ZapDevLogger, err = zap.NewDevelopment()
	if err != nil {
		log.Println(
			"[WARN] failed to initialize Development Logger.",
			"Initializing ExampleLogger instead...",
		)
		ZapDevLogger = zap.NewExample()
	}

	ZapProdLogger, err = zap.NewProduction()
	if err != nil {
		log.Println(
			"[WARN] failed to initialize Production Logger.",
			"Initializing ExampleLogger instead...",
		)
		ZapProdLogger = zap.NewExample()
	}
}
