jsonpp
======

**A command line JSON pretty printer.**

Pretty print web service responses like so:

    curl -s -L http://t.co/tYTq5Pu | jsonpp

and make beautiful the files running around on your disk:

    jsonpp data/long_malformed.json

By default jsonpp uses two spaces for indentation. Use -tabs to indent with
tabs and -spaces=N (N > 0) to indent with an arbitrary number of spaces.

Install
-------

Installable with `go get github.com/jmhodges/jsonpp`, `brew install jsonpp`,
or copying the binary to your `$PATH`. See the [live
documentation](http://jmhodges.github.com/jsonpp/) for details.

Special Note on JSON files
--------------------------

`jsonpp` assumes that there are one or more JSON objects in a given file,
seperated by newlines.

