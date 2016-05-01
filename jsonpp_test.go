package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestMultipleObjectFiles(t *testing.T) {
	testCmd(t, "./testdata/multiple")
}

func TestAlreadyFormattedSingleJSON(t *testing.T) {
	testCmd(t, "./testdata/one", "-s")
}

func testCmd(t *testing.T, dir string, argPrefix ...string) {
	t.Logf("Reminder: This tests a jsonpp binary in the current directory.")
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	seen := false
	for _, i := range infos {
		if strings.HasPrefix(i.Name(), "expected_") {
			continue
		}
		seen = true
		path := filepath.Join(dir, i.Name())
		expectedPath := filepath.Join(dir, "expected_"+i.Name())
		args := append(argPrefix, path)
		cmd := exec.Command("./jsonpp", args...)
		expected, err := ioutil.ReadFile(expectedPath)
		if err != nil {
			t.Errorf("unable to read %s: %s", expectedPath, err)
		}
		out, cerr := cmd.CombinedOutput()
		if !bytes.Equal(expected, out) {
			if cerr != nil {
				t.Logf("jsonpp cmd errored: %s", cerr)
			}
			t.Errorf("On %#v, expected:\n%s\n=====\nGot:\n%s\n=====", strings.Join(cmd.Args, " "), string(expected), string(out))
		}
	}

	if !seen {
		t.Fatalf("no test files in %#v", dir)
	}

}
