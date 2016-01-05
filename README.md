# Buskin

> a thick-soled laced boot worn by an ancient Athenian tragic actor to gain height

This is a 'bootstrap' script of sorts, built as a executable binary using [Go](https://golang.org/) instead of your traditional Bash scripts.

## Usage

```bash
buskin path/to/config
```

> Note: the config path can be local or remote

## Configuration

The binary is controlled by a [yaml](http://www.yaml.org/) configuration file.

The following is an example:

```yaml
github:
  - username: foobar
  - email: foo@bar.com
homebrew:
  - neovim
  - git
  - rbenv
  - ruby-build 
  - irssi 
  - leiningen 
  - reattach-to-user-namespace 
  - siege 
  - terminal-notifier 
  - the_silver_searcher 
  - tmux 
  - tree 
  - gpg 
  - watch 
  - go
shell:
  - type: zsh
dotfiles:
  - path: https://github.com/integralist/dotfiles
go:
  - golang.org/x/tools/cmd/goimports
  - github.com/svent/sift
```
