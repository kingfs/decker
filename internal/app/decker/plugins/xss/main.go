package main

import (
	"fmt"
	"os"
	"os/exec"
)

type plugin string

// --- wants input: ---
// inputsMap{
//   "host": "example.com",
// }
//
// --- gives output: ---
// resultsMap{
//  "raw_output": "...",
// }
func (p plugin) Run(inputsMap, resultsMap *map[string]string) {
	var (
		cmdOut []byte
		err    error
	)

	targetHost := (*inputsMap)["host"]

	cmdName := "python"
	cmdArgs := []string{"/usr/bin/XssPy/XssPy.py", "-v", "-u", targetHost}

	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running xss command: ", err)
		return
	}

	output := string(cmdOut)

	(*resultsMap)["raw_output"] = output
	fmt.Println("If vulnerable, try these payloads later: https://pastebin.com/J1hCfL9J")
}

// Plugin is an implementation of github.com/stevenaldinger/decker/pkg/plugins.Plugin
// All this includes is a single function, "Run(*map[string]string, *map[string]string)"
// which takes a map of inputs and an empty map of outputs that the Plugin
// is expected to populate
var Plugin plugin
