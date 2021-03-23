package data

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"sendbot/internal/zip"
)

func PrepareData(filePath string) (string, io.Reader, error) {

	fInfo, err := os.Stat(filePath)
	if err != nil {
		return "", nil, fmt.Errorf("file info %s: %v", filePath, err)
	}

	if fInfo.Size() == 0 {
		return "", nil, fmt.Errorf("you cant send empty file %s: %v", filePath, err)
	}

	var r io.Reader
	if fInfo.IsDir() {
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
			return "", nil, fmt.Errorf("something wrong with open file %s  :%v", filePath, err)
		}
		defer fi.Close()
		log.Printf("sucsessed open %s ", fi.Name())
		r = fi
	}

	fName := filepath.Base(fInfo.Name())
	if fInfo.IsDir() {
		fName += ".zip"
	}

	return fName, r, nil
}
