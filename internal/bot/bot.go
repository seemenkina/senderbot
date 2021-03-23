package bot

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

var client = http.Client{}

type Bot struct {
	BotId  string
	ChatId string
}

func (b *Bot) ConstructUrlForSend() string {
	return "https://api.telegram.org/bot" + b.BotId + "/sendDocument"
}

// SendFile sends the document with fileName using the bot. The sendDocument Telegram API method is used
func (b *Bot) SendFile(fileName string, reader io.Reader) error {

	pipeReader, pipeWriter := io.Pipe()
	multipartWriter := multipart.NewWriter(pipeWriter)
	go func() {
		defer pipeWriter.Close()

		// start write field in multipartWriter
		if err := multipartWriter.WriteField("chat_id", b.ChatId); err != nil {
			_ = pipeWriter.CloseWithError(err)
			return
		}

		const disableNotification = "true"
		if err := multipartWriter.WriteField("disable_notification", disableNotification); err != nil {
			_ = pipeWriter.CloseWithError(err)
			return
		}

		part, err := multipartWriter.CreateFormFile("document", fileName)
		if err != nil {
			_ = pipeWriter.CloseWithError(err)
			return
		}

		log.Printf("[send file] start copy %s", fileName)
		if _, err := io.Copy(part, reader); err != nil {
			_ = pipeWriter.CloseWithError(err)
			return
		}
		log.Printf("[send file] sucess copy %s", fileName)

		if err := multipartWriter.Close(); err != nil {
			_ = pipeWriter.CloseWithError(err)
			return
		}

	}()

	resp, err := client.Post(b.ConstructUrlForSend(), multipartWriter.FormDataContentType(), pipeReader)
	if err != nil {
		_ = pipeReader.CloseWithError(err)
		return err
	}

	resp.Close = true
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[send file] bad HTTP response: %v, %s", resp.StatusCode, data)
	}
	log.Printf("[send file] sucess send file %s", fileName)

	return nil
}
