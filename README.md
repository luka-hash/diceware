# diceware

Quick and dirty Go program to generate passphrases.

## Usage

The default length of the passphrase is 8 words, but you can pass a number to denote the number of words to be generated.

```
$> diceware
overflow revenge estate visitor confidant google junction spoof
$> diceware 5
cadmium nemeses maturely crazily squid
```

## Is it any good?

Yes.

## TODO
- [x] embed wordlist into a program, so it can actually be installed

## Licence

The word list: CC BY 3.0 US 2016 Electronic Frontier Foundation https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases

This code is licensed under MIT licence (see LICENCE for details).
