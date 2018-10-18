package sigfox

import (
	"log"
	"os"
)

// dprintf prints debug messages to stderr if the environment variable
// SIGFOX_DEBUG is non-zero.
func dprintf(format string, a ...interface{}) {
	if os.Getenv("SIGFOX_DEBUG") != "" {
		log.Printf(format, a...)
	}
}
