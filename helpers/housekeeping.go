package helpers

import (
	"log"

	"github.com/juju/ansiterm"
)

// AutoFlush handles the management of any errors that may be outputted after flushing the tab writer
func AutoFlush(tr *ansiterm.TabWriter) {
	err := tr.Flush()

	if err != nil {
		log.Fatal(err)
	}
}
