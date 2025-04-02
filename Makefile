NAME = diceware
SERVER_NAME = diceware-server

# PREFIX ?= /usr/local
PREFIX ?= $(HOME)/.local
BINDIR ?= $(PREFIX)/bin

# GO ?= $(shell which go)
GO ?= go
GOFLAGS ?=

# GOSRC := $(shell find . -name '*.go') # old
# GOSRC += go.mod go.sum
GOSRC += go.mod

RM ?= rm -f

default: $(NAME)

$(NAME): cli

cli:
	$(GO) build $(GOFLAGS) -o $(NAME) ./cmd/diceware

server:
	$(GO) build $(GOFLAGS) -o $(SERVER_NAME) ./cmd/server

install: default
	install -D -m755 $(NAME) $(DESTDIR)$(BINDIR)/$(NAME)


uninstall:
	$(RM) $(DESTDIR)$(BINDIR)/$(NAME)

clean:
	$(RM) $(NAME)

.PHONY: default $(NAME) cli server install uninstall clean
