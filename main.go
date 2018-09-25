package main

import (
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

var (
	use_filepath bool
	debug_mode   bool
	echo_mode    bool
	silent_mode  bool
)

func f_handler(dat string, status bool) (interface{}, error) {
	if status {
		log.Info(fmt.Sprintf("The -f flag is enabled: locating a JSON file at %s", dat))
		json_file, err := ioutil.ReadFile(dat)
		check("Failed to read the JSON file", err)

		json_bytes := []byte(json_file)

		var json_face map[string]interface{}
		valid := json.Unmarshal(json_bytes, &json_face)

		return json_face, valid

	} else {
		log.Info("The -f flag is disabled: parsing JSON from stdin")
		json_bytes := []byte(dat)

		var json_face map[string]interface{}
		err := json.Unmarshal(json_bytes, &json_face)

		return json_face, err
	}
}

func logging_handler(s_flag bool, d_flag bool) {
	// Give the silent flag precedence over the debug flag
	if s_flag {
		s_handler(s_flag)
	} else {
		d_handler(d_flag)
	}
}

func d_handler(status bool) {
	// Set the log level to fatal if the flag is down
	log.SetOutput(os.Stdout)
	if status {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.FatalLevel)
	}
	log.Info("Logger initialized")
}

func s_handler(status bool) {
	// Separate logic for the silent flag for extensibility
	log.SetOutput(ioutil.Discard)
}

func e_handler(json_data interface{}, status bool) {
	// Handle the echo flag.
	if status {
		log.Info("Echo flag is enabled, sending the JSON to stdout...")
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(json_data)
	}
}

func check(msg string, e error) error {
	// Convenience error handler.
	if e != nil {
		log.Fatal(msg)
		return nil
	}
	return nil
}

func main() {

	// Initialize the flags
	flag.BoolVar(&use_filepath, "f", false, "Use a filepath to a JSON file instead of a string as the argument.")
	flag.BoolVar(&debug_mode, "d", false, "Run in debug mode.")
	flag.BoolVar(&silent_mode, "s", false, "Silent mode. Note: Echo and debug mode will be overrriden")
	flag.BoolVar(&echo_mode, "e", false, "Echo the JSON in the command line.")
	flag.Parse()

	// Debug and silent handlers
	logging_handler(silent_mode, debug_mode)

	// Argparse
	log.Info("Parsing flags and arguments...")
	flagArgs := strings.Join(flag.Args(), "")
	json_data, err := f_handler(flagArgs, use_filepath)
	check("JSON could not be read, the file may be corrupt or invalid JSON.", err)

	// Process json and/or perform program logic here.
	// The logger and json data are available.

	// Echo handler
	e_handler(json_data, echo_mode)

	// Goodbye
	log.Info("Finished successfully and exiting...")
}
