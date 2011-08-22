= `jsonpp` is a command line JSON pretty printer.

You'll love it.

Pretty print web service responses like so:

    curl -s -L http://t.co/tYTq5Pu | jsonpp

and make beautiful the files running around on your disk:

    jsonpp data/long_malformed.json

= Install

Grabbing the binary is the easiest way to do install `jsonpp`. It's a simple file.

To begin, you'll want to download the zip file that matches your machine:

  * [jsonpp for OSX](http://github.com/jmhodges/jsonpp/downloads/jsonpp-1.0.0-osx-x86_64.zip)
  * [jsonpp for Linux](http://github.com/jmhodges/jsonpp/downloads/jsonpp-1.0.0-linux-x86_64.zip)
      
Then, simply decompress that file and copy the `jsonpp` binary
inside to somewhere in your $PATH`. All you have left to do is to run:</p>

A source install requires a working [install of
Go](http://golang.org/doc/install.html). Once Go is installed run:

    goinstall github.com/jmhodges/jsonpp

or run `make install` yourself.
