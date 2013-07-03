package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
)

func TestFiles(t *testing.T) {
	t.Logf("Reminder: This tests a jsonpp in the repo directory.")
	infos, err := ioutil.ReadDir("./data")
	if err != nil {
		t.Fatal(err)
	}
	for _, i := range infos {
		if i.Name() == "tweet.json" {
			continue
		}
		if strings.HasPrefix(i.Name(), "expected_") {
			continue
		}
		path := "./data/" + i.Name()
		expectedPath := "./data/expected_" + i.Name()
		cmd := exec.Command("./jsonpp", path)
		expected, err := ioutil.ReadFile(expectedPath)
		if err != nil {
			t.Error("unable to read %s: %s", expectedPath, err)
		}
		out, cerr := cmd.CombinedOutput()
		if !bytes.Equal(expected, out) {
			if cerr != nil {
				t.Logf("jsonpp cmd errored: %s", cerr)
			}
			t.Errorf("On %s, expected:\n%s\n=====\nGot:\n%s\n=====", i.Name(), string(expected), string(out))
		}
	}
}
