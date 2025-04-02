# diceware

Passphrase generator written in Go.

## Usage

The default length of the passphrase is 8 words, but you can pass a number to
specify the number of words to be generated.

``` sh
$ diceware
overflow revenge estate visitor confidant google junction spoof
$ diceware 5
cadmium nemeses maturely crazily squid
```

## Installation

```sh
$ git clone github.com/luka-hash/diceware
$ make # to build and try it
$ make install  # to install it in your local bin dir
$ make uninstall # to uninstall it
$ sudo make install PREFIX=/usr/local # to install it system-wide
$ sudo make uninstall PREFIX=/usr/local # to uninstall it if installed system-wide
```

Also, you can set DESTDIR to install it into a different system root, e.g. use
make install DESTDIR="$pkgdir" PREFIX=/usr on Arch.

## Is it any good?

Yes.

## TODO
- [ ] create non-english wordlist and add a flag to switch the wordlist

## Licence

This code is licensed under the terms of the MIT licence (see LICENCE for details).
