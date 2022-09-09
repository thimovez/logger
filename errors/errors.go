//This package demonstrated go docs.
//
//This will only be in the overview
package errors

import (
	"log"
)

//LogInfo represent the log of the error
func logInfo(message string) {
	log.Printf("INFO - %v", message)
}

//logWarning represent the log of the warning
func logWarning(message string) {
	log.Printf("WARN - %v", message)
}

//logError represent the log of the error
func logError(message string) {
	log.Printf("ERROR - %v", message)
}
