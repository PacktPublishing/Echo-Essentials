package main

import (
	"testing"
)

func TestRunMain(t *testing.T) {
	TestRun = true
	go main()
	<-StopTestServer
	TestRun = false
}
