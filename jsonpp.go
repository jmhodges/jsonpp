package main

import (
	"bytes"
	"bufio"
	"flag"
	"fmt"
	"json"
	"os"
)

var newline = []uint8("\n")

func main() {
	bufIn := bufio.NewReader(fileFromArguments())
	lastLine := []uint8("")

	arr := make([]byte, 0, 1024*1024)
	buf := bytes.NewBuffer(arr)
	lineNum := int64(1)
	for {
		line, isPrefix, err := bufIn.ReadLine()
		if err != nil && err != os.EOF {
			genericError(err)
		}

		lastLine = append(lastLine, line...)
		if !isPrefix && len(lastLine) != 0 {
			indentAndPrint(buf, lastLine, lineNum)
			lineNum += 1
			lastLine = []uint8("")
		}

		if err == os.EOF {
			os.Exit(0)
		}
	}
}

func indentAndPrint(buf *bytes.Buffer, js []uint8, lineNum int64) {
	jsErr := json.Indent(buf, js, "", "  ")
	if (jsErr != nil) {
		malformedJson(jsErr, js, lineNum)
	}
	os.Stdout.Write(buf.Bytes())
	os.Stdout.Write(newline)
	buf.Reset()
}

func malformedJson(jsErr os.Error, js []uint8, lineNum int64) {
	os.Stdout.Sync()

	synErr, isSynError := (jsErr).(*json.SyntaxError)

	fmt.Fprintf(os.Stderr, "ERROR: Broken json on line %d: %s\n", lineNum, jsErr)

	if (isSynError) {
		prefix := ""
		suffix := ""

		begin := 0
		if synErr.Offset > 15 {
			begin = int(synErr.Offset - 15)
			prefix = "..."
		}

		end := begin + 30
		if end > len(js) {
			end = len(js)
		} else {
			suffix = "..."
		}

		fmt.Fprintf(os.Stderr, "  Context: %s%s%s\n", prefix, js[begin:end], suffix)
	}

	os.Exit(1)
}

func fileFromArguments() *os.File {
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		msg := fmt.Sprintf("only specify 0 or 1 files in the arguments, not %d\n", len(args))
		genericError(os.NewError(msg))
	}
	if len(args) == 0 {
		return os.Stdin
	}
	
	file, err := os.OpenFile(args[0], os.O_RDONLY, 0)
	if err != nil {
		genericError(err)
	}
	return file
}

func genericError(err os.Error) {
	os.Stdout.Sync()
	fmt.Fprintf(os.Stderr, "ERROR: %s", err)
	os.Exit(2)
}
