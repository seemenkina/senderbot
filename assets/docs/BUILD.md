## How to create your bot

Detailed instructions on how to create a bot can be found at this link [Telegram Bot API](https://core.telegram.org/bots).
Further, I will outline the steps sufficient to use this functionality.

1.  Find this bot in the telegram `@BotFather` and start it.

Also, you can use this invite link [Father Bot](https://t.me/botfather)

![father-bot](./assets/img/father.png?raw=true)

2. Send `/newbot` and follow bot instruction. Like this:

![step](./assets/img/step.png?raw=true)

3. Copy the bot's token, it is required on the build step.

## How to find chat id

### First method

You can use this method if you don't want to add another bot to the group, except this of course.

1. Add your bot to your group.

2. After that, you need to request the console:

```
curl https://api.telegram.org/bot$(botToken)/getUpdates | jq
```

`$(botToken)` –– your bot's token.

The response will display the last actions with the bot, including the id of the group to which it was added:

![chat-id](./assets/img/chat-id.png?raw=true)

**Don't lose the ` - ` sign when you copy the id.**

3. You're excellent. Go to the building section.

### Second method

This method for you if you don't have access to the console. You can use any other bot you know.

**I warn you, this bot is not mine and I am not responsible for its actions.**

1. Start this bot `@username_to_id_bot` in telegram.

![other-bot](./assets/img/other-bot.png?raw=true)

2. Send the bot an invitation link to your group or channel. The response will contain the chat id.

## Building

`chat-id` –– the identifier for the target chat or username of the target channel (in the format `@channelusername`).

`bot-id` –– the identifier for your bot in the format `"X..X:X..X"`

### Build for *nix system

```
$ make build-linux BOTID="bot-id" CHATID="chat-id"
```
### Build for windows

```
$ make build-windows BOTID="bot-id" CHATID="chat-id"
```

### Build for MacOS

```
$ make build-darwin BOTID="bot-id" CHATID="chat-id"
```

### Build for all systems

```
$ make BOTID="bot-id" CHATID="chat-id"
```

## Requirements

- go1.15 or above

## Install
If you have [Golang](https://golang.org) on your system you just can do:

```
$ make install BOTID="bot-id" CHATID="chat-id"
```

## Install for all system

If necessary add the execution rights:

```
$ chmod +x senderbot
```

Copy binaries file to somewhere `$PATH`.