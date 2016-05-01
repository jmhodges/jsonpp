jsonpp
======

**A command line JSON pretty printer.**

Pretty print web service responses like so:

    curl -s -L http://t.co/tYTq5Pu | jsonpp

and make beautiful the files running around on your disk:

    jsonpp testdata/multiple/multiple.json

You can also format previously pretty-printed code with "-s":

    jsonpp -s testdata/one/singular.json

Install
-------

Installable with `go get github.com/jmhodges/jsonpp`, `brew install jsonpp`,
or copying the binary to your `$PATH`. See the [live
documentation](http://jmhodges.github.com/jsonpp/) for details.

Options
-------

The string used for indentation defaults to 2 spaces, but can be overridden
by the environment variable `JSONPP_INDENT`.

Adding the "-s" parameter will allow you to format already formatted JSON code by assuming the entire input stream is one JSON object.
