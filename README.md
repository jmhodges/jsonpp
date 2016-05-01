jsonpp
======

**A command line JSON pretty printer.**

Pretty print web service responses like so:

    curl -s -L http://t.co/tYTq5Pu | jsonpp

and make beautiful the files running around on your disk:

    jsonpp testdata/multiple/multiple.json

You can also format previously pretty-printed code with "-s":

    jsonpp -s testdata/one/singular.json

Jsonpp exists because a friend was building against an API with large JSON APIs
and was tired of the noticable wait times that other languages' tooling
had. Then, jsonpp turned out to be pretty nice for parsing the JSON log files we
had lying around and I made it a thing.


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
