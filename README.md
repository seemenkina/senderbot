# senderbot

`senderbot` is a simple CLI tool for sending files or directories 
to your telegram chat or channel using one simple command 
(and no config files!).

Imagine you want to send a log file to your teammate:

```bash
$ senderbot /var/log/nginx/access.log
```

And voila:

![team-log](./assets/img/team-log.png?raw=true)

## FAQ

### Can I take your binary and use it in my chat?

Not yet. By design, I wanted it to be a single binary file without configs.
But this is in the plans.

### How can I make a bot for myself?

Look at this instruction manual : [BUILD.md](./assets/docs/BUILD.md)
