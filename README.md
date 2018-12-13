
[![Build Status](https://travis-ci.org/Pirionfr/goDcCrypto.svg?branch=master)](https://travis-ci.org/Pirionfr/goDcCrypto)

goDcCrypto
==========

encrypt/decrypt lookatch agent message

Build
-----

You can also download the source code and install it manually:

    go build


Usage
-----
encrypt your message

    goDcCrypto encrypt -k <master key> -m <message>
    
decrypt your message

    goDcCrypto decrypt -k <master key> -m <message>
