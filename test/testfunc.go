// Package test init test directories
package test

// https://intellij-support.jetbrains.com/hc/en-us/community/posts/360009685279-Go-test-working-directory-keeps-changing-to-dir-of-the-test-file-instead-of-value-in-template
import (
	"os"
	"path"
	"runtime"
	"testing"
)

// TestDir working dir for test
var TestDir string

// TestData directory for working Attachments
var TestData string

// Testinit set test directory
func Testinit(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)
	err := os.Chdir(dir)
	if err == nil {
		TestDir = dir
		TestData = path.Join(TestDir, "testdata")
		// create data directory and ignore errors
		err = os.Mkdir(TestData, 0750)
		if err != nil && !os.IsExist(err) {
			t.Fatalf("Init error:%s", err)
		}
		t.Logf("Test in %s", dir)
	} else {
		t.Fatalf("Init error:%s", err)
	}
}
