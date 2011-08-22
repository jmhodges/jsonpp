jsonpp
======

**A command line JSON pretty printer.**

Pretty print web service responses like so:

    curl -s -L http://t.co/tYTq5Pu | jsonpp

and make beautiful the files running around on your disk:

    jsonpp data/long_malformed.json

Install
-------

Grabbing the binary is the easiest way to do install `jsonpp`. It's a simple file.

To begin, download the zip file that matches your machine:

  * [jsonpp for OSX](http://github.com/jmhodges/jsonpp/downloads/jsonpp-1.0.0-osx-x86_64.zip)
  * [jsonpp for Linux](http://github.com/jmhodges/jsonpp/downloads/jsonpp-1.0.0-linux-x86_64.zip)
      
Then, decompress that zip file, and copy the `jsonpp` file inside to somewhere
in your `$PATH`.

To see it in action, pipe some example JSON into it:
    curl -s -L http://t.co/tYTq5Pu | jsonpp

A source install requires a working [install of
Go](http://golang.org/doc/install.html). Once Go is installed run:

    goinstall github.com/jmhodges/jsonpp

or run `make install` yourself.
