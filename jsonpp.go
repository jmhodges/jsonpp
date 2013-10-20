package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var help = flag.Bool("help", false, "help")
var tabs = flag.Bool("tabs", false, "indent using tabs")
var spaces = flag.Int("spaces", 2, "indent using spaces (default: 2)")

func main() {
	flag.Parse()
	if *help {
		cmd := os.Args[0]
		if cmd[0:2] == "./" {
			cmd = cmd[2:]
		}
		fmt.Fprintf(os.Stderr, "Usage: "+cmd+" [-tabs | -spaces=N] [file]"+"\n")
		fmt.Fprintf(os.Stderr, "   or: $COMMAND | "+cmd+" [-tabs | -spaces=N]\n")
		os.Exit(0)
	}

	var indent string

	if *tabs {
		indent = "\t"
	} else {
		if *spaces <= 0 {
			fmt.Fprintf(os.Stderr, "spaces must be greater than 0\n")
			os.Exit(1)
		}

		indent = strings.Repeat(" ", *spaces)
	}

	var exitStatus = 0
	if len(flag.Args()) > 0 {
		for _, filename := range flag.Args() {
			file, err := os.OpenFile(filename, os.O_RDONLY, 0)
			if err != nil {
				printError(err)
				exitStatus = 1
				continue
			}
			defer file.Close()

			status := processFile(file, indent)
			if status > 0 {
				exitStatus = status
			}
		}
	} else {
		status := processFile(os.Stdin, indent)
		if status > 0 {
			exitStatus = status
		}
	}
	os.Exit(exitStatus)
}

func processFile(fn *os.File, indent string) int {
	bufIn := bufio.NewReader(fn)
	arr := make([]byte, 0, 1024*1024)
	buf := bytes.NewBuffer(arr)
	lineNum := int64(1)
	for {
		lastLine, err := bufIn.ReadBytes('\n')
		if err != nil && err != io.EOF {
			printError(err)
			return 2
		}

		if err == io.EOF && len(lastLine) == 0 {
			break
		}

		status := indentAndPrint(buf, indent, lastLine, lineNum)
		if status > 0 {
			return status
		}

		buf.Reset()
		lineNum += 1

		if err == io.EOF {
			break
		}
	}

	return 0
}

func indentAndPrint(buf *bytes.Buffer, indent string, js []byte, lineNum int64) int {
	jsErr := json.Indent(buf, js, "", indent)
	if jsErr != nil {
		malformedJSON(jsErr, js, lineNum)
		return 1
	}
	os.Stdout.Write(buf.Bytes())

	return 0
}

func malformedJSON(jsErr error, js []uint8, lineNum int64) {
	os.Stdout.Sync()

	synErr, isSynError := (jsErr).(*json.SyntaxError)
	if isSynError {
		fmt.Fprintf(os.Stderr, "ERROR: Broken json on line %d, char %d: %s\n", lineNum, synErr.Offset, jsErr)
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
		b := bytes.TrimRight(js[begin:end], "\r\n")
		fmt.Fprintf(os.Stderr, "  Context: %s%s%s\n", prefix, b, suffix)
	} else {
		fmt.Fprintf(os.Stderr, "ERROR: Broken json on line %d: %s\n", lineNum, jsErr)
	}
}

func printError(err error) {
	os.Stdout.Sync()
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
}
