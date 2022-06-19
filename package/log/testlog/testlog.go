package testlog

import (
	"flag"
	"fmt"
	"os"

	"github.com/livebud/bud/package/log/filter"

	"github.com/livebud/bud/package/log"
	"github.com/livebud/bud/package/log/console"
)

var pattern = flag.String("log", "info", "choose a log level")

// Pattern returns the log level pattern so we can pass it through arguments.
func Pattern() string {
	return *pattern
}

// Log is a logger for testing. You can set the log level for a given test by
// using the --log=pattern flag. For example, --log=debug will log all debug
// messages.
func Log() log.Interface {
	handler, err := filter.Load(console.New(os.Stderr), *pattern)
	if err != nil {
		// TODO: use a custom flag to fail earlier
		panic(fmt.Sprintf("testlog: invalid log pattern %q" + *pattern))
	}
	return log.New(handler)
}
