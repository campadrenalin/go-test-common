package gtc

import (
	"io/ioutil"
	"os"
	"testing"
)

type TempDirTracker struct {
	Dirs   []string
	Tester *testing.T
}

func NewTempDirTracker(t *testing.T) TempDirTracker {
	return TempDirTracker{make([]string, 0), t}
}

func (tracker *TempDirTracker) Create() string {
	tempdir, err := ioutil.TempDir("", "contentgremlin")
	if err == nil {
		tracker.Dirs = append(tracker.Dirs, tempdir)
	} else {
		tracker.Tester.Fatal(err)
	}
	return tempdir
}

func (tracker *TempDirTracker) Cleanup() {
	for _, dir := range tracker.Dirs {
		err := os.RemoveAll(dir)
		if err != nil {
			tracker.Tester.Error(err)
		}
	}
}
