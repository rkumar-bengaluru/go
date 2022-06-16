package config

import (
	"fmt"
	"github.com/rkumar-bengaluru/go/logger"
	"testing"
)

func TestConfigReader(t *testing.T) {
	config := NewConfigReader(logger.NewConsoleLogger())
	if r := config.Initialize(); r != nil {
		t.Fatalf("error initializing config reader..")
	}
	port, err := config.GetIntegerKey("server.port")
	if err != nil {
		t.Errorf("was expecting server:port %v, but got %v\n", port, err)
	}
	fmt.Println(port)
}
