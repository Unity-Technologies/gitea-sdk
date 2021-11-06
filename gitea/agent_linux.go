package gitea

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh/agent"
)

// hasAgent returns true if the ssh agent is available
func hasAgent() bool {
	if _, err := os.Stat(os.Getenv("SSH_AUTH_SOCK")); err != nil {
		return false
	}

	return true
}

// getAgent returns a ssh agent
func getAgent() (agent.Agent, error) {
	if !hasAgent() {
		return nil, fmt.Errorf("no ssh agent available")
	}

	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
	}

	return agent.NewClient(sshAgent), nil
}