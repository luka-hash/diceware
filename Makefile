NAME = diceware

# PREFIX ?= /usr/local
PREFIX ?= $(HOME)/.local
BINDIR ?= $(PREFIX)/bin

# GO ?= $(shell which go)
GO ?= go
GOFLAGS ?=

GOSRC := $(shell find . -name '*.go')
# GOSRC += go.mod go.sum
GOSRC += go.mod

RM ?= rm -f

default: $(NAME)

$(NAME): $(GOSRC)
	$(GO) build $(GOFLAGS) -o $@

install: default
	install -D -m755 $(NAME) $(DESTDIR)$(BINDIR)/$(NAME)

uninstall:
	$(RM) $(DESTDIR)$(BINDIR)/$(NAME)

clean:
	$(RM) $(NAME)

.PHONY: default $(NAME) install uninstall clean