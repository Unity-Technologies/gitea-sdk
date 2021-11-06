import (
	"fmt"

	"github.com/davidmz/go-pageant"
	"golang.org/x/crypto/ssh/agent"
)

// hasAgent returns true if pageant is available
func hasAgent() bool {
	return pageant.Available()
}

// getAgent returns a ssh agent
func getAgent() (agent.Agent, error) {
	if !hasAgent() {
		return nil, fmt.Errorf("no pageant available")
	}

	return pageant.New(), nil
}