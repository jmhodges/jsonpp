jsonpp
======

**A command line JSON pretty printer.**

Pretty print web service responses like so:

    curl -s -L http://t.co/tYTq5Pu | jsonpp

and make beautiful the files running around on your disk:

    jsonpp data/long_malformed.json

Install
-------

Installable with `go get github.com/jmhodges/jsonpp`, `brew install jsonpp`,
or copying the binary to your `$PATH`. See the [live
documentation](http://jmhodges.github.com/jsonpp/) for details.

Options
-------

The string used for indentation defaults to 2 spaces, but can be overridden
by the environment variable `JSONPP_INDENT`.

Special Note on JSON files
--------------------------

`jsonpp` assumes that there are one or more JSON objects in a given file,
seperated by newlines.

