# natsctl
The simplest CLI, but one of the most useful ever! If you ever wanted to use NATS client within your bash :D

## Inspiration
My main inspiration was the `kubectl` command line interface, for controlling the Kubernetes cluster. So I decided to build my own for my favorite message system! [NATS](https://nats.io) is a Go based messaging system, responsible for (retirei texto) one of my best messaging experiences! It's Pretty simple, easy to configure, and the only fault was that there is no way that I can find some metrics inside my terminal, and **I DON'T LIKE WEB INTERFACES SORRY!**. Then, within a talk on the [Nerdzão Community](http://meetup.com/pt-BR/nerdzao), I'll build this tool according to what I feel is cool, but if you want to help me or put some feature that you consider cool, please [open an issue](https://github.com/Horaddrim/natsctl/issues/new) and lets discuss this!

## [Building from source](#)
Today the only way to build `natsctl` is compiling the source code for your platform!
So, you just have to grab the source code, and follow this steps, and have `go` with version equal or higher
to `v1.10`:

- `git clone https://github.com/Horaddrim/natsctl ~/.natsctl`
- `cd ~/.natsctl`
- `go build .`

And after that you will have a functional binary
inside the `~/.natsctl` folder, named `natsctl`.
If you want to access the binary everywhere else in your environment, just add the following line to your shell init script, (e.g `.bashrc` or `.zshrc`)

- `export PATH=$PATH:$HOME/.natsctl`

And you're done! :sparkling_heart: you can use the command `natsctl` from any bash or ZSH instance!

## [Available commands](#)
Today we have only the `status` command, wich you can run something like this:

- `natsctl status --host <You NATS server IP or address here>`

*OBS: The `host` flag is required*

#### TO DO
- [x] Add the documentation for building the tool from source
- [ ] Add documentation for downloading a pre-compiled version
- [ ] Add the ISSUE TEMPLATE
- [ ] Think about support Windows versions
- [ ] Upgrade the code documentation too
- [ ] Improve the `help` flag and command description
