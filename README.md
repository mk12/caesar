# Caesar

Caesar is a little program that cracks messages encrypted with the [Caesar
cipher][1] using frequency analysis. It also allows you to just encrypt messages
with a shift value or display a message's relative letter frequencies.

    $ go get github.com/mk12/caesar

[1]: http://en.wikipedia.org/wiki/Caesar_cipher

## Usage

Here is Caesar in action:

    $ caesar --help
    usage: caesar -e shift | -f | -c
    $ echo hello | caesar -e 1
    ifmmp
    $ echo ifmmp | caesar -e -1
    hello
    $ echo nothing | caesar -e a
    caesar: a: invalid syntax
    $ echo abcdefg | caesar -f
    [0.25 0.25 0.25 0.25 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    $ echo The quick brown fox jumps over the lazy dog | caesar -e $RANDOM > msg
    $ cat msg
    Vjg swkem dtqyp hqz lworu qxgt vjg ncba fqi
    $ caesar -c < msg
    The quick brown fox jumps over the lazy dog

The `cat msg` line will probably be different for you since it used a random
shift value, but 96.15% of the time it will be unintelligible (the shift value
won't be zero).

## License

Copyright © 2013 Mitchell Kember

Caesar is available under the MIT License; see [LICENSE](LICENSE.md) for
details.
