package log

import (
	"fmt"
	"os"
	"syscall"
)

func errorHandler(err error) error {
	if e, ok := err.(*os.PathError); ok {
		err = e.Err
	}
	if err != syscall.EPIPE && err != syscall.ERROR_BROKEN_PIPE {
		fmt.Fprintf(os.Stderr, "logger output causes an error: %v\n", err)
		return err
	}
	os.Exit(5)
	return nil
}
