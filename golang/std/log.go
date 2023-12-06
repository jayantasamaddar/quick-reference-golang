package std

import (
	"log"
	"os"
)

func LogOperationsDemo() {
	// Define Custom Loggers using the log.New() method. Takes in an io.Writer type, a prefix string and flags defined as constants as shown below
	// log.Ldate, log.Ltime, log.Lshortfile are constants defined in the Go standard library at: https://pkg.go.dev/log#pkg-constants

	// Write to the standard output
	infoLogger := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	/********************************************************************************************************************/
	// Write Logs to a file
	// STEPS:

	// (1) Create the file
	// (2) Create a Custom Logger using log.New() that takes the created file as io.Writer type
	/********************************************************************************************************************/

	file, err := os.Create("logfile.txt")
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer file.Close()

	fileLogger := log.New(file, "Log:\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger.Println("This message is a piece of information")
	errorLogger.Println("This is an error message")
	fileLogger.Println("Demo log message to file")
}
