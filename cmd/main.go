package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	agentRegistrationSocketName string = "agent-registration.sock"
	discoveryHandlerPrefix      string = "akri-discovery-handler-template"
)

var (
	agentRegistrationSocketPath string
	discoveryHandlerName        string
	discoveryServiceSocketPath  string
)

func init() {
	envVar := os.Getenv("DISCOVERY_HANDLER_SUFFIX")
	discoveryHandlerName = fmt.Sprintf("%s%s", discoveryHandlerPrefix, envVar)

	envVar = os.Getenv("DISCOVERY_HANDLERS_DIRECTORY")
	if envVar == "" {
		envVar = "/var/lib/akri"
	}
	agentRegistrationSocketPath = filepath.Join(envVar, agentRegistrationSocketName)
	discoveryServiceSocketPath = filepath.Join(envVar, fmt.Sprintf("%s.sock", discoveryHandlerName))
	os.RemoveAll(discoveryServiceSocketPath)
}

func main() {
	fmt.Println("akri-discovery-handler-template started")
}
