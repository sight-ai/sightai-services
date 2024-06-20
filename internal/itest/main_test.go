package itest

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Set("config-path", "../../build")
	flag.Set("config-name", "config")
	flag.Parse()
	TestContext = NewContext()
	TestContext.Setup()
	code := m.Run()
	TestContext.Teardown()
	os.Exit(code)
}

func TestDummy(t *testing.T) {
	t.Log("Success")
}
