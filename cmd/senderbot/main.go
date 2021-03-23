package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"sendbot/internal/bot"
	"sendbot/internal/data"
)

var (
	botId  string
	chatId string
)

func main() {

	if botId == "" {
		log.Fatalf("boId is empty. Consider building using makefile")
	}
	if chatId == "" {
		log.Fatalf("chatId is empty. Consider building using makefile")
	}

	var (
		optSilent = flag.Bool("silent", false, "Stop sending logs to stderr")
	)
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatalf("Usage: %s [opts] <filePath>", filepath.Base(os.Args[0]))
	}
	if *optSilent {
		log.SetOutput(ioutil.Discard)
	}

	filePath := flag.Args()[0]

	if filePath == "" {
		log.Fatalf("file path is empty")
	}

	fName, r, err := data.PrepareData(filePath)
	if err != nil {
		log.Fatalf("%s", err)
	}

	sendBot := bot.Bot{
		BotId:  botId,
		ChatId: chatId,
	}
	err = sendBot.SendFile(fName, r)
	if err != nil {
		log.Fatalf("%s", err)
	}

}
