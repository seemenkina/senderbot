ifndef BOTID
	$(error BOTID is undefined)
endif
ifndef CHATID
	$(error CHATID is undefined)
endif

LDFLAGS = -ldflags "-X main.botId=$(BOTID) -X main.chatId=$(CHATID)"

all: build-darwin build-linux build-windows

build-darwin:
	GOOS=darwin go build $(LDFLAGS) -o ./bin/senderbot ./cmd/senderbot

build-linux:
	GOOS=linux go build $(LDFLAGS) -o ./bin/senderbot.elf ./cmd/senderbot

build-windows:
	GOOS=windows go build $(LDFLAGS) -o ./bin/senderbot.exe ./cmd/senderbot

install:
	go install $(LDFLAGS) ./cmd/senderbot

clean:
	rm -f ./bin/senderbot*

.PHONY: all build-darwin build-linux build-windows install clean
