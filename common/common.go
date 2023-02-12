package common

import "log"

// HandleError checks if the returned error is nil. If it is not nil, it will log the error using log.Printf().
// An additional description can be added to clarify the type of error.
func HandleError(err error, desc ...string) {
	if err != nil {
		log.Printf("%v: %s", desc, err)
	}
}
