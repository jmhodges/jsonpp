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
		if strings.HasPrefix(i.Name(), "expected_") {
			continue
		}
		testName := strings.TrimRight(i.Name(), ".json")
		path := "./data/" + testName + ".json"
		for j, expectedPath := range []string{"./data/expected_" + testName + ".json", "./data/expected_" + testName + "_tabs.json", "./data/expected_" + testName + "_4spaces.json"} {
			var cmd *exec.Cmd
			if j == 0 {
				cmd = exec.Command("./jsonpp", path)
			} else if j == 1 {
				cmd = exec.Command("./jsonpp", "-tabs", path)
			} else {
				cmd = exec.Command("./jsonpp", "-spaces=4", path)
			}
			expected, err := ioutil.ReadFile(expectedPath)
			if err != nil {
				t.Error("unable to read %s: %s", expectedPath, err)
			}
			out, cerr := cmd.CombinedOutput()
			if !bytes.Equal(expected, out) {
				if cerr != nil {
					t.Logf("jsonpp cmd errored: %s", cerr)
				}
				flag := ""
				if j == 1 {
					flag = "-tabs "
				} else if j == 2 {
					flag = "-spaces=4 "
				}
				t.Errorf("On %s%s, expected:\n%s\n=====\nGot:\n%s\n=====", flag, i.Name(), string(expected), string(out))
			}
		}
	}
}
