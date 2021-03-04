package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"sendbot/internal/bot"
	"sendbot/internal/zip"
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
		log.Fatalf("filePath is empty. Consider building using makefile")
	}

	finfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("File Info %s: %v", filePath, err)
	}

	if finfo.Size() == 0 {
		log.Fatalf("You cant send empty file %s: %v", filePath, err)
	}

	var r io.Reader
	if finfo.IsDir() {
		// make zip from directory
		pipeReader, pipeWriter := io.Pipe()
		go func() {
			_ = pipeWriter.CloseWithError(zip.Make(pipeWriter, filePath))
		}()
		r = pipeReader
	} else {
		// open file
		log.Printf("try open %s ", filePath)
		fi, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("something wrong with open file %s  :%v", filePath, err)
		}
		defer fi.Close()
		log.Printf("sucsessed open %s ", fi.Name())
		r = fi
	}

	fName := filepath.Base(finfo.Name())
	if finfo.IsDir() {
		fName += ".zip"
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
